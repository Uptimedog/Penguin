// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/clivern/penguin/core/model"

	"sync"
)

// Memory type
type Memory struct {
	sync.RWMutex
	items map[int]model.Metric
}

// NewCache creates a new instance of Memory
func NewMemoryCache() Memory {
	return Memory{items: make(map[int]model.Metric)}
}

// Get a key from a concurrent map
func (m *Memory) Get(key int) (model.Metric, bool) {
	m.Lock()
	defer m.Unlock()

	value, ok := m.items[key]

	return value, ok
}

// Set a key in a concurrent map
func (m *Memory) Set(key int, value model.Metric) {
	m.Lock()
	defer m.Unlock()

	m.items[key] = value
}

// Delete deletes a key
func (m *Memory) Delete(key int) {
	m.Lock()
	defer m.Unlock()

	delete(m.items, key)
}

// Length gets the  length
func (m *Memory) Length() int {
	m.Lock()
	defer m.Unlock()

	return len(m.items)
}

// Clear clears the map
func (m *Memory) Clear() {
	m.Lock()
	defer m.Unlock()

	m.items = make(map[int]model.Metric)
}

// Clear convert into a list
func (m *Memory) List() []model.Metric {
	m.Lock()

	defer m.Unlock()

	list := make([]model.Metric, 0)

	for _, v := range m.items {
		list = append(list, v)
	}

	return list
}
