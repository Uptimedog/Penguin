// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/clivern/penguin/core/backend"
	"github.com/clivern/penguin/core/cache"
	"github.com/clivern/penguin/core/model"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Daemon function
func Daemon(messages <-chan string) {
	i := 1
	bucket := cache.NewMemoryCache()
	console := backend.NewConsole()
	graphite := backend.NewGraphite()
	prometheus := backend.NewPrometheus()

	for message := range messages {
		metric := model.Metric{}
		metric.LoadFromJSON([]byte(message))

		log.WithFields(log.Fields{
			"metric": metric,
		}).Info("Metric received")

		if viper.GetBool("cache.enabled") && viper.GetString("cache.type") == "memory" {
			// Caching enabled
			if bucket.Length() < viper.GetInt("cache.drivers.memory.buffer_size")-1 {
				bucket.Set(i, metric)
				i += 1
				log.WithFields(log.Fields{
					"count": bucket.Length(),
				}).Info("Current metrics in cache")
				continue
			} else {
				bucket.Set(i, metric)
				i += 1
			}
		} else {
			// No caching
			bucket.Set(i, metric)
		}

		log.WithFields(log.Fields{
			"count": bucket.Length(),
		}).Info("Metrics aggregated for sending")

		// Send to console backend
		if viper.GetBool("output.console.enabled") {
			log.WithFields(log.Fields{
				"count": bucket.Length(),
			}).Info("Flush metrics to console backend")

			err := console.Send(bucket.List())

			if err != nil {
				log.WithFields(log.Fields{
					"error": err.Error(),
				}).Error("Failure during flushing metrics to console backend")
			}
		}

		// Send to graphite backend
		if viper.GetBool("output.graphite.enabled") {
			log.WithFields(log.Fields{
				"count": bucket.Length(),
			}).Info("Flush metrics to graphite backend")

			err := graphite.Send(bucket.List())

			if err != nil {
				log.WithFields(log.Fields{
					"error": err.Error(),
				}).Error("Failure during flushing metrics to graphite backend")
			}
		}

		// Send to prometheus backend
		if viper.GetBool("output.prometheus.enabled") {
			log.WithFields(log.Fields{
				"count": bucket.Length(),
			}).Info("Flush metrics to prometheus backend")

			err := prometheus.Send(bucket.List())

			if err != nil {
				log.WithFields(log.Fields{
					"error": err.Error(),
				}).Error("Failure during flushing metrics to prometheus backend")
			}
		}

		// Reset Bucket
		i = 1
		bucket.Clear()

		log.WithFields(log.Fields{
			"count": bucket.Length(),
		}).Info("Bucket cleared")
	}
}
