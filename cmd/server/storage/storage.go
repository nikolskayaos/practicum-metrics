package storage

import "fmt"

type MemStorage struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Gauges:   make(map[string]float64),
		Counters: make(map[string]int64),
	}
}

func (m *MemStorage) SaveGauges(name string, gauge float64) {
	m.Gauges[name] = gauge
	fmt.Println(m.Gauges)
}

func (m *MemStorage) SaveCounters(name string, count int64) {
	m.Counters[name] += count
	fmt.Println(m.Counters)
}
