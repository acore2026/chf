package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const ( /* Enum Type */
	CoreNetworkTypePresentFiveGC asn.Enumerated = 0
	CoreNetworkTypePresentEPC    asn.Enumerated = 1
)

type CoreNetworkType struct {
	Value asn.Enumerated
}
