package resources

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
	"os"
)

type Config struct {
	ConfigEntryKind	string
	Datacenter	string
	Address		string
	Token		string
	Namespace	string
}

type Client struct {
	TerraformBridge	*bridge.TerraformBridge

	// TODO You can continue to refine your client
	ConsulClient	*consulapi.Client

	conf	*Config
}

func newClient(conf *Config) *Client {
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
		Address:	addr,
		Token:		token,
		Datacenter:	dc,
		Namespace:	ns,
	}

	os.Setenv("CONSUL_TOKEN", token)

	consulClient, err := consulapi.NewClient(&config)
	if err != nil {
		panic(err)
	}

	return &Client{
		ConsulClient:	consulClient,
		conf:		conf,
	}
}
