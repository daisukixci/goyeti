/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"encoding/json"
	"fmt"
)

// checks if the NewExtendedObservableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewExtendedObservableRequest{}

// NewExtendedObservableRequest struct for NewExtendedObservableRequest
type NewExtendedObservableRequest struct {
	Tags                 []string    `json:"tags,omitempty"`
	Observable           Observable1 `json:"observable"`
	AdditionalProperties map[string]interface{}
}

type _NewExtendedObservableRequest NewExtendedObservableRequest

// NewNewExtendedObservableRequest instantiates a new NewExtendedObservableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewExtendedObservableRequest(observable Observable1) *NewExtendedObservableRequest {
	this := NewExtendedObservableRequest{}
	this.Observable = observable
	return &this
}

// NewNewExtendedObservableRequestWithDefaults instantiates a new NewExtendedObservableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewExtendedObservableRequestWithDefaults() *NewExtendedObservableRequest {
	this := NewExtendedObservableRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NewExtendedObservableRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewExtendedObservableRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NewExtendedObservableRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *NewExtendedObservableRequest) SetTags(v []string) {
	o.Tags = v
}

// GetObservable returns the Observable field value
func (o *NewExtendedObservableRequest) GetObservable() Observable1 {
	if o == nil {
		var ret Observable1
		return ret
	}

	return o.Observable
}

// GetObservableOk returns a tuple with the Observable field value
// and a boolean to check if the value has been set.
func (o *NewExtendedObservableRequest) GetObservableOk() (*Observable1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Observable, true
}

// SetObservable sets field value
func (o *NewExtendedObservableRequest) SetObservable(v Observable1) {
	o.Observable = v
}

func (o NewExtendedObservableRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewExtendedObservableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["observable"] = o.Observable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewExtendedObservableRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"observable",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNewExtendedObservableRequest := _NewExtendedObservableRequest{}

	err = json.Unmarshal(data, &varNewExtendedObservableRequest)

	if err != nil {
		return err
	}

	*o = NewExtendedObservableRequest(varNewExtendedObservableRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		delete(additionalProperties, "observable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewExtendedObservableRequest struct {
	value *NewExtendedObservableRequest
	isSet bool
}

func (v NullableNewExtendedObservableRequest) Get() *NewExtendedObservableRequest {
	return v.value
}

func (v *NullableNewExtendedObservableRequest) Set(val *NewExtendedObservableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNewExtendedObservableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNewExtendedObservableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewExtendedObservableRequest(val *NewExtendedObservableRequest) *NullableNewExtendedObservableRequest {
	return &NullableNewExtendedObservableRequest{value: val, isSet: true}
}

func (v NullableNewExtendedObservableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewExtendedObservableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}