/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NsmfEventExposureNotification struct {
	// Notification correlation ID
	NotifId string `json:"notifId" yaml:"notifId" bson:"notifId" mapstructure:"NotifId"`
	// Notifications about Individual Events
	EventNotifs []EventNotification `json:"eventNotifs" yaml:"eventNotifs" bson:"eventNotifs" mapstructure:"EventNotifs"`
}