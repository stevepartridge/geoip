package geoip_test

import (
	"testing"

	"github.com/stevepartridge/geoip"
)

const (
	IPTestIP1 = "208.80.152.201"
)

func TestGeoipInit(t *testing.T) {
	client := geoip.New()
	if client == nil {
		t.Error("Failed to init ip client.")
	}
}

func TestGeoipLookup(t *testing.T) {
	client := geoip.New()
	if client == nil {
		t.Error("Failed to init ip client.")
	}

	ip, err := client.Lookup(IPTestIP1)
	if err != nil {
		t.Error("Error looking up ip.")
		t.Log(err.Error())
	}

	if ip.IP != IPTestIP1 {
		t.Error("IP look up did not match " + IPTestIP1)
	}
}
