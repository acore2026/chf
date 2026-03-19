package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type SMInterface struct { /* Sequence Type */
	InterfaceId   *asn.GraphicString `ber:"tagNum:0,optional"`
	InterfaceText *asn.GraphicString `ber:"tagNum:1,optional"`
	InterfacePort *asn.GraphicString `ber:"tagNum:2,optional"`
	InterfaceType *SMInterfaceType   `ber:"tagNum:3,optional"`
}
