package main

import (
	"fmt"
	"log"
	"net"

	yaml "gopkg.in/yaml.v2"
)

// Proto holds the protocol we are speaking.
type Provider string

// MetalLB supported protocols.
const (
	OpnSense Provider = "opnsense"
	VCloud   Provider = "vcloud"
)

type provider struct {
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
	Providers []provider `yaml:"providers"`
}

//
func Parse(bs []byte) ([]BgpProvider, error) {
	var raw config
	if err := yaml.UnmarshalStrict(bs, &raw); err != nil {
		return nil, fmt.Errorf("could not parse config: %s", err)
	}

	providers := []BgpProvider{}

	for _, prov := range raw.Providers {

		if prov.Provider == "" {
			return nil, fmt.Errorf("'provider' has to be set for provider: %#v", prov)
		}
		if prov.Name == "" {
			return nil, fmt.Errorf("'name' has to be set for provider: %#v", prov)
		}
		if prov.URL == "" {
			return nil, fmt.Errorf("'url' has to be set for provider: %#v", prov)
		}

		if prov.PeerAddress == "" {
			return nil, fmt.Errorf("'peer-address' has to be set for provider: %#v", prov)
		}

		peerIPAddress := net.ParseIP(prov.PeerAddress)
		if peerIPAddress == nil {
			return nil, fmt.Errorf("'peer-address' has to be a valid IP: %s", prov.PeerAddress)

		}

		switch prov.Provider {
		case OpnSense:
			if prov.Key == "" {
				return nil, fmt.Errorf("'key' has to be set for OPNsense provider: %#v", prov)
			}
			if prov.Secret == "" {
				return nil, fmt.Errorf("'secret' has to be set for OPNsense provider: %#v", prov)
			}
			opn, err := NewOpnSenseProvider(prov.URL, prov.Key, prov.Secret, peerIPAddress, prov.InSecure)
			if err != nil {
				return nil, err
			}
			providers = append(providers, opn)

		case VCloud:
			if prov.User == "" {
				return nil, fmt.Errorf("'user' has to be set for OPNsense provider: %#v", prov)
			}
			if prov.Password == "" {
				return nil, fmt.Errorf("'password' has to be set for OPNsense provider: %#v", prov)
			}
			if prov.Org == "" {
				return nil, fmt.Errorf("'org' has to be set for OPNsense provider: %#v", prov)
			}
			if prov.Vdc == "" {
				return nil, fmt.Errorf("'vdc' has to be set for OPNsense provider: %#v", prov)
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
