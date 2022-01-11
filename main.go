package sdk

import (
	"fmt"
	"math/big"
)

type TokenBalance struct {
	OwnerAccount string `json:"ownerAccount"`
	Weight_      string `json:"weight,omitempty"`
}

func (tb TokenBalance) Weight() *big.Int {
	bn, _ := new(big.Int).SetString(tb.Weight_, 10)
	return bn
}

func (c *Client) FetchTokenBalance(tokenMint string, timestamp uint64) ([]TokenBalance, error) {
	path := fmt.Sprintf("/token-balance/%s?epoch=%d", tokenMint, timestamp)
	result := []TokenBalance{}
	err := c.get(path, &result)
	return result, err
}

type ParrotVault struct {
	Owner             string `json:"ownerAccount"`
	DebtWeight_       string `json:"debtWeight,omitempty"`
	CollateralWeight_ string `json:"collateralWeight,omitempty"`
}

func (pv ParrotVault) DebtWeight() *big.Int {
	bn, _ := new(big.Int).SetString(pv.DebtWeight_, 10)
	return bn
}

func (pv ParrotVault) CollateralWeight() *big.Int {
	bn, _ := new(big.Int).SetString(pv.CollateralWeight_, 10)
	return bn
}

func (c *Client) FetchParrotVault(vaultType string, timestamp uint64) ([]ParrotVault, error) {
	path := fmt.Sprintf("/parrot-vaults/%s?epoch=%d", vaultType, timestamp)
	result := []ParrotVault{}
	err := c.get(path, &result)
	return result, err
}
