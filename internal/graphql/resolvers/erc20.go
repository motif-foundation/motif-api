// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"motif-api/internal/repository"
	"motif-api/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ERC20Token represents a generic ERC20 token
type ERC20Token struct {
	types.Erc20Token
}

// NewErc20Token creates a new instance of resolvable ERC20 token, it also validates
// the token existence by loading the total supply of the token
// before making a resolvable instance.
func NewErc20Token(adr *common.Address) *ERC20Token {
	// get the total supply of the token and validate the token existence
	erc20, err := repository.R().Erc20Token(adr)
	if err != nil {
		return nil
	}
	// make the instance of the token
	return &ERC20Token{*erc20}
}

// Erc20Token resolves an instance of ERC20 token if available.
func (rs *rootResolver) Erc20Token(args *struct{ Token common.Address }) *ERC20Token {
	return NewErc20Token(&args.Token)
}

// FMintTokenAllowance resolves the amount of ERC20 tokens unlocked
// by the token owner for DeFi operations.
func (rs *rootResolver) FMintTokenAllowance(args *struct {
	Owner common.Address
	Token common.Address
}) (hexutil.Big, error) {
	return repository.R().Erc20Allowance(&args.Token, &args.Owner, nil)
}

// ErcTotalSupply resolves the current total supply of the specified token.
func (rs *rootResolver) ErcTotalSupply(args *struct{ Token common.Address }) (hexutil.Big, error) {
	return repository.R().Erc20TotalSupply(&args.Token)
}

// ErcTokenBalance resolves the current available balance of the specified token
// for the specified owner.
func (rs *rootResolver) ErcTokenBalance(args *struct {
	Owner common.Address
	Token common.Address
}) (hexutil.Big, error) {
	return repository.R().Erc20BalanceOf(&args.Token, &args.Owner)
}

// ErcTokenAllowance resolves the current amount of ERC20 tokens unlocked
// by the token owner for the spender to be manipulated with.
func (rs *rootResolver) ErcTokenAllowance(args *struct {
	Token   common.Address
	Owner   common.Address
	Spender common.Address
}) (hexutil.Big, error) {
	return repository.R().Erc20Allowance(&args.Token, &args.Owner, &args.Spender)
}

// TotalSupply resolves the total supply of the given ERC20 token.
func (token *ERC20Token) TotalSupply() (hexutil.Big, error) {
	return repository.R().Erc20TotalSupply(&token.Address)
}

// BalanceOf resolves the available balance of the given ERC20 token to a user.
func (token *ERC20Token) BalanceOf(args *struct{ Owner common.Address }) (hexutil.Big, error) {
	return repository.R().Erc20BalanceOf(&token.Address, &args.Owner)
}

// Allowance resolves the unlocked allowance of the given ERC20 token from the owner to spender.
func (token *ERC20Token) Allowance(args *struct {
	Owner   common.Address
	Spender common.Address
}) (hexutil.Big, error) {
	return repository.R().Erc20Allowance(&token.Address, &args.Owner, &args.Spender)
}

// LogoURL resolves an URL of the token logo.
func (token *ERC20Token) LogoURL() string {
	return repository.R().Erc20LogoURL(&token.Address)
}

// TotalDeposit represents the total amount of tokens deposited to fMint as collateral.
func (token *ERC20Token) TotalDeposit() (hexutil.Big, error) {
	return repository.R().FMintTokenTotalBalance(&token.Address, types.DefiTokenTypeCollateral)
}

// TotalDebt represents the total amount of tokens borrowed/minted on fMint.
func (token *ERC20Token) TotalDebt() (hexutil.Big, error) {
	return repository.R().FMintTokenTotalBalance(&token.Address, types.DefiTokenTypeDebt)
}
