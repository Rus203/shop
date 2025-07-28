package utils

import (
	"math/rand"
	"time"

	"github.com/Rus203/shop/logger"
)

func GenerateRandomDuration(min, max int) time.Duration {
	if min > max || min < 1 {
		logger.Panic("Invalid range of time")
	}

	randomSec := rand.Intn(max - min + 1) + min
	return time.Duration(randomSec) * time.Second
}


