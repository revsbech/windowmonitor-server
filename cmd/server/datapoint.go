package main

import (
	"time"
)

type DataPoint struct {
	Timestamp time.Time
	Value    float64
}