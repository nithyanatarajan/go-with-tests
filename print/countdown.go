package print

import (
	"fmt"
	"io"
)

const countDownStarter = 3
const finalWord = "Go!"

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countDownStarter; i > 0; i-- {
		// revive:disable-next-line
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	// revive:disable-next-line
	fmt.Fprint(writer, finalWord)
}
