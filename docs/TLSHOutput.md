# TLSHOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "tlsh"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 

## Methods

### NewTLSHOutput

`func NewTLSHOutput(value string, id string, tags map[string]TagRelationshipOutput, rootType string, ) *TLSHOutput`

NewTLSHOutput instantiates a new TLSHOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSHOutputWithDefaults

`func NewTLSHOutputWithDefaults() *TLSHOutput`

NewTLSHOutputWithDefaults instantiates a new TLSHOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *TLSHOutput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TLSHOutput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TLSHOutput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *TLSHOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSHOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSHOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TLSHOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *TLSHOutput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TLSHOutput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TLSHOutput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TLSHOutput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *TLSHOutput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TLSHOutput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TLSHOutput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *TLSHOutput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *TLSHOutput) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *TLSHOutput) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *TLSHOutput) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *TLSHOutput) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetId

`func (o *TLSHOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TLSHOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TLSHOutput) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *TLSHOutput) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TLSHOutput) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TLSHOutput) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *TLSHOutput) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *TLSHOutput) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *TLSHOutput) SetRootType(v string)`

SetRootType sets RootType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

