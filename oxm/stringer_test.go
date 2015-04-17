package oxm

import (
	"bytes"
	"testing"
)

func TestStrings(t *testing.T) {
	tokens := []string{
		"in_port=10",
		"in_phy_port=10",
		"metadata=0x5/0xff",
		"ipv4_src=192.168.0.1",
		"ipv4_src=192.168.0.1/255.255.255.0",
		"ipv4_src=192.168.0.1/255.0.255.255",
	}
	for _, token := range tokens {
		if o, n, err := ParseOne(token); err != nil {
			t.Error(err)
		} else if n != len(token) {
			t.Errorf("consumed length error %d for %s len=%d", n, token, len(token))
		} else if token != Oxm(o).String() {
			t.Errorf("stringer %s != %s", token, Oxm(o).String())
		}
	}
}

func TestToOxm(t *testing.T) {
	token := "ipv4_src=192.168.0.1/24"
	suffix := []byte{192, 168, 0, 1, 255, 255, 255, 0}
	if o, n, err := ParseOne(token); err != nil {
		t.Error(err)
	} else if n != len(token) {
		t.Errorf("consumed length error %d for %s len=%d", n, token, len(token))
	} else if !bytes.Equal(o[len(o)-len(suffix):], suffix) {
		t.Errorf("stringer suffix error %v %v", o[len(o)-len(suffix):], suffix)
	}
}