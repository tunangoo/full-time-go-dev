package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tunangoo/full-time-go-dev/toll-calculator/types"
)

type LoggingMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LoggingMiddleware{
		next: next,
	}
}

func (m *LoggingMiddleware) CalculateDistance(data types.OBUdata) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"dist": dist,
			"err":  err,
		}).Info("calculate distance")
	}(time.Now())

	dist, err = m.next.CalculateDistance(data)
	return
}
