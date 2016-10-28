package main

import (
	"time"
)

type DataPoint struct {
	Timestamp  time.Time `json:"timestamp"`
	Value      float64   `json:"value"`
}