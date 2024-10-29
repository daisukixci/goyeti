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

// checks if the NewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewRequest{}

// NewRequest struct for NewRequest
type NewRequest struct {
	Name                 string   `json:"name"`
	DefaultExpiration    *string  `json:"default_expiration,omitempty"`
	Produces             []string `json:"produces,omitempty"`
	Replaces             []string `json:"replaces,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NewRequest NewRequest

// NewNewRequest instantiates a new NewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewRequest(name string) *NewRequest {
	this := NewRequest{}
	this.Name = name
	var defaultExpiration string = "P90D"
	this.DefaultExpiration = &defaultExpiration
	return &this
}

// NewNewRequestWithDefaults instantiates a new NewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewRequestWithDefaults() *NewRequest {
	this := NewRequest{}
	var defaultExpiration string = "P90D"
	this.DefaultExpiration = &defaultExpiration
	return &this
}

// GetName returns the Name field value
func (o *NewRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewRequest) SetName(v string) {
	o.Name = v
}

// GetDefaultExpiration returns the DefaultExpiration field value if set, zero value otherwise.
func (o *NewRequest) GetDefaultExpiration() string {
	if o == nil || IsNil(o.DefaultExpiration) {
		var ret string
		return ret
	}
	return *o.DefaultExpiration
}

// GetDefaultExpirationOk returns a tuple with the DefaultExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRequest) GetDefaultExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultExpiration) {
		return nil, false
	}
	return o.DefaultExpiration, true
}

// HasDefaultExpiration returns a boolean if a field has been set.
func (o *NewRequest) HasDefaultExpiration() bool {
	if o != nil && !IsNil(o.DefaultExpiration) {
		return true
	}

	return false
}

// SetDefaultExpiration gets a reference to the given string and assigns it to the DefaultExpiration field.
func (o *NewRequest) SetDefaultExpiration(v string) {
	o.DefaultExpiration = &v
}

// GetProduces returns the Produces field value if set, zero value otherwise.
func (o *NewRequest) GetProduces() []string {
	if o == nil || IsNil(o.Produces) {
		var ret []string
		return ret
	}
	return o.Produces
}

// GetProducesOk returns a tuple with the Produces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRequest) GetProducesOk() ([]string, bool) {
	if o == nil || IsNil(o.Produces) {
		return nil, false
	}
	return o.Produces, true
}

// HasProduces returns a boolean if a field has been set.
func (o *NewRequest) HasProduces() bool {
	if o != nil && !IsNil(o.Produces) {
		return true
	}

	return false
}

// SetProduces gets a reference to the given []string and assigns it to the Produces field.
func (o *NewRequest) SetProduces(v []string) {
	o.Produces = v
}

// GetReplaces returns the Replaces field value if set, zero value otherwise.
func (o *NewRequest) GetReplaces() []string {
	if o == nil || IsNil(o.Replaces) {
		var ret []string
		return ret
	}
	return o.Replaces
}

// GetReplacesOk returns a tuple with the Replaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewRequest) GetReplacesOk() ([]string, bool) {
	if o == nil || IsNil(o.Replaces) {
		return nil, false
	}
	return o.Replaces, true
}

// HasReplaces returns a boolean if a field has been set.
func (o *NewRequest) HasReplaces() bool {
	if o != nil && !IsNil(o.Replaces) {
		return true
	}

	return false
}

// SetReplaces gets a reference to the given []string and assigns it to the Replaces field.
func (o *NewRequest) SetReplaces(v []string) {
	o.Replaces = v
}

func (o NewRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.DefaultExpiration) {
		toSerialize["default_expiration"] = o.DefaultExpiration
	}
	if !IsNil(o.Produces) {
		toSerialize["produces"] = o.Produces
	}
	if !IsNil(o.Replaces) {
		toSerialize["replaces"] = o.Replaces
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varNewRequest := _NewRequest{}

	err = json.Unmarshal(data, &varNewRequest)

	if err != nil {
		return err
	}

	*o = NewRequest(varNewRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "default_expiration")
		delete(additionalProperties, "produces")
		delete(additionalProperties, "replaces")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewRequest struct {
	value *NewRequest
	isSet bool
}

func (v NullableNewRequest) Get() *NewRequest {
	return v.value
}

func (v *NullableNewRequest) Set(val *NewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewRequest(val *NewRequest) *NullableNewRequest {
	return &NullableNewRequest{value: val, isSet: true}
}

func (v NullableNewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}