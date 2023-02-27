package resources

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
	"os"
)

type Config struct {
	ConfigEntryKind string
	Datacenter      string
}

type Client struct {
	TerraformBridge *bridge.TerraformBridge

	// TODO You can continue to refine your client
	ConsulClient *consulapi.Client

	conf *Config
}

func newClient() *Client {
	config := consulapi.Config{
		Address: "localhost:8500",
		Token:   "e95b599e-166e-7d80-08ad-aee76e7ddf19",
	}

	os.Setenv("CONSUL_TOKEN", "e95b599e-166e-7d80-08ad-aee76e7ddf19")

	consulClient, err := consulapi.NewClient(&config)
	if err != nil {
		panic(err)
	}

	return &Client{
		ConsulClient: consulClient,
		conf: &Config{
			ConfigEntryKind: "proxy-defaults",
			Datacenter:      "dc1",
		},
	}
}
