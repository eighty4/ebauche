package ebauche

import "time"

// IntervalIndefinitely invokes fn indefinitely on interval specified with time.Duration
func IntervalIndefinitely(interval time.Duration, fn func()) {
	Interval(interval, func() bool {
		fn()
		return true
	})
}

// Interval invokes fn until fn returns false on interval specified with time.Duration
func Interval(interval time.Duration, fn func() bool) {
	cont := true
	for cont {
		start := time.Now()
		if cont = fn(); cont {
			elapsed := time.Now().Sub(start)
			delay := interval - elapsed
			time.Sleep(delay)
		}
	}
}
