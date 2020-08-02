// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Listener controller
func Listener(c *gin.Context, messages chan<- string) {
	messages <- "wip1"

	c.Status(http.StatusAccepted)
	return
}
