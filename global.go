// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Global Global
//
// HAProxy global configuration
//
// swagger:model global
type Global struct {

	// CPU maps
	CPUMaps []*CPUMap `json:"cpu_maps"`

	// runtime a p is
	RuntimeAPIs []*RuntimeAPI `json:"runtime_apis"`

	// chroot
	// Pattern: ^[^\s]+$
	Chroot string `json:"chroot,omitempty"`

	// daemon
	// Enum: [enabled disabled]
	Daemon string `json:"daemon,omitempty"`

	// external check
	ExternalCheck bool `json:"external_check,omitempty"`

	// group
	// Pattern: ^[^\s]+$
	Group string `json:"group,omitempty"`

	// lua loads
	LuaLoads []*LuaLoad `json:"lua_loads"`

	// master worker
	MasterWorker bool `json:"master-worker,omitempty"`

	// maxconn
	Maxconn int64 `json:"maxconn,omitempty"`

	// nbproc
	Nbproc int64 `json:"nbproc,omitempty"`

	// nbthread
	Nbthread int64 `json:"nbthread,omitempty"`

	// pidfile
	Pidfile string `json:"pidfile,omitempty"`

	// ssl default bind ciphers
	SslDefaultBindCiphers string `json:"ssl_default_bind_ciphers,omitempty"`

	// ssl default bind options
	SslDefaultBindOptions string `json:"ssl_default_bind_options,omitempty"`

	// ssl default server ciphers
	SslDefaultServerCiphers string `json:"ssl_default_server_ciphers,omitempty"`

	// ssl default server options
	SslDefaultServerOptions string `json:"ssl_default_server_options,omitempty"`

	// stats timeout
	StatsTimeout *int64 `json:"stats_timeout,omitempty"`

	// tune ssl default dh param
	TuneSslDefaultDhParam int64 `json:"tune_ssl_default_dh_param,omitempty"`

	// user
	// Pattern: ^[^\s]+$
	User string `json:"user,omitempty"`
}

// Validate validates this global
func (m *Global) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUMaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimeAPIs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChroot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaemon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaLoads(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Global) validateCPUMaps(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUMaps) { // not required
		return nil
	}

	for i := 0; i < len(m.CPUMaps); i++ {
		if swag.IsZero(m.CPUMaps[i]) { // not required
			continue
		}

		if m.CPUMaps[i] != nil {
			if err := m.CPUMaps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpu_maps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateRuntimeAPIs(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeAPIs) { // not required
		return nil
	}

	for i := 0; i < len(m.RuntimeAPIs); i++ {
		if swag.IsZero(m.RuntimeAPIs[i]) { // not required
			continue
		}

		if m.RuntimeAPIs[i] != nil {
			if err := m.RuntimeAPIs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runtime_apis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateChroot(formats strfmt.Registry) error {

	if swag.IsZero(m.Chroot) { // not required
		return nil
	}

	if err := validate.Pattern("chroot", "body", string(m.Chroot), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var globalTypeDaemonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTypeDaemonPropEnum = append(globalTypeDaemonPropEnum, v)
	}
}

const (

	// GlobalDaemonEnabled captures enum value "enabled"
	GlobalDaemonEnabled string = "enabled"

	// GlobalDaemonDisabled captures enum value "disabled"
	GlobalDaemonDisabled string = "disabled"
)

// prop value enum
func (m *Global) validateDaemonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalTypeDaemonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Global) validateDaemon(formats strfmt.Registry) error {

	if swag.IsZero(m.Daemon) { // not required
		return nil
	}

	// value enum
	if err := m.validateDaemonEnum("daemon", "body", m.Daemon); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if err := validate.Pattern("group", "body", string(m.Group), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateLuaLoads(formats strfmt.Registry) error {

	if swag.IsZero(m.LuaLoads) { // not required
		return nil
	}

	for i := 0; i < len(m.LuaLoads); i++ {
		if swag.IsZero(m.LuaLoads[i]) { // not required
			continue
		}

		if m.LuaLoads[i] != nil {
			if err := m.LuaLoads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lua_loads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := validate.Pattern("user", "body", string(m.User), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Global) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Global) UnmarshalBinary(b []byte) error {
	var res Global
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CPUMap CPU map
//
// swagger:model CPUMap
type CPUMap struct {

	// cpu set
	// Required: true
	CPUSet *string `json:"cpu_set"`

	// process
	// Required: true
	Process *string `json:"process"`
}

// Validate validates this CPU map
func (m *CPUMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUMap) validateCPUSet(formats strfmt.Registry) error {

	if err := validate.Required("cpu_set", "body", m.CPUSet); err != nil {
		return err
	}

	return nil
}

func (m *CPUMap) validateProcess(formats strfmt.Registry) error {

	if err := validate.Required("process", "body", m.Process); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CPUMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUMap) UnmarshalBinary(b []byte) error {
	var res CPUMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LuaLoad lua load
//
// swagger:model LuaLoad
type LuaLoad struct {

	// file
	// Required: true
	// Pattern: ^[^\s]+$
	File *string `json:"file"`
}

// Validate validates this lua load
func (m *LuaLoad) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LuaLoad) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", m.File); err != nil {
		return err
	}

	if err := validate.Pattern("file", "body", string(*m.File), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LuaLoad) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LuaLoad) UnmarshalBinary(b []byte) error {
	var res LuaLoad
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RuntimeAPI runtime API
//
// swagger:model RuntimeAPI
type RuntimeAPI struct {

	// address
	// Required: true
	// Pattern: ^[^\s]+$
	Address *string `json:"address"`

	// expose fd listeners
	ExposeFdListeners bool `json:"exposeFdListeners,omitempty"`

	// level
	// Enum: [user operator admin]
	Level string `json:"level,omitempty"`

	// mode
	// Pattern: ^[^\s]+$
	Mode string `json:"mode,omitempty"`

	// process
	// Pattern: ^[^\s]+$
	Process string `json:"process,omitempty"`
}

// Validate validates this runtime API
func (m *RuntimeAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuntimeAPI) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if err := validate.Pattern("address", "body", string(*m.Address), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var runtimeApiTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","operator","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runtimeApiTypeLevelPropEnum = append(runtimeApiTypeLevelPropEnum, v)
	}
}

const (

	// RuntimeAPILevelUser captures enum value "user"
	RuntimeAPILevelUser string = "user"

	// RuntimeAPILevelOperator captures enum value "operator"
	RuntimeAPILevelOperator string = "operator"

	// RuntimeAPILevelAdmin captures enum value "admin"
	RuntimeAPILevelAdmin string = "admin"
)

// prop value enum
func (m *RuntimeAPI) validateLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, runtimeApiTypeLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RuntimeAPI) validateLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeAPI) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := validate.Pattern("mode", "body", string(m.Mode), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeAPI) validateProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.Process) { // not required
		return nil
	}

	if err := validate.Pattern("process", "body", string(m.Process), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeAPI) UnmarshalBinary(b []byte) error {
	var res RuntimeAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
