/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateWorkspaceRequestProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkspaceRequestProject{}

// CreateWorkspaceRequestProject struct for CreateWorkspaceRequestProject
type CreateWorkspaceRequestProject struct {
	Build   *ProjectBuild                        `json:"build,omitempty"`
	EnvVars *map[string]string                   `json:"envVars,omitempty"`
	Image   *string                              `json:"image,omitempty"`
	Name    string                               `json:"name"`
	Source  *CreateWorkspaceRequestProjectSource `json:"source,omitempty"`
	User    *string                              `json:"user,omitempty"`
}

type _CreateWorkspaceRequestProject CreateWorkspaceRequestProject

// NewCreateWorkspaceRequestProject instantiates a new CreateWorkspaceRequestProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceRequestProject(name string) *CreateWorkspaceRequestProject {
	this := CreateWorkspaceRequestProject{}
	this.Name = name
	return &this
}

// NewCreateWorkspaceRequestProjectWithDefaults instantiates a new CreateWorkspaceRequestProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceRequestProjectWithDefaults() *CreateWorkspaceRequestProject {
	this := CreateWorkspaceRequestProject{}
	return &this
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *CreateWorkspaceRequestProject) GetBuild() ProjectBuild {
	if o == nil || IsNil(o.Build) {
		var ret ProjectBuild
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequestProject) GetBuildOk() (*ProjectBuild, bool) {
	if o == nil || IsNil(o.Build) {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *CreateWorkspaceRequestProject) HasBuild() bool {
	if o != nil && !IsNil(o.Build) {
		return true
	}

	return false
}

// SetBuild gets a reference to the given ProjectBuild and assigns it to the Build field.
func (o *CreateWorkspaceRequestProject) SetBuild(v ProjectBuild) {
	o.Build = &v
}

// GetEnvVars returns the EnvVars field value if set, zero value otherwise.
func (o *CreateWorkspaceRequestProject) GetEnvVars() map[string]string {
	if o == nil || IsNil(o.EnvVars) {
		var ret map[string]string
		return ret
	}
	return *o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequestProject) GetEnvVarsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.EnvVars) {
		return nil, false
	}
	return o.EnvVars, true
}

// HasEnvVars returns a boolean if a field has been set.
func (o *CreateWorkspaceRequestProject) HasEnvVars() bool {
	if o != nil && !IsNil(o.EnvVars) {
		return true
	}

	return false
}

// SetEnvVars gets a reference to the given map[string]string and assigns it to the EnvVars field.
func (o *CreateWorkspaceRequestProject) SetEnvVars(v map[string]string) {
	o.EnvVars = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CreateWorkspaceRequestProject) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequestProject) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CreateWorkspaceRequestProject) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CreateWorkspaceRequestProject) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value
func (o *CreateWorkspaceRequestProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequestProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWorkspaceRequestProject) SetName(v string) {
	o.Name = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CreateWorkspaceRequestProject) GetSource() CreateWorkspaceRequestProjectSource {
	if o == nil || IsNil(o.Source) {
		var ret CreateWorkspaceRequestProjectSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequestProject) GetSourceOk() (*CreateWorkspaceRequestProjectSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CreateWorkspaceRequestProject) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given CreateWorkspaceRequestProjectSource and assigns it to the Source field.
func (o *CreateWorkspaceRequestProject) SetSource(v CreateWorkspaceRequestProjectSource) {
	o.Source = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateWorkspaceRequestProject) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequestProject) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateWorkspaceRequestProject) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CreateWorkspaceRequestProject) SetUser(v string) {
	o.User = &v
}

func (o CreateWorkspaceRequestProject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkspaceRequestProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Build) {
		toSerialize["build"] = o.Build
	}
	if !IsNil(o.EnvVars) {
		toSerialize["envVars"] = o.EnvVars
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

func (o *CreateWorkspaceRequestProject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateWorkspaceRequestProject := _CreateWorkspaceRequestProject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWorkspaceRequestProject)

	if err != nil {
		return err
	}

	*o = CreateWorkspaceRequestProject(varCreateWorkspaceRequestProject)

	return err
}

type NullableCreateWorkspaceRequestProject struct {
	value *CreateWorkspaceRequestProject
	isSet bool
}

func (v NullableCreateWorkspaceRequestProject) Get() *CreateWorkspaceRequestProject {
	return v.value
}

func (v *NullableCreateWorkspaceRequestProject) Set(val *CreateWorkspaceRequestProject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceRequestProject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceRequestProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceRequestProject(val *CreateWorkspaceRequestProject) *NullableCreateWorkspaceRequestProject {
	return &NullableCreateWorkspaceRequestProject{value: val, isSet: true}
}

func (v NullableCreateWorkspaceRequestProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceRequestProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
