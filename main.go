package sdk

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/shopspring/decimal"
)

type TokenBalance struct {
	OwnerAccount solana.PublicKey `json:"ownerAccount"`
	Weight       decimal.Decimal  `json:"weight"`
}

func (c *Client) FetchTokenBalance(tokenMint solana.PublicKey, timestamp int64) ([]TokenBalance, error) {
	path := fmt.Sprintf("/token-balance/%s?epoch=%d", tokenMint, timestamp)
	result := []TokenBalance{}
	err := c.get(path, &result)
	if err != nil {
		c.Log.Info().Msgf("fetched tokenBalance owners len: %d", len(result))
	}
	return result, err
}

func (c *Client) FetchTokenBalance2(tokenMint solana.PublicKey, start, end int64) ([]TokenBalance, error) {
	path := fmt.Sprintf("/token-balance/%s?start=%d&end=%d", tokenMint, start, end)
	result := []TokenBalance{}
	err := c.get(path, &result)
	if err != nil {
		c.Log.Info().Msgf("fetched tokenBalance owners len: %d", len(result))
	}
	return result, err
}

type ParrotVault struct {
	OwnerAccount     solana.PublicKey `json:"ownerAccount"`
	DebtWeight       decimal.Decimal  `json:"debtWeight"`
	CollateralWeight decimal.Decimal  `json:"collateralWeight"`
}

func (c *Client) FetchParrotVault(vaultType solana.PublicKey, timestamp int64) ([]ParrotVault, error) {
	path := fmt.Sprintf("/parrot-vaults/%s?epoch=%d", vaultType, timestamp)
	result := []ParrotVault{}
	err := c.get(path, &result)
	if err != nil {
		c.Log.Info().Msgf("fetched parrotVault owners len: %d", len(result))
	}
	return result, err
}

func (c *Client) FetchParrotVault2(vaultType solana.PublicKey, start, end int64) ([]ParrotVault, error) {
	path := fmt.Sprintf("/parrot-vaults/%s?start=%d&end=%d", vaultType, start, end)
	fmt.Println("path:", path)
	result := []ParrotVault{}
	err := c.get(path, &result)
	if err != nil {
		c.Log.Info().Msgf("fetched parrotVault owners len: %d", len(result))
	}
	return result, err
}
