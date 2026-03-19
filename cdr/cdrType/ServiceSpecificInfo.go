package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type ServiceSpecificInfo struct { /* Sequence Type */
	ServiceSpecificData *asn.GraphicString `ber:"tagNum:0,optional"`
	ServiceSpecificType *int64             `ber:"tagNum:1,optional"`
}
