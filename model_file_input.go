/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the FileInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileInput{}

// FileInput Represents a file.  One of sha256, md5, or sha1 should be provided. Value should to be in the form FILE:<HASH>.
type FileInput struct {
	Value        string                   `json:"value"`
	Type         *string                  `json:"type,omitempty"`
	Created      *time.Time               `json:"created,omitempty"`
	Context      []map[string]interface{} `json:"context,omitempty"`
	LastAnalysis map[string]time.Time     `json:"last_analysis,omitempty"`
	Name         NullableString           `json:"name,omitempty"`
	Size         NullableInt32            `json:"size,omitempty"`
	Sha256       NullableString           `json:"sha256,omitempty"`
	Md5          NullableString           `json:"md5,omitempty"`
	Sha1         NullableString           `json:"sha1,omitempty"`
	MimeType     NullableString           `json:"mime_type,omitempty"`
}

type _FileInput FileInput

// NewFileInput instantiates a new FileInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileInput(value string) *FileInput {
	this := FileInput{}
	this.Value = value
	var type_ string = "file"
	this.Type = &type_
	return &this
}

// NewFileInputWithDefaults instantiates a new FileInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileInputWithDefaults() *FileInput {
	this := FileInput{}
	var type_ string = "file"
	this.Type = &type_
	return &this
}

// GetValue returns the Value field value
func (o *FileInput) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FileInput) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FileInput) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FileInput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FileInput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FileInput) SetType(v string) {
	o.Type = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FileInput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FileInput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *FileInput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *FileInput) GetContext() []map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInput) GetContextOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *FileInput) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []map[string]interface{} and assigns it to the Context field.
func (o *FileInput) SetContext(v []map[string]interface{}) {
	o.Context = v
}

// GetLastAnalysis returns the LastAnalysis field value if set, zero value otherwise.
func (o *FileInput) GetLastAnalysis() map[string]time.Time {
	if o == nil || IsNil(o.LastAnalysis) {
		var ret map[string]time.Time
		return ret
	}
	return o.LastAnalysis
}

// GetLastAnalysisOk returns a tuple with the LastAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileInput) GetLastAnalysisOk() (map[string]time.Time, bool) {
	if o == nil || IsNil(o.LastAnalysis) {
		return map[string]time.Time{}, false
	}
	return o.LastAnalysis, true
}

// HasLastAnalysis returns a boolean if a field has been set.
func (o *FileInput) HasLastAnalysis() bool {
	if o != nil && !IsNil(o.LastAnalysis) {
		return true
	}

	return false
}

// SetLastAnalysis gets a reference to the given map[string]time.Time and assigns it to the LastAnalysis field.
func (o *FileInput) SetLastAnalysis(v map[string]time.Time) {
	o.LastAnalysis = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileInput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FileInput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FileInput) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *FileInput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FileInput) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileInput) GetSize() int32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInput) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *FileInput) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *FileInput) SetSize(v int32) {
	o.Size.Set(&v)
}

// SetSizeNil sets the value for Size to be an explicit nil
func (o *FileInput) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *FileInput) UnsetSize() {
	o.Size.Unset()
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileInput) GetSha256() string {
	if o == nil || IsNil(o.Sha256.Get()) {
		var ret string
		return ret
	}
	return *o.Sha256.Get()
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInput) GetSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha256.Get(), o.Sha256.IsSet()
}

// HasSha256 returns a boolean if a field has been set.
func (o *FileInput) HasSha256() bool {
	if o != nil && o.Sha256.IsSet() {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given NullableString and assigns it to the Sha256 field.
func (o *FileInput) SetSha256(v string) {
	o.Sha256.Set(&v)
}

// SetSha256Nil sets the value for Sha256 to be an explicit nil
func (o *FileInput) SetSha256Nil() {
	o.Sha256.Set(nil)
}

// UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
func (o *FileInput) UnsetSha256() {
	o.Sha256.Unset()
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileInput) GetMd5() string {
	if o == nil || IsNil(o.Md5.Get()) {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInput) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *FileInput) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *FileInput) SetMd5(v string) {
	o.Md5.Set(&v)
}

// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *FileInput) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *FileInput) UnsetMd5() {
	o.Md5.Unset()
}

// GetSha1 returns the Sha1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileInput) GetSha1() string {
	if o == nil || IsNil(o.Sha1.Get()) {
		var ret string
		return ret
	}
	return *o.Sha1.Get()
}

// GetSha1Ok returns a tuple with the Sha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInput) GetSha1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha1.Get(), o.Sha1.IsSet()
}

// HasSha1 returns a boolean if a field has been set.
func (o *FileInput) HasSha1() bool {
	if o != nil && o.Sha1.IsSet() {
		return true
	}

	return false
}

// SetSha1 gets a reference to the given NullableString and assigns it to the Sha1 field.
func (o *FileInput) SetSha1(v string) {
	o.Sha1.Set(&v)
}

// SetSha1Nil sets the value for Sha1 to be an explicit nil
func (o *FileInput) SetSha1Nil() {
	o.Sha1.Set(nil)
}

// UnsetSha1 ensures that no value is present for Sha1, not even an explicit nil
func (o *FileInput) UnsetSha1() {
	o.Sha1.Unset()
}

// GetMimeType returns the MimeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileInput) GetMimeType() string {
	if o == nil || IsNil(o.MimeType.Get()) {
		var ret string
		return ret
	}
	return *o.MimeType.Get()
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileInput) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MimeType.Get(), o.MimeType.IsSet()
}

// HasMimeType returns a boolean if a field has been set.
func (o *FileInput) HasMimeType() bool {
	if o != nil && o.MimeType.IsSet() {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given NullableString and assigns it to the MimeType field.
func (o *FileInput) SetMimeType(v string) {
	o.MimeType.Set(&v)
}

// SetMimeTypeNil sets the value for MimeType to be an explicit nil
func (o *FileInput) SetMimeTypeNil() {
	o.MimeType.Set(nil)
}

// UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
func (o *FileInput) UnsetMimeType() {
	o.MimeType.Unset()
}

func (o FileInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.LastAnalysis) {
		toSerialize["last_analysis"] = o.LastAnalysis
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.Sha256.IsSet() {
		toSerialize["sha256"] = o.Sha256.Get()
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if o.Sha1.IsSet() {
		toSerialize["sha1"] = o.Sha1.Get()
	}
	if o.MimeType.IsSet() {
		toSerialize["mime_type"] = o.MimeType.Get()
	}
	return toSerialize, nil
}

func (o *FileInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varFileInput := _FileInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileInput)

	if err != nil {
		return err
	}

	*o = FileInput(varFileInput)

	return err
}

type NullableFileInput struct {
	value *FileInput
	isSet bool
}

func (v NullableFileInput) Get() *FileInput {
	return v.value
}

func (v *NullableFileInput) Set(val *FileInput) {
	v.value = val
	v.isSet = true
}

func (v NullableFileInput) IsSet() bool {
	return v.isSet
}

func (v *NullableFileInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileInput(val *FileInput) *NullableFileInput {
	return &NullableFileInput{value: val, isSet: true}
}

func (v NullableFileInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}