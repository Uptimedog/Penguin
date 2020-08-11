// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package backend

import (
	"fmt"

	"github.com/clivern/penguin/core/model"

	log "github.com/sirupsen/logrus"
)

// Prometheus struct
type Prometheus struct{}

func (p *Prometheus) Send(metrics []model.Metric) error {
	log.Info(fmt.Sprintf(
		"Send %d metrics to prometheus backend",
		len(metrics),
	))

	return nil
}
