package service_test

import (

	// We need to include the service so it is started.
	"gopkg.in/dedis/kyber.v2/suites"
)

var tSuite = suites.MustFind("Ed25519")
