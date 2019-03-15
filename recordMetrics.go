package main

import (
	"fmt"
	"time"

	"github.com/prometheus/common/log"
)

func recordMetrics() {
	conf, err := getConfig()
	if err != nil {
		fmt.Println(err)
		log.Errorln("Failed during config reading")
		return
	}
	go func() {
		for {
			time.Sleep(conf.ProbSec * time.Second)
			resAuth := authResponse{}

			// Write profile file
			profile()

			// Auth and check Keystobe API
			for {
				resAuth, err = osAuth(conf)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Authentication failed")
					osAPIcheck.WithLabelValues(keystone).Set(1)
					time.Sleep(10 * time.Second) // wait a 10 sec
					// os.Exit(1)
				}
				osAPIcheck.WithLabelValues(keystone).Set(1)
				break
			}

			// Get Nova services state
			go func() {
				novaServs, err := getNovaServices(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get nova services")
					osAPIcheck.WithLabelValues(nova).Set(0)
					return
				}
				novaServs.Metrics()
				osAPIcheck.WithLabelValues(nova).Set(1)
			}()

			// Get Nova hypervisors metrics
			go func() {
				hypers, err := getNovaHypervisors(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get nova hypervisors")
					osAPIcheck.WithLabelValues(nova).Set(0)
					return
				}
				hypers.Metrics()
				osAPIcheck.WithLabelValues(nova).Set(1)
			}()

			// Get Nova projects limits and usage
			go func() {
				pl, err := getNovaLimits(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get nova limits")
					osAPIcheck.WithLabelValues(nova).Set(0)
					return
				}
				pl.Metrics()
				osAPIcheck.WithLabelValues(nova).Set(1)
			}()

			// Get Nova projects Root Disk usage
			go func() {
				pl, err := getAllProjectsUsage(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get nova root disk usage")
					osAPIcheck.WithLabelValues(nova).Set(0)
					return
				}
				pl.Metrics()
				osAPIcheck.WithLabelValues(nova).Set(1)
			}()

			// Get Nova instances statuse for projects and usage
			go func() {
				apai, err := getServersInfo(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get nova limits")
					osAPIcheck.WithLabelValues(nova).Set(0)
					return
				}
				apai.Metrics()
				osAPIcheck.WithLabelValues(nova).Set(1)
			}()

			// Get Neutron agent state
			go func() {
				netAgents, err := getNeutronAgents(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get neutron services")
					osAPIcheck.WithLabelValues(neutron).Set(0)
					return
				}
				netAgents.Metrics()
				osAPIcheck.WithLabelValues(neutron).Set(1)
			}()

			// Check Glance API
			go func() {
				_, err := getGlanceImages(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get glance images")
					osAPIcheck.WithLabelValues(glance).Set(0)
					return
				}
				osAPIcheck.WithLabelValues(glance).Set(1)
			}()

			// Check Cinder services
			go func() {
				cinderServ, err := getCinderServices(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get cinder services")
					osAPIcheck.WithLabelValues(cinder).Set(0)
					return
				}
				cinderServ.Metrics()
				osAPIcheck.WithLabelValues(cinder).Set(1)
			}()

			// Get Cinder projects limits and usage
			go func() {
				pl, err := getCinderLimits(resAuth)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Can't get cinder limits")
					osAPIcheck.WithLabelValues(cinder).Set(0)
					return
				}
				pl.Metrics()
				osAPIcheck.WithLabelValues(cinder).Set(1)
			}()

			// Get Horizon availability metric
			go func() {
				err := loginHorizon(conf)
				if err != nil {
					fmt.Println(err)
					log.Errorln("Horizon not available")
					horizonAvailable.Set(0)
					return
				}
				horizonAvailable.Set(1)
			}()

		}
	}()
}
