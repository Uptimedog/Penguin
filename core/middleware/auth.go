// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Auth middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		if path == viper.GetString("inputs.http.path") {
			apiKey := c.GetHeader("X-API-KEY")
			if viper.GetString("inputs.http.api_key") != "" && apiKey != viper.GetString("inputs.http.api_key") {
				log.WithFields(log.Fields{
					"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
					"http_method":    method,
					"http_path":      path,
					"api_key":        apiKey,
				}).Info(`Unauthorized access`)

				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
