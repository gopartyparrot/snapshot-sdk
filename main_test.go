package sdk

import (
	"net/http"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/test-go/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var testLogger *zap.SugaredLogger

var prodClient *Client

func init() {
	logger, _, _ := NewLogger(ZapConf{
		Level: zapcore.DebugLevel,
	}, os.Stdout)
	testLogger = logger.Sugar()

	c := &Client{
		Host:   "https://snapshot-api.bunnyducky.com",
		Client: http.DefaultClient,
		Logger: testLogger,
	}
	prodClient = c
}

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

func TestClient_FetchParrotVault2(t *testing.T) {
	got, err := prodClient.FetchParrotVault2("2chxdDkAYXuhcosfasR6sMDMhZkHUW28ngmw5dnojufd", 1646752126, 1646824126)
	require.NoError(t, err)
	spew.Dump(got)
}

func TestClient_FetchTokenBalance2(t *testing.T) {
	got, err := prodClient.FetchTokenBalance2("yAC4RaXvFZpNhnXZVgiiYj4cgC1G5SbgX5jzdaEXgyA", 1646752126, 1646824126)
	require.NoError(t, err)
	spew.Dump(got)
}
