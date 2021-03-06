// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"sync"

	"github.com/clivern/penguin/core/model"

	"github.com/nxadm/tail"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Watcher function
func Watcher(messages chan<- string) {
	paths := viper.GetStringSlice("inputs.log.paths")
	var wg sync.WaitGroup

	for _, path := range paths {

		log.WithFields(log.Fields{
			"log_file": path,
		}).Info("Watch log file")
		wg.Add(1)

		go func(swg *sync.WaitGroup, file string, channel chan<- string) {
			t, err := tail.TailFile(
				file,
				tail.Config{Follow: true, ReOpen: true},
			)

			if err != nil {
				panic(err)
			}

			for line := range t.Lines {
				metric := &model.Metric{}
				metric.LoadFromJSON([]byte(fmt.Sprintf("%s", line.Text)))
				message, err := metric.ConvertToJSON()

				if err != nil {
					log.WithFields(log.Fields{
						"error": err.Error(),
					}).Warn("Invalid metric received")
					continue
				}

				channel <- message
			}

			defer swg.Done()
		}(&wg, path, messages)
	}

	wg.Wait()
}
