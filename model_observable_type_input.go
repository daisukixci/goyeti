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

// ObservableTypeInput the model 'ObservableTypeInput'
type ObservableTypeInput string

// List of ObservableType-Input
const (
	CIDR         ObservableTypeInput = "cidr"
	HOSTNAME     ObservableTypeInput = "hostname"
	MAC_ADDRESS  ObservableTypeInput = "mac_address"
	CERTIFICATE  ObservableTypeInput = "certificate"
	IPV6         ObservableTypeInput = "ipv6"
	USER_ACCOUNT ObservableTypeInput = "user_account"
	SSDEEP       ObservableTypeInput = "ssdeep"
	BIC          ObservableTypeInput = "bic"
	PACKAGE      ObservableTypeInput = "package"
	SHA256       ObservableTypeInput = "sha256"
	DOCKER_IMAGE ObservableTypeInput = "docker_image"
	EMAIL        ObservableTypeInput = "email"
	JARM         ObservableTypeInput = "jarm"
	FILE         ObservableTypeInput = "file"
	MUTEX        ObservableTypeInput = "mutex"
	IBAN         ObservableTypeInput = "iban"
	COMMAND_LINE ObservableTypeInput = "command_line"
	URL          ObservableTypeInput = "url"
	REGISTRY_KEY ObservableTypeInput = "registry_key"
	PATH         ObservableTypeInput = "path"
	IPV4         ObservableTypeInput = "ipv4"
	TLSH         ObservableTypeInput = "tlsh"
	GENERIC      ObservableTypeInput = "generic"
	SHA1         ObservableTypeInput = "sha1"
	USER_AGENT   ObservableTypeInput = "user_agent"
	WALLET       ObservableTypeInput = "wallet"
	MD5          ObservableTypeInput = "md5"
	IMPHASH      ObservableTypeInput = "imphash"
	ASN          ObservableTypeInput = "asn"
	NAMED_PIPE   ObservableTypeInput = "named_pipe"
	JA3          ObservableTypeInput = "ja3"
	GUESS        ObservableTypeInput = "guess"
)

// All allowed values of ObservableTypeInput enum
var AllowedObservableTypeInputEnumValues = []ObservableTypeInput{
	"cidr",
	"hostname",
	"mac_address",
	"certificate",
	"ipv6",
	"user_account",
	"ssdeep",
	"bic",
	"package",
	"sha256",
	"docker_image",
	"email",
	"jarm",
	"file",
	"mutex",
	"iban",
	"command_line",
	"url",
	"registry_key",
	"path",
	"ipv4",
	"tlsh",
	"generic",
	"sha1",
	"user_agent",
	"wallet",
	"md5",
	"imphash",
	"asn",
	"named_pipe",
	"ja3",
	"guess",
}

func (v *ObservableTypeInput) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObservableTypeInput(value)
	for _, existing := range AllowedObservableTypeInputEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObservableTypeInput", value)
}

// NewObservableTypeInputFromValue returns a pointer to a valid ObservableTypeInput
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObservableTypeInputFromValue(v string) (*ObservableTypeInput, error) {
	ev := ObservableTypeInput(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObservableTypeInput: valid values are %v", v, AllowedObservableTypeInputEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObservableTypeInput) IsValid() bool {
	for _, existing := range AllowedObservableTypeInputEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservableType-Input value
func (v ObservableTypeInput) Ptr() *ObservableTypeInput {
	return &v
}

type NullableObservableTypeInput struct {
	value *ObservableTypeInput
	isSet bool
}

func (v NullableObservableTypeInput) Get() *ObservableTypeInput {
	return v.value
}

func (v *NullableObservableTypeInput) Set(val *ObservableTypeInput) {
	v.value = val
	v.isSet = true
}

func (v NullableObservableTypeInput) IsSet() bool {
	return v.isSet
}

func (v *NullableObservableTypeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservableTypeInput(val *ObservableTypeInput) *NullableObservableTypeInput {
	return &NullableObservableTypeInput{value: val, isSet: true}
}

func (v NullableObservableTypeInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservableTypeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
