package structs

import (
	"errors"
	"math"
	"math/rand"
)

const (
	RAND_SUM_DEFAULT_MIN = -100000
	RAND_SUM_DEFAULT_MAX = 100000
)

func generateSum(unsigned bool, limitTop *int64, limitBottom *int64) (float64, error) {
	var min, max, rndLimit int64
	var rnd float64
	if limitBottom == nil {
		if unsigned {
			min = 0
		} else {
			min = RAND_SUM_DEFAULT_MIN
		}
	} else {
		min = *limitBottom
	}
	if max <= min {
		return 0, errors.New("max could not be smaller than min")
	}

	if limitTop == nil {
		max = RAND_SUM_DEFAULT_MAX
	} else {
		max = *limitTop
	}

	if min > 0 {
		rndLimit := max - min
		rnd = float64(rand.Int63n(rndLimit) + min)
	} else if max < 0 && min < 0 {
		rndLimit = max - min
		rnd = float64(-1*rand.Int63n(rndLimit) + max)
	} else {
		var pow float64
		rndSwitch := rand.Int63n(100)
		if rndSwitch > 50 {
			pow = 2
			rndLimit = min
		} else {
			pow = 1
			rndLimit = max
		}
		rnd = math.Pow(-1, pow) * float64(rand.Int63n(rndLimit))
	}

	return rnd, nil
}
