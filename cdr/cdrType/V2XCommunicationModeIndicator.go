package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const ( /* Enum Type */
	V2XCommunicationModeIndicatorPresentV2XComSupported    asn.Enumerated = 0
	V2XCommunicationModeIndicatorPresentV2XComNotSupported asn.Enumerated = 1
)

type V2XCommunicationModeIndicator struct {
	Value asn.Enumerated
}
