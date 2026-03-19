package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type GNbId struct { /* Sequence Type */
	BitLength int64         `ber:"tagNum:0"`
	GNbValue  asn.IA5String `ber:"tagNum:1"`
}
