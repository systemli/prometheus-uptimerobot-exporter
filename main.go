package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("web.listen-address", ":13121", "Address on which to expose metrics and web interface.")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if os.Getenv("UPTIMEROBOT_API_KEY") == "" {
		log.Fatal("UPTIMEROBOT_API_KEY is not set")
	}

	prometheus.MustRegister(NewCollector(os.Getenv("UPTIMEROBOT_API_KEY")))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
