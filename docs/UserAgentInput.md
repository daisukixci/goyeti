# UserAgentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "user_agent"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]

## Methods

### NewUserAgentInput

`func NewUserAgentInput(value string, ) *UserAgentInput`

NewUserAgentInput instantiates a new UserAgentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAgentInputWithDefaults

`func NewUserAgentInputWithDefaults() *UserAgentInput`

NewUserAgentInputWithDefaults instantiates a new UserAgentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UserAgentInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserAgentInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserAgentInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *UserAgentInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAgentInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAgentInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserAgentInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *UserAgentInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserAgentInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserAgentInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserAgentInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *UserAgentInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserAgentInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserAgentInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *UserAgentInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *UserAgentInput) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *UserAgentInput) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *UserAgentInput) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *UserAgentInput) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

