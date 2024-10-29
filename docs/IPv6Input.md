# IPv6Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "ipv6"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]

## Methods

### NewIPv6Input

`func NewIPv6Input(value string, ) *IPv6Input`

NewIPv6Input instantiates a new IPv6Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv6InputWithDefaults

`func NewIPv6InputWithDefaults() *IPv6Input`

NewIPv6InputWithDefaults instantiates a new IPv6Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *IPv6Input) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IPv6Input) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IPv6Input) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *IPv6Input) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPv6Input) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPv6Input) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPv6Input) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *IPv6Input) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IPv6Input) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IPv6Input) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IPv6Input) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *IPv6Input) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *IPv6Input) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *IPv6Input) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *IPv6Input) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *IPv6Input) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *IPv6Input) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *IPv6Input) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *IPv6Input) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

