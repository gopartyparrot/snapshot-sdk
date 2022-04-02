package sdk

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/rs/zerolog"
	"github.com/test-go/testify/require"
)

var prodClient *Client

func init() {
	logger := zerolog.New(os.Stdout)

	c := &Client{
		Host:   "https://snapshot-api.bunnyducky.com",
		Client: http.DefaultClient,
		Log:    &logger,
	}
	prodClient = c
}

func getTestHttpClient(t *testing.T) *Client {
	logger := zerolog.New(os.Stdout)
	client := &Client{
		// Host: "http://127.0.0.1:3500",
		// Host:   "https://staging.partyparrot.finance",
		Host:   "https://snapshot-api.bunnyducky.com",
		Client: http.DefaultClient,
		Log:    &logger,
	}

	return client
}

func TestSDK(t *testing.T) {
	client := getTestHttpClient(t)
	const timestamp = 1641499200

	tokenMint := solana.MustPublicKeyFromBase58("yPRTUpLDftNej7p6QofNYgRArRXsm6Mvkzohj4bh4WM")
	tokenBalances, err := client.FetchTokenBalance(tokenMint, timestamp)
	require.NoError(t, err)
	t.Logf("tokenBalances length %d", len(tokenBalances))

	vaultType := solana.MustPublicKeyFromBase58("8PcJ5FmtmuYQCvBhaHkVY5DKVBn8BsMtV5RVqHU4h8ir")
	vaults, err := client.FetchParrotVault(vaultType, timestamp)
	require.NoError(t, err)
	t.Logf("vaults length %d", len(vaults))
}

func TestLegacyParrotVault(t *testing.T) {
	got, err := prodClient.FetchParrotVault(solana.MustPublicKeyFromBase58("2xjZS6GzCN3UAQQMxe57eaWt7bkAYN6yzNDagEsviLBH"), 1648483200)
	require.NoError(t, err)
	spew.Dump(got)
}

//  2022-03-28 16:00:00+00  2022-03-29 12:00:00+00
// SaberLP_USTUSDC  yUT3Dqe2Nz46iy9urmqEdntQSBywvSC1KoiktkEycvN  2xjZS6GzCN3UAQQMxe57eaWt7bkAYN6yzNDagEsviLBH
// SaberLP_USDTUSDC y2poavXfuAHWRizzK52vegqzzkHUo5CzooxFzDDX6vp 2pHHpffMpLaLGc522EmV5w8T9BvHJwrDjpQEoa8fiU8Z

var (
	twentyHAgo = time.Now().Add(-20 * time.Hour)
	now        = time.Now()
)

func init() {
	fmt.Println("20h ago", twentyHAgo.Unix())
	fmt.Println("now", now.Unix())
}

func TestClient_FetchParrotVault2(t *testing.T) {
	got, err := prodClient.FetchParrotVault2(solana.MustPublicKeyFromBase58("2pHHpffMpLaLGc522EmV5w8T9BvHJwrDjpQEoa8fiU8Z"), twentyHAgo.Unix(), now.Unix())
	require.NoError(t, err)
	spew.Dump(got)
}

func TestClient_FetchTokenBalance2(t *testing.T) {
	got, err := prodClient.FetchTokenBalance2(solana.MustPublicKeyFromBase58("y2poavXfuAHWRizzK52vegqzzkHUo5CzooxFzDDX6vp"), twentyHAgo.Unix(), now.Unix())
	require.NoError(t, err)
	spew.Dump(got)
}
