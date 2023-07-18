package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var m runtime.MemStats
	var pollCount int64
	var randomValue float64

	reportInterval := 10 * time.Second
	pollInterval := 2 * time.Second

	client := NewClient()

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

		client.SendGaugeMetric("Alloc", float64(m.Alloc))
		client.SendGaugeMetric("BuckHashSys", float64(m.BuckHashSys))
		client.SendGaugeMetric("GCCPUFraction", m.GCCPUFraction)
		client.SendGaugeMetric("GCSys", float64(m.GCSys))
		client.SendGaugeMetric("HeapAlloc", float64(m.HeapAlloc))
		client.SendGaugeMetric("HeapIdle", float64(m.HeapIdle))
		client.SendGaugeMetric("HeapInuse", float64(m.HeapInuse))
		client.SendGaugeMetric("HeapObjects", float64(m.HeapObjects))
		client.SendGaugeMetric("HeapReleased", float64(m.HeapReleased))
		client.SendGaugeMetric("HeapSys", float64(m.HeapSys))
		client.SendGaugeMetric("LastGC", float64(m.LastGC))
		client.SendGaugeMetric("Lookups", float64(m.Lookups))
		client.SendGaugeMetric("MCacheInuse", float64(m.MCacheInuse))
		client.SendGaugeMetric("MCacheSys", float64(m.MCacheSys))
		client.SendGaugeMetric("MSpanInuse", float64(m.MSpanInuse))
		client.SendGaugeMetric("MSpanSys", float64(m.MSpanSys))
		client.SendGaugeMetric("Mallocs", float64(m.Mallocs))
		client.SendGaugeMetric("NextGC", float64(m.NextGC))
		client.SendGaugeMetric("NumForcedGC", float64(m.NumForcedGC))
		client.SendGaugeMetric("NumGC", float64(m.NumGC))
		client.SendGaugeMetric("OtherSys", float64(m.OtherSys))
		client.SendGaugeMetric("PauseTotalNs", float64(m.PauseTotalNs))
		client.SendGaugeMetric("StackInuse", float64(m.StackInuse))
		client.SendGaugeMetric("StackSys", float64(m.StackSys))
		client.SendGaugeMetric("Sys", float64(m.Sys))
		client.SendGaugeMetric("TotalAlloc", float64(m.TotalAlloc))
		client.SendGaugeMetric("TotalAlloc", float64(m.TotalAlloc))
		client.SendGaugeMetric("RandomValue", randomValue)

		client.SendCounterMetric("PollCount", pollCount)
		fmt.Println("alloc", alloc)
		fmt.Println("buckHashSys", buckHashSys)
		time.Sleep(reportInterval)
	}
}
