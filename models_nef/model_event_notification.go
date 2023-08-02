/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models_nef

import (
	"github.com/free5gc/openapi/models"
)

type EventNotification struct {
	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransId string `json:"afTransId,omitempty" bson:"afTransId"`

	DnaiChgType models.DnaiChangeType `json:"dnaiChgType" bson:"dnaiChgType"`

	SourceTrafficRoute *models.RouteToLocation `json:"sourceTrafficRoute,omitempty" bson:"sourceTrafficRoute"`

	SubscribedEvent models.SubscribedEvent `json:"subscribedEvent" bson:"subscribedEvent"`

	TargetTrafficRoute *models.RouteToLocation `json:"targetTrafficRoute,omitempty" bson:"targetTrafficRoute"`

	SourceDnai string `json:"sourceDnai,omitempty" bson:"sourceDnai"`

	TargetDnai string `json:"targetDnai,omitempty" bson:"targetDnai"`

	Gpsi string `json:"gpsi,omitempty" bson:"gpsi"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	SrcUeIpv4Addr string `json:"srcUeIpv4Addr,omitempty" bson:"srcUeIpv4Addr"`

	SrcUeIpv6Prefix string `json:"srcUeIpv6Prefix,omitempty" bson:"srcUeIpv6Prefix"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	TgtUeIpv4Addr string `json:"tgtUeIpv4Addr,omitempty" bson:"tgtUeIpv4Addr"`

	TgtUeIpv6Prefix string `json:"tgtUeIpv6Prefix,omitempty" bson:"tgtUeIpv6Prefix"`

	UeMac string `json:"ueMac,omitempty" bson:"ueMac"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	AfAckUri string `json:"afAckUri,omitempty" bson:"afAckUri"`
}