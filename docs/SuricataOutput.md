# SuricataOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "suricata"]
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**ValidFrom** | Pointer to **time.Time** |  | [optional] 
**ValidUntil** | Pointer to **time.Time** |  | [optional] 
**Pattern** | **string** |  | 
**Location** | Pointer to **string** |  | [optional] [default to ""]
**Diamond** | [**DiamondModel**](DiamondModel.md) |  | 
**KillChainPhases** | Pointer to **[]string** |  | [optional] [default to []]
**RelevantTags** | Pointer to **[]string** |  | [optional] [default to []]
**Sid** | Pointer to **int32** |  | [optional] [default to 0]
**Metadata** | Pointer to **[]string** |  | [optional] [default to []]
**References** | Pointer to **[]string** |  | [optional] [default to []]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 

## Methods

### NewSuricataOutput

`func NewSuricataOutput(name string, pattern string, diamond DiamondModel, id string, tags map[string]TagRelationshipOutput, rootType string, ) *SuricataOutput`

NewSuricataOutput instantiates a new SuricataOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuricataOutputWithDefaults

`func NewSuricataOutputWithDefaults() *SuricataOutput`

NewSuricataOutputWithDefaults instantiates a new SuricataOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SuricataOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuricataOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuricataOutput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SuricataOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SuricataOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SuricataOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SuricataOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *SuricataOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SuricataOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SuricataOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SuricataOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *SuricataOutput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SuricataOutput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SuricataOutput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SuricataOutput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *SuricataOutput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SuricataOutput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SuricataOutput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SuricataOutput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetValidFrom

`func (o *SuricataOutput) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *SuricataOutput) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *SuricataOutput) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *SuricataOutput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *SuricataOutput) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *SuricataOutput) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *SuricataOutput) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *SuricataOutput) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPattern

`func (o *SuricataOutput) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SuricataOutput) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SuricataOutput) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetLocation

`func (o *SuricataOutput) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SuricataOutput) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SuricataOutput) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SuricataOutput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDiamond

`func (o *SuricataOutput) GetDiamond() DiamondModel`

GetDiamond returns the Diamond field if non-nil, zero value otherwise.

### GetDiamondOk

`func (o *SuricataOutput) GetDiamondOk() (*DiamondModel, bool)`

GetDiamondOk returns a tuple with the Diamond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiamond

`func (o *SuricataOutput) SetDiamond(v DiamondModel)`

SetDiamond sets Diamond field to given value.


### GetKillChainPhases

`func (o *SuricataOutput) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *SuricataOutput) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *SuricataOutput) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *SuricataOutput) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetRelevantTags

`func (o *SuricataOutput) GetRelevantTags() []string`

GetRelevantTags returns the RelevantTags field if non-nil, zero value otherwise.

### GetRelevantTagsOk

`func (o *SuricataOutput) GetRelevantTagsOk() (*[]string, bool)`

GetRelevantTagsOk returns a tuple with the RelevantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTags

`func (o *SuricataOutput) SetRelevantTags(v []string)`

SetRelevantTags sets RelevantTags field to given value.

### HasRelevantTags

`func (o *SuricataOutput) HasRelevantTags() bool`

HasRelevantTags returns a boolean if a field has been set.

### GetSid

`func (o *SuricataOutput) GetSid() int32`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SuricataOutput) GetSidOk() (*int32, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SuricataOutput) SetSid(v int32)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SuricataOutput) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetMetadata

`func (o *SuricataOutput) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SuricataOutput) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SuricataOutput) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SuricataOutput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *SuricataOutput) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *SuricataOutput) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *SuricataOutput) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *SuricataOutput) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetId

`func (o *SuricataOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuricataOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuricataOutput) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *SuricataOutput) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SuricataOutput) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SuricataOutput) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *SuricataOutput) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *SuricataOutput) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *SuricataOutput) SetRootType(v string)`

SetRootType sets RootType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

