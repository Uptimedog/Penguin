// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/uptimedog/penguin/core/backend"
	"github.com/uptimedog/penguin/core/model"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// Listen controller
func Listen(c echo.Context, prom *backend.Prometheus) error {

	apiKey := viper.GetString("app.api_key")

	if apiKey != "" && apiKey != c.Request().Header.Get("x-api-key") {
		return c.NoContent(http.StatusForbidden)
	}

	var bodyBytes []byte

	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	metric := &model.Metric{}
	metric.LoadFromJSON(bodyBytes)
	prom.Send(*metric)

	return c.NoContent(http.StatusCreated)
}
