/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MessageIn struct for MessageIn
type MessageIn struct {
	// List of free-form identifiers that endpoints can filter by
	Channels []string `json:"channels,omitempty"`
	// Optional unique identifier for the message
	EventId NullableString `json:"eventId,omitempty"`
	EventType string `json:"eventType"`
	Payload map[string]interface{} `json:"payload"`
	// The retention period for the payload (in days).
	PayloadRetentionPeriod *int32 `json:"payloadRetentionPeriod,omitempty"`
}

// NewMessageIn instantiates a new MessageIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageIn(eventType string, payload map[string]interface{}) *MessageIn {
	this := MessageIn{}
	this.EventType = eventType
	this.Payload = payload
	var payloadRetentionPeriod int32 = 90
	this.PayloadRetentionPeriod = &payloadRetentionPeriod
	return &this
}

// NewMessageInWithDefaults instantiates a new MessageIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageInWithDefaults() *MessageIn {
	this := MessageIn{}
	var payloadRetentionPeriod int32 = 90
	this.PayloadRetentionPeriod = &payloadRetentionPeriod
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageIn) GetChannels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageIn) GetChannelsOk() (*[]string, bool) {
	if o == nil || o.Channels == nil {
		return nil, false
	}
	return &o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *MessageIn) HasChannels() bool {
	if o != nil && o.Channels != nil {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *MessageIn) SetChannels(v []string) {
	o.Channels = v
}

// GetEventId returns the EventId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageIn) GetEventId() string {
	if o == nil || o.EventId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EventId.Get()
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageIn) GetEventIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EventId.Get(), o.EventId.IsSet()
}

// HasEventId returns a boolean if a field has been set.
func (o *MessageIn) HasEventId() bool {
	if o != nil && o.EventId.IsSet() {
		return true
	}

	return false
}

// SetEventId gets a reference to the given NullableString and assigns it to the EventId field.
func (o *MessageIn) SetEventId(v string) {
	o.EventId.Set(&v)
}
// SetEventIdNil sets the value for EventId to be an explicit nil
func (o *MessageIn) SetEventIdNil() {
	o.EventId.Set(nil)
}

// UnsetEventId ensures that no value is present for EventId, not even an explicit nil
func (o *MessageIn) UnsetEventId() {
	o.EventId.Unset()
}

// GetEventType returns the EventType field value
func (o *MessageIn) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MessageIn) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MessageIn) SetEventType(v string) {
	o.EventType = v
}

// GetPayload returns the Payload field value
func (o *MessageIn) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *MessageIn) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *MessageIn) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetPayloadRetentionPeriod returns the PayloadRetentionPeriod field value if set, zero value otherwise.
func (o *MessageIn) GetPayloadRetentionPeriod() int32 {
	if o == nil || o.PayloadRetentionPeriod == nil {
		var ret int32
		return ret
	}
	return *o.PayloadRetentionPeriod
}

// GetPayloadRetentionPeriodOk returns a tuple with the PayloadRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageIn) GetPayloadRetentionPeriodOk() (*int32, bool) {
	if o == nil || o.PayloadRetentionPeriod == nil {
		return nil, false
	}
	return o.PayloadRetentionPeriod, true
}

// HasPayloadRetentionPeriod returns a boolean if a field has been set.
func (o *MessageIn) HasPayloadRetentionPeriod() bool {
	if o != nil && o.PayloadRetentionPeriod != nil {
		return true
	}

	return false
}

// SetPayloadRetentionPeriod gets a reference to the given int32 and assigns it to the PayloadRetentionPeriod field.
func (o *MessageIn) SetPayloadRetentionPeriod(v int32) {
	o.PayloadRetentionPeriod = &v
}

func (o MessageIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if o.EventId.IsSet() {
		toSerialize["eventId"] = o.EventId.Get()
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if o.PayloadRetentionPeriod != nil {
		toSerialize["payloadRetentionPeriod"] = o.PayloadRetentionPeriod
	}
	return json.Marshal(toSerialize)
}

type NullableMessageIn struct {
	value *MessageIn
	isSet bool
}

func (v NullableMessageIn) Get() *MessageIn {
	return v.value
}

func (v *NullableMessageIn) Set(val *MessageIn) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageIn) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageIn(val *MessageIn) *NullableMessageIn {
	return &NullableMessageIn{value: val, isSet: true}
}

func (v NullableMessageIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


