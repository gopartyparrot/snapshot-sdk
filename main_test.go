package sdk

import (
	"net/http"
	"os"
	"testing"

	"github.com/test-go/testify/require"
	"go.uber.org/zap/zapcore"
)

func getTestHttpClient(t *testing.T) *Client {
	logger, recoverLog, err := NewLogger(ZapConf{
		Level: zapcore.DebugLevel,
	}, os.Stdout)
	t.Cleanup(recoverLog)
	require.NoError(t, err)
	sLogger := logger.Sugar()

	client := &Client{
		// Host: "http://127.0.0.1:3500",
		Host:   "https://staging.partyparrot.finance",
		Client: http.DefaultClient,
		Logger: sLogger,
	}

	return client
}

func TestSDK(t *testing.T) {
	client := getTestHttpClient(t)
	timestamp := uint64(1641499200)

	tokenMint := "yPRTUpLDftNej7p6QofNYgRArRXsm6Mvkzohj4bh4WM"
	tokenBalances, err := client.FetchTokenBalance(tokenMint, timestamp)
	require.NoError(t, err)
	t.Logf("tokenBalances length %d", len(tokenBalances))

	vaultType := "8PcJ5FmtmuYQCvBhaHkVY5DKVBn8BsMtV5RVqHU4h8ir"
	vaults, err := client.FetchParrotVault(vaultType, timestamp)
	require.NoError(t, err)
	t.Logf("vaults length %d", len(vaults))
}
