package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	fmt.Println("Go datadog metrics client")

	var ip = "127.0.0.1"
	var port = "8125"
	var addr = ip + ":" + port

	client, err := statsd.New(addr)
	if err != nil {
		log.Fatal(err)
	}

	var metricName = "datadog-client.main.forloop.increment"
	var metricValue float64 = 1

	for i := 1; true; i++ {
		client.Incr(metricName, []string{}, metricValue)
		time.Sleep(time.Second)
		fmt.Printf("Incrementing metric %s:%f => StatsD at %s\n", metricName, metricValue, addr)
	}
}
