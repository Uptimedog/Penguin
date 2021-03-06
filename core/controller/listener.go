// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/clivern/penguin/core/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Listener controller
func Listener(c *gin.Context, messages chan<- string) {
	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	metric := &model.Metric{}
	metric.LoadFromJSON(bodyBytes)

	body, err := metric.ConvertToJSON()

	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Warn("Invalid metric received")
	} else {
		messages <- body
	}

	c.Status(http.StatusAccepted)
	return
}
