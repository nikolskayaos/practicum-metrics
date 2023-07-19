package main

import (
	"fmt"
	"github.com/nikolskayaos/practicum-metrics/cmd/agent/client"
	"math/rand"
	"runtime"
	"time"
)

const BaseURL = "http://localhost:8080/update"

func main() {
	var m runtime.MemStats
	var pollCount int64
	var randomValue float64

	reportInterval := 10 * time.Second
	pollInterval := 2 * time.Second

	cl := client.NewClient(BaseURL)

	go func() {
		for {
			runtime.ReadMemStats(&m)
			pollCount++
			randomValue = rand.ExpFloat64()
			fmt.Println(pollCount, randomValue)
			time.Sleep(pollInterval)
		}
	}()

	for {
		alloc := m.Alloc
		buckHashSys := m.BuckHashSys

		cl.SendGaugeMetric("Alloc", float64(m.Alloc))
		cl.SendGaugeMetric("BuckHashSys", float64(m.BuckHashSys))
		cl.SendGaugeMetric("GCCPUFraction", m.GCCPUFraction)
		cl.SendGaugeMetric("GCSys", float64(m.GCSys))
		cl.SendGaugeMetric("HeapAlloc", float64(m.HeapAlloc))
		cl.SendGaugeMetric("HeapIdle", float64(m.HeapIdle))
		cl.SendGaugeMetric("HeapInuse", float64(m.HeapInuse))
		cl.SendGaugeMetric("HeapObjects", float64(m.HeapObjects))
		cl.SendGaugeMetric("HeapReleased", float64(m.HeapReleased))
		cl.SendGaugeMetric("HeapSys", float64(m.HeapSys))
		cl.SendGaugeMetric("LastGC", float64(m.LastGC))
		cl.SendGaugeMetric("Lookups", float64(m.Lookups))
		cl.SendGaugeMetric("MCacheInuse", float64(m.MCacheInuse))
		cl.SendGaugeMetric("MCacheSys", float64(m.MCacheSys))
		cl.SendGaugeMetric("MSpanInuse", float64(m.MSpanInuse))
		cl.SendGaugeMetric("MSpanSys", float64(m.MSpanSys))
		cl.SendGaugeMetric("Mallocs", float64(m.Mallocs))
		cl.SendGaugeMetric("NextGC", float64(m.NextGC))
		cl.SendGaugeMetric("NumForcedGC", float64(m.NumForcedGC))
		cl.SendGaugeMetric("NumGC", float64(m.NumGC))
		cl.SendGaugeMetric("OtherSys", float64(m.OtherSys))
		cl.SendGaugeMetric("PauseTotalNs", float64(m.PauseTotalNs))
		cl.SendGaugeMetric("StackInuse", float64(m.StackInuse))
		cl.SendGaugeMetric("StackSys", float64(m.StackSys))
		cl.SendGaugeMetric("Sys", float64(m.Sys))
		cl.SendGaugeMetric("TotalAlloc", float64(m.TotalAlloc))
		cl.SendGaugeMetric("TotalAlloc", float64(m.TotalAlloc))
		cl.SendGaugeMetric("RandomValue", randomValue)

		cl.SendCounterMetric("PollCount", pollCount)
		fmt.Println("alloc", alloc)
		fmt.Println("buckHashSys", buckHashSys)
		time.Sleep(reportInterval)
	}
}
