package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type SMAddressDomain struct { /* Sequence Type */
	SMDomainName       *asn.GraphicString `ber:"tagNum:0,optional"`
	ThreeGPPIMSIMCCMNC *PLMNId            `ber:"tagNum:1,optional"`
}
