// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721VotesUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";

/**
 * @title IERC5192
 * @notice Interface for contracts that are compatible with the ERC721 standard.
 */
interface IERC5192 {
    /**
     * @notice Emitted when the locking status is changed to locked.
     * @dev If a token is minted and the status is locked, this event should be emitted.
     * @param tokenId The identifier for a token.
     */
    event Locked(uint256 tokenId);

    /**
     * @notice Emitted when the locking status is changed to unlocked.
     * @dev If a token is minted and the status is unlocked, this event should be emitted.
     * @param tokenId The identifier for a token.
     */
    event Unlocked(uint256 tokenId);

    /**
     * @notice Returns the locking status of an Soulbound Token
     * @dev SBTs assigned to zero address are considered invalid, and queries about them do throw.
     * @param tokenId The identifier for an SBT.
     */
    function locked(uint256 tokenId) external view returns (bool);
}

abstract contract KromaSoulBoundERC721 is
    Initializable,
    IERC5192,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable,
    PausableUpgradeable,
    OwnableUpgradeable,
    EIP712Upgradeable,
    ERC721VotesUpgradeable
{
    using CountersUpgradeable for CountersUpgradeable.Counter;

    CountersUpgradeable.Counter private _tokenIdCounter;
    bool private isLocked;

    error ErrLocked();
    error ErrNotFound();

    modifier checkLock() {
        if (isLocked) revert ErrLocked();
        _;
    }

    /**
     * @custom:oz-upgrades-unsafe-allow constructor
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @param _name   ERC721 name.
     * @param _symbol ERC721 symbol.
     * @param _owner  Owner of token.
     */
    function __KromaSoulBoundERC721_init(
        string memory _name,
        string memory _symbol,
        address _owner
    ) internal onlyInitializing {
        __KromaSoulBoundERC721_init_unchained(true);
        __ERC721_init(_name, _symbol);
        __ERC721Enumerable_init();
        __ERC721URIStorage_init();
        __Pausable_init();
        __EIP712_init(_name, "1");
        __ERC721Votes_init();
        _transferOwnership(_owner);
    }

    function __KromaSoulBoundERC721_init_unchained(bool _isLocked) internal onlyInitializing {
        isLocked = _isLocked;
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function safeMint(address to, string memory uri) public onlyOwner {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
        _delegate(to, to);
    }

    function burn(uint256 tokenId) public onlyOwner {
        _burn(tokenId);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId,
        uint256 batchSize
    ) internal override(ERC721Upgradeable, ERC721EnumerableUpgradeable) whenNotPaused {
        super._beforeTokenTransfer(from, to, tokenId, batchSize);
    }

    function locked(uint256 tokenId) external view returns (bool) {
        if (!_exists(tokenId)) revert ErrNotFound();
        return isLocked;
    }

    // The following functions are overridden cause required by Solidity.

    function _afterTokenTransfer(
        address from,
        address to,
        uint256 tokenId,
        uint256 batchSize
    ) internal override(ERC721Upgradeable, ERC721VotesUpgradeable) {
        super._afterTokenTransfer(from, to, tokenId, batchSize);
    }

    function _burn(uint256 tokenId)
        internal
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
    {
        super._burn(tokenId);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory data
    ) public override(IERC721Upgradeable, ERC721Upgradeable) checkLock {
        super.safeTransferFrom(from, to, tokenId, data);
    }

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override(IERC721Upgradeable, ERC721Upgradeable) checkLock {
        super.safeTransferFrom(from, to, tokenId);
    }

    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override(IERC721Upgradeable, ERC721Upgradeable) checkLock {
        super.transferFrom(from, to, tokenId);
    }

    function approve(address approved, uint256 tokenId)
        public
        virtual
        override(IERC721Upgradeable, ERC721Upgradeable)
        checkLock
    {
        super.approve(approved, tokenId);
    }

    function setApprovalForAll(address operator, bool approved)
        public
        override(IERC721Upgradeable, ERC721Upgradeable)
        checkLock
    {
        super.setApprovalForAll(operator, approved);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable, ERC721URIStorageUpgradeable)
        returns (bool)
    {
        return interfaceId == type(IERC5192).interfaceId || super.supportsInterface(interfaceId);
    }
}
