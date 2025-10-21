package common

import (
	"math"
	"math/rand/v2"
	"strconv"

	// "time"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
)

func GenerateOtp() string {
	cfg := config.GetConfig()
	// rand.Seed(time.Now().UnixNano())
	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))   // 10 ^ d - 1 100000
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1) // 999999 = 100000 - 1 (10 ^ d) -1

	var num = rand.IntN(max-min) + min
	return strconv.Itoa(num)

}
