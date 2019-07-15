package main

import (
	"log"
	"net"
	"strconv"

	"github.com/kradalby/opnsense-go/opnsense"
)

type BgpProvider interface {
	Add(net.IP, uint32) error
	Delete(net.IP) error
	Name() string
	URL() string
	PeerIP() net.IP
}

type OpnSenseProvider struct {
	APIURL        string
	Key           string
	Secret        string
	PeerIPAddress net.IP
	InSecure      bool
	c             *opnsense.Client
}

func NewOpnSenseProvider(url, key, secret string, peerIPAddress net.IP, inSecure bool) (*OpnSenseProvider, error) {
	opn := &OpnSenseProvider{
		APIURL:        url,
		Key:           key,
		Secret:        secret,
		PeerIPAddress: peerIPAddress,
		InSecure:      inSecure,
	}

	client, err := opnsense.NewClient(opn.APIURL, opn.Key, opn.Secret, opn.InSecure)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Init OPNsense client: %#v", client)
	opn.c = client

	return opn, err
}

func (opn OpnSenseProvider) Add(ip net.IP, as uint32) error {

	neighbours, err := opn.c.BgpNeighborList()
	if err != nil {
		return err
	}

	for _, neighbour := range neighbours {
		if ip.Equal(net.ParseIP(neighbour.Address)) &&
			neighbour.Remoteas == strconv.Itoa(int(as)) {
			return nil
		}
	}

	newNeighbour := opnsense.BgpNeighborSet{}
	newNeighbour.Enabled = "1"
	newNeighbour.Address = ip.String()
	newNeighbour.Remoteas = strconv.Itoa(int(as))
	// newNeighbour.Nexthopself = "0"
	// newNeighbour.Defaultoriginate = "0"
	newNeighbour.Updatesource = "wan"
	// newNeighbour.LinkedPrefixlistIn = ""
	// newNeighbour.LinkedPrefixlistOut = ""
	// newNeighbour.LinkedRoutemapIn = ""
	// newNeighbour.LinkedRoutemapOut = ""

	log.Printf("[INFO] Adding neighbour: %s with AS number: %d", ip.String(), as)
	_, err = opn.c.BgpNeighborAdd(newNeighbour)
	if err != nil {
		return err
	}

	return nil
}

func (opn OpnSenseProvider) Delete(ip net.IP) error {

	neighbours, err := opn.c.BgpNeighborList()
	if err != nil {
		return err
	}

	for _, neighbour := range neighbours {
		if ip.Equal(net.ParseIP(neighbour.Address)) {
			log.Printf("[INFO] Removing neighbour %s with IP %s", neighbour.UUID.String(), ip.String())

			_, err := opn.c.BgpNeighborDelete(*neighbour.UUID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (opn OpnSenseProvider) Name() string {
	return opn.APIURL
}

func (opn OpnSenseProvider) URL() string {
	return opn.APIURL
}

func (opn OpnSenseProvider) PeerIP() net.IP {
	return opn.PeerIPAddress
}
