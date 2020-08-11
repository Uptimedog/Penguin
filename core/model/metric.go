// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
)

// Metric struct
type Metric struct {
	Type   string            `json:"type"`
	Name   string            `json:"name"`
	Help   string            `json:"help"`
	Method string            `json:"method"`
	Value  string            `json:"value"`
	Labels map[string]string `json:"labels"`
}

// LoadFromJSON update object from json
func (m *Metric) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (m *Metric) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
