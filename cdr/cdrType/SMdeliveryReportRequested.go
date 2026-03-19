package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const ( /* Enum Type */
	SMdeliveryReportRequestedPresentYes asn.Enumerated = 0
	SMdeliveryReportRequestedPresentNo  asn.Enumerated = 1
)

type SMdeliveryReportRequested struct {
	Value asn.Enumerated
}
