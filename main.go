package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	conf, err := getConfig()
	if err != nil {
		fmt.Println(err)
		log.Errorln("Failed during config reading")
		// log.Errorln("Failed during config reading")
		return
	}
	regSrvMetrics()
	recordMetrics()
	log.Infoln("Starting openstack-exporter")
	log.Infoln("Path of metrics /metrics")
	log.Infoln("Port to listen on for telemetry", conf.Port)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+conf.Port, nil)
}
