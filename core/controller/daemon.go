// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	log "github.com/sirupsen/logrus"
)

// Daemon function
func Daemon(messages <-chan string) {
	for message := range messages {
		log.Info(message)
	}
}
