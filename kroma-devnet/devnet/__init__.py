import argparse
import logging
import os
import subprocess
import json
import socket
import calendar
import datetime
import time
import shutil
import http.client
import gzip
from multiprocessing import Process, Queue
import concurrent.futures
from collections import namedtuple


import devnet.log_setup

pjoin = os.path.join

parser = argparse.ArgumentParser(description='Bedrock devnet launcher')
parser.add_argument('--monorepo-dir', help='Directory of the monorepo', default=os.getcwd())
parser.add_argument('--allocs', help='Only create the allocs and exit', type=bool, action=argparse.BooleanOptionalAction)
parser.add_argument('--test', help='Tests the deployment, must already be deployed', type=bool, action=argparse.BooleanOptionalAction)

log = logging.getLogger()

# Global environment variables
DEVNET_NO_BUILD = os.getenv('DEVNET_NO_BUILD') == "true"
DEVNET_PLASMA = os.getenv('DEVNET_PLASMA') == "true"

class Bunch:
    def __init__(self, **kwds):
        self.__dict__.update(kwds)

class ChildProcess:
    def __init__(self, func, *args):
        self.errq = Queue()
        self.process = Process(target=self._func, args=(func, args))

    def _func(self, func, args):
        try:
            func(*args)
        except Exception as e:
            self.errq.put(str(e))

    def start(self):
        self.process.start()

    def join(self):
        self.process.join()

    def get_error(self):
        return self.errq.get() if not self.errq.empty() else None


def main():
    args = parser.parse_args()

    monorepo_dir = os.path.abspath(args.monorepo_dir)
    devnet_dir = pjoin(monorepo_dir, '.devnet')
    contracts_bedrock_dir = pjoin(monorepo_dir, 'packages', 'contracts')
    deployment_dir = pjoin(contracts_bedrock_dir, 'deployments', 'devnetL1')
    op_node_dir = pjoin(args.monorepo_dir, 'op-node')
    ops_bedrock_dir = pjoin(monorepo_dir, 'ops-devnet')
    deploy_config_dir = pjoin(contracts_bedrock_dir, 'deploy-config')
    devnet_config_path = pjoin(deploy_config_dir, 'devnetL1.json')
    devnet_config_template_path = pjoin(deploy_config_dir, 'devnetL1-template.json')
    ops_chain_ops = pjoin(monorepo_dir, 'kroma-chain-ops')
    sdk_dir = pjoin(monorepo_dir, 'packages', 'sdk')

    paths = Bunch(
      mono_repo_dir=monorepo_dir,
      devnet_dir=devnet_dir,
      contracts_bedrock_dir=contracts_bedrock_dir,
      deployment_dir=deployment_dir,
      l1_deployments_path=pjoin(deployment_dir, '.deploy'),
      deploy_config_dir=deploy_config_dir,
      devnet_config_path=devnet_config_path,
      devnet_config_template_path=devnet_config_template_path,
      op_node_dir=op_node_dir,
      ops_bedrock_dir=ops_bedrock_dir,
      ops_chain_ops=ops_chain_ops,
      sdk_dir=sdk_dir,
      genesis_l1_path=pjoin(devnet_dir, 'genesis-l1.json'),
      genesis_l2_path=pjoin(devnet_dir, 'genesis-l2.json'),
      allocs_path=pjoin(devnet_dir, 'allocs-l1.json'),
      addresses_json_path=pjoin(devnet_dir, 'addresses.json'),
      sdk_addresses_json_path=pjoin(devnet_dir, 'sdk-addresses.json'),
      rollup_config_path=pjoin(devnet_dir, 'rollup.json')
    )

    if args.test:
      log.info('Testing deployed devnet')
      devnet_test(paths)
      return

    os.makedirs(devnet_dir, exist_ok=True)

    if args.allocs:
        devnet_l1_genesis(paths)
        return

    git_commit = subprocess.run(['git', 'rev-parse', 'HEAD'], capture_output=True, text=True).stdout.strip()
    git_date = subprocess.run(['git', 'show', '-s', "--format=%ct"], capture_output=True, text=True).stdout.strip()

    # CI loads the images from workspace, and does not otherwise know the images are good as-is
    if DEVNET_NO_BUILD:
        log.info('Skipping docker images build')
    else:
        log.info(f'Building docker images for git commit {git_commit} ({git_date})')
        run_command(['docker', 'compose', 'build', '--progress', 'plain',
                     '--build-arg', f'GIT_COMMIT={git_commit}', '--build-arg', f'GIT_DATE={git_date}'],
                    cwd=paths.ops_bedrock_dir, env={
            'PWD': paths.ops_bedrock_dir,
            'DOCKER_BUILDKIT': '1', # (should be available by default in later versions, but explicitly enable it anyway)
            'COMPOSE_DOCKER_CLI_BUILD': '1'  # use the docker cache
        })

    log.info('Devnet starting')
    devnet_deploy(paths)


def deploy_contracts(paths):
    wait_up(8545)
    wait_for_rpc_server('127.0.0.1:8545')
    res = eth_accounts('127.0.0.1:8545')

    response = json.loads(res)
    account = response['result'][0]
    log.info(f'Deploying with {account}')

    # send some ether to the create2 deployer account
    run_command([
        'cast', 'send', '--from', account,
        '--rpc-url', 'http://127.0.0.1:8545',
        '--unlocked', '--value', '1ether', '0x3fAB184622Dc19b6109349B94811493BF2a45362'
    ], env={}, cwd=paths.contracts_bedrock_dir)

    # deploy the create2 deployer
    run_command([
      'cast', 'publish', '--rpc-url', 'http://127.0.0.1:8545',
      '0xf8a58085174876e800830186a08080b853604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf31ba02222222222222222222222222222222222222222222222222222222222222222a02222222222222222222222222222222222222222222222222222222222222222'
    ], env={}, cwd=paths.contracts_bedrock_dir)

    # [Kroma: START]
    # send some ether to the l1 contract deployer account
    run_command([
      'cast', 'send', '--from', account,
      '--rpc-url', 'http://127.0.0.1:8545',
      '--unlocked', '--value', '1ether', '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266'
    ], env={}, cwd=paths.contracts_bedrock_dir)

    run_command([
        'npx', 'hardhat', 'deploy', '--network', 'devnetL1', '--tags', 'setup'
    ], env={}, cwd=paths.contracts_bedrock_dir)

    run_command([
      'npx', 'hardhat', 'export-addresses', '--network', 'devnetL1'
    ], env={}, cwd=paths.contracts_bedrock_dir)
    # [Kroma: END]

    shutil.copy(paths.l1_deployments_path, paths.addresses_json_path)


def init_devnet_l1_deploy_config(paths, update_timestamp=False):
    deploy_config = read_json(paths.devnet_config_template_path)
    if update_timestamp:
        deploy_config['l1GenesisBlockTimestamp'] = '{:#x}'.format(int(time.time()))
    if DEVNET_PLASMA:
        deploy_config['usePlasma'] = True
    write_json(paths.devnet_config_path, deploy_config)

def devnet_l1_genesis(paths):
    log.info('Generating L1 genesis state')

    # [Kroma: START]
    run_command([
      'rm', '-rf', 'deploy-config/devnetL1.json', 'deployments/devnetL1'
    ], env={}, cwd=paths.contracts_bedrock_dir)
    # [Kroma: END]

    init_devnet_l1_deploy_config(paths)

    geth = subprocess.Popen([
        'geth', '--dev', '--http', '--http.api', 'eth,debug,web3',
        '--verbosity', '4', '--gcmode', 'archive', '--dev.gaslimit', '30000000',
        '--rpc.allow-unprotected-txs'
    ])

    try:
        forge = ChildProcess(deploy_contracts, paths)
        forge.start()
        forge.join()
        err = forge.get_error()
        if err:
            raise Exception(f"Exception occurred in child process: {err}")

        res = debug_dumpBlock('127.0.0.1:8545')
        response = json.loads(res)
        allocs = response['result']

        write_json(paths.allocs_path, allocs)
    finally:
        geth.terminate()


# Bring up the devnet where the contracts are deployed to L1
def devnet_deploy(paths):
    if os.path.exists(paths.genesis_l1_path):
        log.info('L1 genesis already generated.')
    else:
        log.info('Generating L1 genesis.')
        if os.path.exists(paths.allocs_path) == False:
            devnet_l1_genesis(paths)

        # It's odd that we want to regenerate the devnetL1.json file with
        # an updated timestamp different than the one used in the devnet_l1_genesis
        # function.  But, without it, CI flakes on this test rather consistently.
        # If someone reads this comment and understands why this is being done, please
        # update this comment to explain.
        init_devnet_l1_deploy_config(paths, update_timestamp=True)
        run_command([
            'go', 'run', 'cmd/main.go', 'genesis', 'l1',
            '--deploy-config', paths.devnet_config_path,
            '--l1-allocs', paths.allocs_path,
            '--l1-deployments', paths.addresses_json_path,
            '--outfile.l1', paths.genesis_l1_path,
        ], cwd=paths.op_node_dir)

    log.info('Starting L1.')
    run_command(['docker', 'compose', 'up', '-d', 'l1', 'validator'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir
    })
    wait_up(8545)
    wait_for_rpc_server('127.0.0.1:8545')

    if os.path.exists(paths.genesis_l2_path):
        log.info('L2 genesis and rollup configs already generated.')
    else:
        log.info('Generating L2 genesis and rollup configs.')
        run_command([
            'go', 'run', 'cmd/main.go', 'genesis', 'l2',
            '--l1-rpc', 'http://localhost:8545',
            '--deploy-config', paths.devnet_config_path,
            '--l1-deployments', paths.addresses_json_path,
            '--outfile.l2', paths.genesis_l2_path,
            '--outfile.rollup', paths.rollup_config_path
        ], cwd=paths.op_node_dir)

    addresses = read_json(paths.addresses_json_path)

    # Start the L2.
    log.info('Bringing up L2.')
    run_command(['docker', 'compose', 'up', '-d', 'l2'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir
    })
    run_command(['docker', 'compose', 'up', '-d', 'l2-historical'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir
    })

    # Wait for the L2 to be available.
    wait_up(9545)
    wait_for_rpc_server('127.0.0.1:9545')
    wait_up(9445)
    wait_for_rpc_server('127.0.0.1:9445')

    # [Kroma: START]
    # Print out the addresses being used for easier debugging.
    l2_output_oracle = addresses['L2OutputOracleProxy']
    log.info(f'Using L2OutputOracle {l2_output_oracle}')
    colosseum = addresses['ColosseumProxy']
    log.info(f'Using Colosseum {colosseum}')
    validator_pool = addresses['ValidatorPoolProxy']
    log.info(f'Using ValidatorPool {validator_pool}')
    validator_manager = addresses['ValidatorManagerProxy']
    log.info(f'Using ValidatorManager {validator_manager}')
    asset_manager = addresses['AssetManagerProxy']
    log.info(f'Using AssetManager {asset_manager}')

    log.info('Bringing up `kroma-node`, `kroma-batcher` and `kroma-validator`.')
    run_command(['docker', 'compose', 'up', '-d', 'kroma-node', 'kroma-node-historical', 'kroma-batcher', 'kroma-validator', 'kroma-challenger'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir,
        'L2OO_ADDRESS': l2_output_oracle,
        'COLOSSEUM_ADDRESS': colosseum,
        'VALPOOL_ADDRESS': validator_pool,
        'VALMGR_ADDRESS': validator_manager,
        'ASSETMANAGER_ADDRESS': asset_manager,
    })

    log.info("Deposit ETH into ValidatorPool contract to be a validator...")
    run_command(['docker', 'compose', 'exec', 'kroma-validator',
                 'kroma-validator', 'deposit', '--amount', '1000000000'], cwd=paths.ops_bedrock_dir)

    log.info("Register to ValidatorManager contract to be a validator...")
    run_command(
        [
            'docker',
            'compose',
            'exec',
            'kroma-validator',
            'kroma-validator',
            'register',
            '--amount',
            '100',
            '--commission-rate',
            '5',
            '--withdraw-account',
            '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
        ],
        cwd=paths.ops_bedrock_dir,
    )
    # [Kroma: END]

    log.info('Bringing up `artifact-server`')
    run_command(['docker', 'compose', 'up', '-d', 'artifact-server'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir
    })

    log.info('Devnet ready.')


def eth_accounts(url):
    log.info(f'Fetch eth_accounts {url}')
    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":2, "jsonrpc":"2.0", "method": "eth_accounts", "params":[]}'
    conn.request('POST', '/', body, headers)
    response = conn.getresponse()
    data = response.read().decode()
    conn.close()
    return data


def debug_dumpBlock(url):
    log.info(f'Fetch debug_dumpBlock {url}')
    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":3, "jsonrpc":"2.0", "method": "debug_dumpBlock", "params":["latest"]}'
    conn.request('POST', '/', body, headers)
    response = conn.getresponse()
    data = response.read().decode()
    conn.close()
    return data


def wait_for_rpc_server(url):
    log.info(f'Waiting for RPC server at {url}')

    headers = {'Content-type': 'application/json'}
    body = '{"id":1, "jsonrpc":"2.0", "method": "eth_chainId", "params":[]}'

    while True:
        try:
            conn = http.client.HTTPConnection(url)
            conn.request('POST', '/', body, headers)
            response = conn.getresponse()
            if response.status < 300:
                log.info(f'RPC server at {url} ready')
                return
        except Exception as e:
            log.info(f'Waiting for RPC server at {url}')
            time.sleep(1)
        finally:
            if conn:
                conn.close()


CommandPreset = namedtuple('Command', ['name', 'args', 'cwd', 'timeout'])


def devnet_test(paths):
    # Run the two commands with different signers, so the ethereum nonce management does not conflict
    # And do not use devnet system addresses, to avoid breaking fee-estimation or nonce values.
    run_commands([
        CommandPreset('erc20-test',
          ['npx', 'hardhat',  'deposit-erc20', '--network',  'devnetL1',
           '--l1-contracts-json-path', paths.addresses_json_path, '--signer-index', '14'],
          cwd=paths.sdk_dir, timeout=8*60),
        CommandPreset('eth-test',
          ['npx', 'hardhat',  'deposit-eth', '--network',  'devnetL1',
           '--l1-contracts-json-path', paths.addresses_json_path, '--signer-index', '15'],
          cwd=paths.sdk_dir, timeout=8*60)
    ], max_workers=2)


def run_commands(commands: list[CommandPreset], max_workers=2):
    with concurrent.futures.ThreadPoolExecutor(max_workers=max_workers) as executor:
        futures = [executor.submit(run_command_preset, cmd) for cmd in commands]

        for future in concurrent.futures.as_completed(futures):
            result = future.result()
            if result:
                print(result.stdout)


def run_command_preset(command: CommandPreset):
    with subprocess.Popen(command.args, cwd=command.cwd,
                          stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True) as proc:
        try:
            # Live output processing
            for line in proc.stdout:
                # Annotate and print the line with timestamp and command name
                timestamp = datetime.datetime.utcnow().strftime('%H:%M:%S.%f')
                # Annotate and print the line with the timestamp
                print(f"[{timestamp}][{command.name}] {line}", end='')

            stdout, stderr = proc.communicate(timeout=command.timeout)

            if proc.returncode != 0:
                raise RuntimeError(f"Command '{' '.join(command.args)}' failed with return code {proc.returncode}: {stderr}")

        except subprocess.TimeoutExpired:
            raise RuntimeError(f"Command '{' '.join(command.args)}' timed out!")

        except Exception as e:
            raise RuntimeError(f"Error executing '{' '.join(command.args)}': {e}")

        finally:
            # Ensure process is terminated
            proc.kill()
    return proc.returncode


def run_command(args, check=True, shell=False, cwd=None, env=None, timeout=None):
    env = env if env else {}
    return subprocess.run(
        args,
        check=check,
        shell=shell,
        env={
            **os.environ,
            **env
        },
        cwd=cwd,
        timeout=timeout
    )


def wait_up(port, retries=10, wait_secs=1):
    for i in range(0, retries):
        log.info(f'Trying 127.0.0.1:{port}')
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        try:
            s.connect(('127.0.0.1', int(port)))
            s.shutdown(2)
            log.info(f'Connected 127.0.0.1:{port}')
            return True
        except Exception:
            time.sleep(wait_secs)

    raise Exception(f'Timed out waiting for port {port}.')


def write_json(path, data):
    with open(path, 'w+') as f:
        json.dump(data, f, indent='  ')


def read_json(path):
    with open(path, 'r') as f:
        return json.load(f)
