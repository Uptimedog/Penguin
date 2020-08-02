// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"time"
)

// Watcher function
func Watcher(messages chan<- string) {
	for {
		messages <- "wip2"
		time.Sleep(1 * time.Second)
	}
}
