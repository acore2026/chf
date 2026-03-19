package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const ( /* Enum Type */
	LineTypePresentDSL asn.Enumerated = 0
	LineTypePresentPON asn.Enumerated = 1
)

type LineType struct {
	Value asn.Enumerated
}
