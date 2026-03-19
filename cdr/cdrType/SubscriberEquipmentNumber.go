package cdrType

import "github.com/acore2026/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type SubscriberEquipmentNumber struct { /* Set Type */
	SubscriberEquipmentNumberType SubscriberEquipmentType `ber:"tagNum:0"`
	SubscriberEquipmentNumberData asn.OctetString         `ber:"tagNum:1"`
}
