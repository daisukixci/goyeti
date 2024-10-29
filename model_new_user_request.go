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

// checks if the NewUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewUserRequest{}

// NewUserRequest struct for NewUserRequest
type NewUserRequest struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	Admin                bool   `json:"admin"`
	AdditionalProperties map[string]interface{}
}

type _NewUserRequest NewUserRequest

// NewNewUserRequest instantiates a new NewUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUserRequest(username string, password string, admin bool) *NewUserRequest {
	this := NewUserRequest{}
	this.Username = username
	this.Password = password
	this.Admin = admin
	return &this
}

// NewNewUserRequestWithDefaults instantiates a new NewUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUserRequestWithDefaults() *NewUserRequest {
	this := NewUserRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *NewUserRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *NewUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *NewUserRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *NewUserRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *NewUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *NewUserRequest) SetPassword(v string) {
	o.Password = v
}

// GetAdmin returns the Admin field value
func (o *NewUserRequest) GetAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Admin
}

// GetAdminOk returns a tuple with the Admin field value
// and a boolean to check if the value has been set.
func (o *NewUserRequest) GetAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Admin, true
}

// SetAdmin sets field value
func (o *NewUserRequest) SetAdmin(v bool) {
	o.Admin = v
}

func (o NewUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["admin"] = o.Admin

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
		"admin",
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

	varNewUserRequest := _NewUserRequest{}

	err = json.Unmarshal(data, &varNewUserRequest)

	if err != nil {
		return err
	}

	*o = NewUserRequest(varNewUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "admin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewUserRequest struct {
	value *NewUserRequest
	isSet bool
}

func (v NullableNewUserRequest) Get() *NewUserRequest {
	return v.value
}

func (v *NullableNewUserRequest) Set(val *NewUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUserRequest(val *NewUserRequest) *NullableNewUserRequest {
	return &NullableNewUserRequest{value: val, isSet: true}
}

func (v NullableNewUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
