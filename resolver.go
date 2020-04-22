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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Resolver Resolver
//
// Runtime DNS configuration
//
// swagger:model resolver
type Resolver struct {

	// accepted payload size
	// Minimum: 1232
	AcceptedPayloadSize int64 `json:"accepted_payload_size,omitempty"`

	// hold nx
	HoldNx int64 `json:"hold_nx,omitempty"`

	// hold obsolete
	HoldObsolete int64 `json:"hold_obsolete,omitempty"`

	// hold other
	HoldOther int64 `json:"hold_other,omitempty"`

	// hold refused
	HoldRefused int64 `json:"hold_refused,omitempty"`

	// hold timeout
	HoldTimeout int64 `json:"hold_timeout,omitempty"`

	// hold valid
	HoldValid int64 `json:"hold_valid,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name string `json:"name"`

	// parse resolv conf
	ParseResolvConf bool `json:"parse-resolv-conf,omitempty"`

	// resolve retries
	// Minimum: 1
	ResolveRetries int64 `json:"resolve_retries,omitempty"`

	// timeout resolve
	TimeoutResolve int64 `json:"timeout_resolve,omitempty"`

	// timeout retry
	TimeoutRetry int64 `json:"timeout_retry,omitempty"`
}

// Validate validates this resolver
func (m *Resolver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedPayloadSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolveRetries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resolver) validateAcceptedPayloadSize(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedPayloadSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("accepted_payload_size", "body", int64(m.AcceptedPayloadSize), 1232, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateResolveRetries(formats strfmt.Registry) error {

	if swag.IsZero(m.ResolveRetries) { // not required
		return nil
	}

	if err := validate.MinimumInt("resolve_retries", "body", int64(m.ResolveRetries), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resolver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resolver) UnmarshalBinary(b []byte) error {
	var res Resolver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
