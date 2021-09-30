package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
)

func main() {
	fmt.Println("Go cactus metrics client")

	cfg := &statsd.ClientConfig{
		Address: "127.0.0.1:8125",
		Prefix:  "cactus-client",
	}

	client, err := statsd.NewClientWithConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	var metricName = "main.forloop.increment"
	var metricValue int64 = 0

	for i := 1; true; i++ {
		metricValue += int64(i)
		client.Inc(metricName, metricValue, 1.0)
		fmt.Printf("Incrementing metric %s.%s:%d => StatsD at %s\n", cfg.Prefix, metricName, metricValue, cfg.Address)
		time.Sleep(time.Second)
	}
}
