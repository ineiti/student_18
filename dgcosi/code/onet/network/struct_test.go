package network

import (
	"testing"

	"github.com/dedis/student_18/dgcosi/code/kyber/util/key"
	"github.com/dedis/student_18/dgcosi/code/onet/log"
)

func TestServerIdentity(t *testing.T) {
	log.OutputToBuf()
	defer log.OutputToOs()
	kp1 := key.NewKeyPair(tSuite)
	kp2 := key.NewKeyPair(tSuite)

	si1 := NewServerIdentity(kp1.Public, NewLocalAddress("1"))
	si2 := NewServerIdentity(kp2.Public, NewLocalAddress("2"))

	if si1.Equal(si2) || !si1.Equal(si1) {
		t.Error("Stg's wrong with ServerIdentity")
	}

	if si1.ID.Equal(si2.ID) || !si1.ID.Equal(si1.ID) {
		t.Error("Stg's wrong with ServerIdentityID")
	}

	t1 := si1.Toml(tSuite)
	if t1.Address != si1.Address || t1.Address == "" {
		t.Error("stg wrong with Toml()")
	}

	si11 := t1.ServerIdentity(tSuite)
	if si11.Address != si1.Address || !si11.Public.Equal(si1.Public) {
		t.Error("Stg wrong with toml -> Si")
	}
	t1.Public = ""
	si12 := t1.ServerIdentity(tSuite)
	if si12.Public != nil && si12.Public.Equal(si1.Public) {
		t.Error("stg wrong with wrong toml -> wrong si")
	}

}

func TestGlobalBind(t *testing.T) {
	_, err := GlobalBind("127.0.0.1:2000")
	if err != nil {
		t.Error("Wrong with global bind")
	}
	_, err = GlobalBind("127.0.0.12000")
	if err == nil {
		t.Error("Wrong with global bind")
	}
}
