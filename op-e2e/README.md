# op-e2e

The end to end tests in this repo depend on genesis state that is
created with the `kroma-devnet` package. To create this state,
run the following commands from the root of the repository:

```bash
make install-geth
make e2e-allocs
```

This will leave artifacts in the `.e2e` directory that will be
read into `op-e2e` at runtime. The default deploy configuration
used for starting all `op-e2e` based tests can be found in
`packages/contracts/deploy-config/e2eL1.json`. There
are some values that are safe to change in memory in `op-e2e` at
runtime, but others cannot be changed or else it will result in
broken tests. Any changes to `e2eL1.json` should result in
rebuilding the `.e2e` artifacts before the new values will
be present in the `op-e2e` tests.

## Running tests
Consult the [Makefile](./Makefile) in this directory. Run, e.g.:

```bash
make test-http
```

### Troubleshooting
If you encounter errors:
* ensure you have the latest version of foundry installed: `pnpm update:foundry`
* try deleting the `packages/contracts/forge-artifacts` directory
* if the above step doesn't fix the error, try `pnpm clean`
