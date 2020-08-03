package config

import (
	"errors"
	"log"
	"net"

	"github.com/kradalby/metallb-neighbour-helper/pkg/provider"
	yaml "gopkg.in/yaml.v2"
)

// Proto holds the protocol we are speaking.
type Provider string

// MetalLB supported protocols.
const (
	OpnSense Provider = "opnsense"
	VCloud   Provider = "vcloud"
)

var (
	errProviderProviderMissing    = errors.New("provider is missing from provider definition")
	errProviderNameMissing        = errors.New("name is missing from provider definition")
	errProviderURLMissing         = errors.New("url is missing from provider definition")
	errProviderPeerAddressMissing = errors.New("peer-address is missing from provider definition")
	errProviderPeerAddressValidIP = errors.New("peer-address is not a valid IP address")

	errOPNsenseProviderKeyMissing    = errors.New("key is missing from OPNsense provider")
	errOPNsenseProviderSecretMissing = errors.New("secret is missing from OPNsense provider")

	errVCloudProviderUserMissing     = errors.New("user is missing from vCloud provider")
	errVCloudProviderPasswordMissing = errors.New("password is missing from vCloud provider")
	errVCloudProviderOrgMissing      = errors.New("org is missing from vCloud provider")
	errVCloudProviderVdcMissing      = errors.New("vdc is missing from vCloud provider")
)

type providerConfig struct {
	Provider    Provider
	Name        string
	URL         string // `yaml:"url"`
	InSecure    bool   `yaml:"in-secure"`
	PeerAddress string `yaml:"peer-address"`

	// OPNSense
	Key    string // `yaml:"key"`
	Secret string // `yaml:"secret"`

	// vCloud
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Org      string `yaml:"org"`
	Vdc      string `yaml:"vdc"`
}

type config struct {
	Providers []providerConfig `yaml:"providers"`
}

//
func Parse(bs []byte) ([]provider.BgpProvider, error) {
	var raw config
	if err := yaml.UnmarshalStrict(bs, &raw); err != nil {
		return nil, err
	}

	providers := []provider.BgpProvider{}

	for _, prov := range raw.Providers {
		if prov.Provider == "" {
			return nil, errProviderProviderMissing
		}

		if prov.Name == "" {
			return nil, errProviderNameMissing
		}

		if prov.URL == "" {
			return nil, errProviderURLMissing
		}

		if prov.PeerAddress == "" {
			return nil, errProviderPeerAddressMissing
		}

		peerIPAddress := net.ParseIP(prov.PeerAddress)
		if peerIPAddress == nil {
			return nil, errProviderPeerAddressValidIP
		}

		switch prov.Provider {
		case OpnSense:
			if prov.Key == "" {
				return nil, errOPNsenseProviderKeyMissing
			}

			if prov.Secret == "" {
				return nil, errOPNsenseProviderSecretMissing
			}

			opn, err := provider.NewOpnSenseProvider(prov.URL, prov.Key, prov.Secret, peerIPAddress, prov.InSecure)
			if err != nil {
				return nil, err
			}

			providers = append(providers, opn)

		case VCloud:
			if prov.User == "" {
				return nil, errVCloudProviderUserMissing
			}

			if prov.Password == "" {
				return nil, errVCloudProviderPasswordMissing
			}

			if prov.Org == "" {
				return nil, errVCloudProviderOrgMissing
			}

			if prov.Vdc == "" {
				return nil, errVCloudProviderVdcMissing
			}

			log.Println("vCloud is not supported yet")

		default:
			log.Println("Got unsupported provider")
		}
	}

	return providers, nil
}

// Testing
// func main() {
// 	config := `providers:
//     - name: "opnTest"
//       provider: "opnsense"
//       url: "http://localhost:8080"
//       in-secure: True
//       key: "6X6860M4fOqJUmoJV9JDHikEucE+UMIi/75uZzo1TzGz1WB0RTbIpBgHdvqBo7Xj6vsWb80rkiYWcZFN"
//       secret: "6X6860M4fOqJUmoJV9JDHikEucE+UMIi/75uZzo1TzGz1WB0RTbIpBgHdvqBo7Xj6vsWb80rkiYWcZFN"
//     - name: "vCloudTest"
//       provider: "vcloud"
//       url: "http://localhost:8080"
//       in-secure: True
//       user: kradalby
//       password: password
//       org: organization
//       vdc: datacenter
//     `
// 	b := []byte(config)
// 	providers, err := Parse(b)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Printf("%#v", providers)

// }
