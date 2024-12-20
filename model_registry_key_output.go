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
	"os"
	"time"
)

// checks if the RegistryKeyOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistryKeyOutput{}

// RegistryKeyOutput Registry Key observable schema class.  Attributes:     key: The registry key name.     value: The registry key value.     hive: The registry hive like SYSEM, SOFTWARE, etc.     path_file: The filesystem path to the file that contains the registry key value.
type RegistryKeyOutput struct {
	Value        string                           `json:"value"`
	Type         *string                          `json:"type,omitempty"`
	Created      *time.Time                       `json:"created,omitempty"`
	Context      []map[string]interface{}         `json:"context,omitempty"`
	LastAnalysis map[string]time.Time             `json:"last_analysis,omitempty"`
	Key          string                           `json:"key"`
	Data         *os.File                         `json:"data"`
	Hive         RegistryHive                     `json:"hive"`
	PathFile     NullableString                   `json:"path_file,omitempty"`
	Id           string                           `json:"id"`
	Tags         map[string]TagRelationshipOutput `json:"tags"`
	RootType     string                           `json:"root_type"`
}

type _RegistryKeyOutput RegistryKeyOutput

// NewRegistryKeyOutput instantiates a new RegistryKeyOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryKeyOutput(value string, key string, data *os.File, hive RegistryHive, id string, tags map[string]TagRelationshipOutput, rootType string) *RegistryKeyOutput {
	this := RegistryKeyOutput{}
	this.Value = value
	var type_ string = "registry_key"
	this.Type = &type_
	this.Key = key
	this.Data = data
	this.Hive = hive
	this.Id = id
	this.Tags = tags
	this.RootType = rootType
	return &this
}

// NewRegistryKeyOutputWithDefaults instantiates a new RegistryKeyOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryKeyOutputWithDefaults() *RegistryKeyOutput {
	this := RegistryKeyOutput{}
	var type_ string = "registry_key"
	this.Type = &type_
	return &this
}

// GetValue returns the Value field value
func (o *RegistryKeyOutput) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RegistryKeyOutput) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegistryKeyOutput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegistryKeyOutput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegistryKeyOutput) SetType(v string) {
	o.Type = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RegistryKeyOutput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RegistryKeyOutput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RegistryKeyOutput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RegistryKeyOutput) GetContext() []map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetContextOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RegistryKeyOutput) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []map[string]interface{} and assigns it to the Context field.
func (o *RegistryKeyOutput) SetContext(v []map[string]interface{}) {
	o.Context = v
}

// GetLastAnalysis returns the LastAnalysis field value if set, zero value otherwise.
func (o *RegistryKeyOutput) GetLastAnalysis() map[string]time.Time {
	if o == nil || IsNil(o.LastAnalysis) {
		var ret map[string]time.Time
		return ret
	}
	return o.LastAnalysis
}

// GetLastAnalysisOk returns a tuple with the LastAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetLastAnalysisOk() (map[string]time.Time, bool) {
	if o == nil || IsNil(o.LastAnalysis) {
		return map[string]time.Time{}, false
	}
	return o.LastAnalysis, true
}

// HasLastAnalysis returns a boolean if a field has been set.
func (o *RegistryKeyOutput) HasLastAnalysis() bool {
	if o != nil && !IsNil(o.LastAnalysis) {
		return true
	}

	return false
}

// SetLastAnalysis gets a reference to the given map[string]time.Time and assigns it to the LastAnalysis field.
func (o *RegistryKeyOutput) SetLastAnalysis(v map[string]time.Time) {
	o.LastAnalysis = v
}

// GetKey returns the Key field value
func (o *RegistryKeyOutput) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *RegistryKeyOutput) SetKey(v string) {
	o.Key = v
}

// GetData returns the Data field value
func (o *RegistryKeyOutput) GetData() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetDataOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RegistryKeyOutput) SetData(v *os.File) {
	o.Data = v
}

// GetHive returns the Hive field value
func (o *RegistryKeyOutput) GetHive() RegistryHive {
	if o == nil {
		var ret RegistryHive
		return ret
	}

	return o.Hive
}

// GetHiveOk returns a tuple with the Hive field value
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetHiveOk() (*RegistryHive, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hive, true
}

// SetHive sets field value
func (o *RegistryKeyOutput) SetHive(v RegistryHive) {
	o.Hive = v
}

// GetPathFile returns the PathFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegistryKeyOutput) GetPathFile() string {
	if o == nil || IsNil(o.PathFile.Get()) {
		var ret string
		return ret
	}
	return *o.PathFile.Get()
}

// GetPathFileOk returns a tuple with the PathFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryKeyOutput) GetPathFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PathFile.Get(), o.PathFile.IsSet()
}

// HasPathFile returns a boolean if a field has been set.
func (o *RegistryKeyOutput) HasPathFile() bool {
	if o != nil && o.PathFile.IsSet() {
		return true
	}

	return false
}

// SetPathFile gets a reference to the given NullableString and assigns it to the PathFile field.
func (o *RegistryKeyOutput) SetPathFile(v string) {
	o.PathFile.Set(&v)
}

// SetPathFileNil sets the value for PathFile to be an explicit nil
func (o *RegistryKeyOutput) SetPathFileNil() {
	o.PathFile.Set(nil)
}

// UnsetPathFile ensures that no value is present for PathFile, not even an explicit nil
func (o *RegistryKeyOutput) UnsetPathFile() {
	o.PathFile.Unset()
}

// GetId returns the Id field value
func (o *RegistryKeyOutput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegistryKeyOutput) SetId(v string) {
	o.Id = v
}

// GetTags returns the Tags field value
func (o *RegistryKeyOutput) GetTags() map[string]TagRelationshipOutput {
	if o == nil {
		var ret map[string]TagRelationshipOutput
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetTagsOk() (map[string]TagRelationshipOutput, bool) {
	if o == nil {
		return map[string]TagRelationshipOutput{}, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *RegistryKeyOutput) SetTags(v map[string]TagRelationshipOutput) {
	o.Tags = v
}

// GetRootType returns the RootType field value
func (o *RegistryKeyOutput) GetRootType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootType
}

// GetRootTypeOk returns a tuple with the RootType field value
// and a boolean to check if the value has been set.
func (o *RegistryKeyOutput) GetRootTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootType, true
}

// SetRootType sets field value
func (o *RegistryKeyOutput) SetRootType(v string) {
	o.RootType = v
}

func (o RegistryKeyOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistryKeyOutput) ToMap() (map[string]interface{}, error) {
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
	toSerialize["key"] = o.Key
	toSerialize["data"] = o.Data
	toSerialize["hive"] = o.Hive
	if o.PathFile.IsSet() {
		toSerialize["path_file"] = o.PathFile.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["tags"] = o.Tags
	toSerialize["root_type"] = o.RootType
	return toSerialize, nil
}

func (o *RegistryKeyOutput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"key",
		"data",
		"hive",
		"id",
		"tags",
		"root_type",
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

	varRegistryKeyOutput := _RegistryKeyOutput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistryKeyOutput)

	if err != nil {
		return err
	}

	*o = RegistryKeyOutput(varRegistryKeyOutput)

	return err
}

type NullableRegistryKeyOutput struct {
	value *RegistryKeyOutput
	isSet bool
}

func (v NullableRegistryKeyOutput) Get() *RegistryKeyOutput {
	return v.value
}

func (v *NullableRegistryKeyOutput) Set(val *RegistryKeyOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryKeyOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryKeyOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryKeyOutput(val *RegistryKeyOutput) *NullableRegistryKeyOutput {
	return &NullableRegistryKeyOutput{value: val, isSet: true}
}

func (v NullableRegistryKeyOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryKeyOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
