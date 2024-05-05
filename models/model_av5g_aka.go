/*
 * AUSF API
 *
 * OpenAPI specification for AUSF
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Av5gAka struct {
	Rand      string `json:"rand" yaml:"rand" bson:"rand"`
	HxresStar string `json:"hxresStar" yaml:"hxresStar" bson:"hxresStar"`
	Autn      string `json:"autn" yaml:"autn" bson:"autn"`
	// NOTE: 增加HNMAC
	HNMAC string `json:"hnMac" yaml:"hnMac" bson:"hnMac"`
}
