// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package backend

import (
	"fmt"

	"github.com/clivern/penguin/core/model"

	log "github.com/sirupsen/logrus"
)

// Graphite struct
type Graphite struct{}

// NewGraphite create a new instance of graphite backend
func NewGraphite() *Graphite {
	return &Graphite{}
}

// Send sends metrics to graphite backend
func (g *Graphite) Send(metrics []model.Metric) error {
	log.Info(fmt.Sprintf(
		"Send %d metrics to graphite backend",
		len(metrics),
	))

	return nil
}
