package sdk

import (
	"fmt"
)

type TokenBalance struct {
	OwnerAccount string `json:"ownerAccount"`
	Weight       uint64 `json:"weight,omitempty"`
}

func (c *Client) FetchTokenBalance(tokenMint string, timestamp uint64) ([]TokenBalance, error) {
	path := fmt.Sprintf("/token-balance/%s?timestamp=%d", tokenMint, timestamp)
	result := []TokenBalance{}
	err := c.get(path, &result)
	return result, err
}

type ParrotVault struct {
	Owner            string `json:"ownerAccount"`
	DebtWeight       uint64 `json:"debtWeight,omitempty"`
	CollateralWeight uint64 `json:"collateralWeight,omitempty"`
}

func (c *Client) FetchParrotVault(vaultType string, timestamp uint64) ([]ParrotVault, error) {
	path := fmt.Sprintf("/parrot-vaults/%s?timestamp=%d", vaultType, timestamp)
	result := []ParrotVault{}
	err := c.get(path, &result)
	return result, err
}
