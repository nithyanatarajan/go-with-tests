package print

import "time"

type Sleeper interface {
	Sleep()
}
type DefaultSleeper struct {
}

func (*DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
