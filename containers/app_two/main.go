package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func PushExampleData() {
	exampleCron := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cron_completion_timestamp",
		Help: "A Phoney Baloney generated timestamp",
	})
	exampleCron.SetToCurrentTime()
	if err := push.New(os.Getenv("DOCKER_NETWORK"), "cron").
		Collector(exampleCron).
		Push(); err != nil {
		fmt.Println("Could not push to Pushgateway:", err)
	}
}

func main() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
	PushExampleData()
}
