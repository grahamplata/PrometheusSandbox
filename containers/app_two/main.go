package main

import (
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

// PushExampleData emits data to prometheus push gateway
func PushExampleData() {

	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})

	completionTime.SetToCurrentTime()

	if err := push.New(os.Getenv("DOCKER_NETWORK"), "postgres").
		Collector(completionTime).
		Grouping("db", "customers").
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}

}

func main() {
	for {
		PushExampleData()
		time.Sleep(5 * time.Second)
	}
}
