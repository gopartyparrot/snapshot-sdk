package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Client struct {
	Host   string
	Client *http.Client
	Log    *zerolog.Logger
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.Client.Get(fmt.Sprintf("%s/api/v1/%s", c.Host, strings.TrimLeft(path, "/")))
	if err != nil {
		return errors.Wrap(err, "http post err")
	}
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.Log.Debug().
		Str("path", path).
		// Str("body", string(respBodyBytes)).
		Int("bodyLen", len(respBodyBytes)).
		Msg("invest sdk post response")
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("none ok status: %d", resp.StatusCode)
	}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return errors.Wrap(err, "json unmarshal result failed")
	}
	return nil
}
