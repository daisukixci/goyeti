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

// checks if the PatchExportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchExportRequest{}

// PatchExportRequest struct for PatchExportRequest
type PatchExportRequest struct {
	Export               ExportTaskInput `json:"export"`
	AdditionalProperties map[string]interface{}
}

type _PatchExportRequest PatchExportRequest

// NewPatchExportRequest instantiates a new PatchExportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchExportRequest(export ExportTaskInput) *PatchExportRequest {
	this := PatchExportRequest{}
	this.Export = export
	return &this
}

// NewPatchExportRequestWithDefaults instantiates a new PatchExportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchExportRequestWithDefaults() *PatchExportRequest {
	this := PatchExportRequest{}
	return &this
}

// GetExport returns the Export field value
func (o *PatchExportRequest) GetExport() ExportTaskInput {
	if o == nil {
		var ret ExportTaskInput
		return ret
	}

	return o.Export
}

// GetExportOk returns a tuple with the Export field value
// and a boolean to check if the value has been set.
func (o *PatchExportRequest) GetExportOk() (*ExportTaskInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Export, true
}

// SetExport sets field value
func (o *PatchExportRequest) SetExport(v ExportTaskInput) {
	o.Export = v
}

func (o PatchExportRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchExportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["export"] = o.Export

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchExportRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"export",
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

	varPatchExportRequest := _PatchExportRequest{}

	err = json.Unmarshal(data, &varPatchExportRequest)

	if err != nil {
		return err
	}

	*o = PatchExportRequest(varPatchExportRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "export")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchExportRequest struct {
	value *PatchExportRequest
	isSet bool
}

func (v NullablePatchExportRequest) Get() *PatchExportRequest {
	return v.value
}

func (v *NullablePatchExportRequest) Set(val *PatchExportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchExportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchExportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchExportRequest(val *PatchExportRequest) *NullablePatchExportRequest {
	return &NullablePatchExportRequest{value: val, isSet: true}
}

func (v NullablePatchExportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchExportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}