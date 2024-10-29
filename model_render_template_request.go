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

// checks if the RenderTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenderTemplateRequest{}

// RenderTemplateRequest struct for RenderTemplateRequest
type RenderTemplateRequest struct {
	TemplateName         string       `json:"template_name"`
	ObservableIds        []string     `json:"observable_ids,omitempty"`
	SearchQuery          *SearchQuery `json:"search_query,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RenderTemplateRequest RenderTemplateRequest

// NewRenderTemplateRequest instantiates a new RenderTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenderTemplateRequest(templateName string) *RenderTemplateRequest {
	this := RenderTemplateRequest{}
	this.TemplateName = templateName
	return &this
}

// NewRenderTemplateRequestWithDefaults instantiates a new RenderTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenderTemplateRequestWithDefaults() *RenderTemplateRequest {
	this := RenderTemplateRequest{}
	return &this
}

// GetTemplateName returns the TemplateName field value
func (o *RenderTemplateRequest) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *RenderTemplateRequest) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *RenderTemplateRequest) SetTemplateName(v string) {
	o.TemplateName = v
}

// GetObservableIds returns the ObservableIds field value if set, zero value otherwise.
func (o *RenderTemplateRequest) GetObservableIds() []string {
	if o == nil || IsNil(o.ObservableIds) {
		var ret []string
		return ret
	}
	return o.ObservableIds
}

// GetObservableIdsOk returns a tuple with the ObservableIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderTemplateRequest) GetObservableIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ObservableIds) {
		return nil, false
	}
	return o.ObservableIds, true
}

// HasObservableIds returns a boolean if a field has been set.
func (o *RenderTemplateRequest) HasObservableIds() bool {
	if o != nil && !IsNil(o.ObservableIds) {
		return true
	}

	return false
}

// SetObservableIds gets a reference to the given []string and assigns it to the ObservableIds field.
func (o *RenderTemplateRequest) SetObservableIds(v []string) {
	o.ObservableIds = v
}

// GetSearchQuery returns the SearchQuery field value if set, zero value otherwise.
func (o *RenderTemplateRequest) GetSearchQuery() SearchQuery {
	if o == nil || IsNil(o.SearchQuery) {
		var ret SearchQuery
		return ret
	}
	return *o.SearchQuery
}

// GetSearchQueryOk returns a tuple with the SearchQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderTemplateRequest) GetSearchQueryOk() (*SearchQuery, bool) {
	if o == nil || IsNil(o.SearchQuery) {
		return nil, false
	}
	return o.SearchQuery, true
}

// HasSearchQuery returns a boolean if a field has been set.
func (o *RenderTemplateRequest) HasSearchQuery() bool {
	if o != nil && !IsNil(o.SearchQuery) {
		return true
	}

	return false
}

// SetSearchQuery gets a reference to the given SearchQuery and assigns it to the SearchQuery field.
func (o *RenderTemplateRequest) SetSearchQuery(v SearchQuery) {
	o.SearchQuery = &v
}

func (o RenderTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenderTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template_name"] = o.TemplateName
	if !IsNil(o.ObservableIds) {
		toSerialize["observable_ids"] = o.ObservableIds
	}
	if !IsNil(o.SearchQuery) {
		toSerialize["search_query"] = o.SearchQuery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RenderTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"template_name",
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

	varRenderTemplateRequest := _RenderTemplateRequest{}

	err = json.Unmarshal(data, &varRenderTemplateRequest)

	if err != nil {
		return err
	}

	*o = RenderTemplateRequest(varRenderTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "template_name")
		delete(additionalProperties, "observable_ids")
		delete(additionalProperties, "search_query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRenderTemplateRequest struct {
	value *RenderTemplateRequest
	isSet bool
}

func (v NullableRenderTemplateRequest) Get() *RenderTemplateRequest {
	return v.value
}

func (v *NullableRenderTemplateRequest) Set(val *RenderTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRenderTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRenderTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenderTemplateRequest(val *RenderTemplateRequest) *NullableRenderTemplateRequest {
	return &NullableRenderTemplateRequest{value: val, isSet: true}
}

func (v NullableRenderTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenderTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
