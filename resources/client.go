package resources

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
	"os"
)

type Config struct {
	ConfigEntryKind string `mapstructure:"config_entry_kind"`
	Datacenter      string `mapstructure:"datacenter"`
	Address         string `mapstructure:"address"`
	Token           string `mapstructure:"token"`
	Namespace       string `mapstructure:"namespace"`
}

type Client struct {
	TerraformBridge *bridge.TerraformBridge

	// TODO You can continue to refine your client
	ConsulClient *consulapi.Client

	conf *Config
}

func newClient(conf *Config) (*Client, error) {
	var dc, addr, token, ns string
	if conf.Address == "" {
		addr = os.Getenv("CONSUL_ADDRESS")
	}
	if conf.Datacenter == "" {
		dc = os.Getenv("CONSUL_DC")
	}
	if conf.ConfigEntryKind == "" {
		conf.ConfigEntryKind = os.Getenv("CONSUL_CONFIG_ENTRY_KIND")
	}
	if conf.Token == "" {
		token = os.Getenv("CONSUL_TOKEN")
	}
	if conf.Namespace == "" {
		ns = os.Getenv("CONSUL_NAMESPACE")
	}

	config := consulapi.Config{
		Address:    addr,
		Token:      token,
		Datacenter: dc,
		Namespace:  ns,
	}

	os.Setenv("CONSUL_TOKEN", token)

	consulClient, err := consulapi.NewClient(&config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("create consul client failed: %v", err))
	}

	return &Client{
		ConsulClient: consulClient,
		conf:         conf,
	}, nil
}
