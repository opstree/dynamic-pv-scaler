package utils

import (
	"dynamic-pv-scaling/logger"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func CalculateUpdatedSize(value int, percentage int) int {
	logger.LogStdout()

	initialValue := value * percentage / 100
	updateValue := value + initialValue
	log.WithFields(log.Fields{
		"Scale Percentage": percentage,
	}).Info("Successfully calculated percentage, the new size is " + strconv.Itoa(updateValue))
	return updateValue
}
