package main

import (
	"./print"
	"fmt"
	"math/rand"
	"time"
)

const (
	INITPOINTS = 100
	MINPOINTS  = 10
)

var MAXPOINTS int = INITPOINTS

func run() {
	print.Testprint()
	defaultDelayPrintln("let's play a craps game")
	defaultDelayPrintf("you have %d points at the begin, when you have less than %d points, you will loss\n", INITPOINTS, MINPOINTS)
	var sum = INITPOINTS
	for play(&sum) {
	}
	defaultDelayPrintf("game over at %d points, max points: %d\n", sum, MAXPOINTS)
}

func play(sum *int) bool {
	defaultDelayPrintf("now you have %d points, input number of points for your rasing\n", *sum)
	var point int
inputpoint:
	defaultDelayPrintf("input:")
	fmt.Scanf("%d\n", &point)
	if point > *sum {
		defaultDelayPrintln("you don't have enough points to raise, input again.")
		goto inputpoint
	}
	defaultDelayPrintln("choose odd or even[o/e]")
	var oddeven rune
inputoddeven:
	defaultDelayPrintf("choice:")
	fmt.Scanf("%c\n", &oddeven)
	if oddeven != 'e' && oddeven != 'o' {
		defaultDelayPrintln("wrong option, need e(for even) or o(for odd)")
		goto inputoddeven
	}
	delayPrintln(500*time.Millisecond, "waiting for your result...")
	rand.Seed(time.Now().UTC().UnixNano())
	result := rand.Intn(6) + 1
	defaultDelayPrintf("reulst is %d\n", result)
	iseven := result%2 == 0
	if (iseven && oddeven == 'e') || (!iseven && oddeven == 'o') {
		*sum = *sum + point
		defaultDelayPrintln("win at this turn!")
	} else {
		defaultDelayPrintln("you loose...")
		*sum = *sum - point
	}
	if *sum > MAXPOINTS {
		MAXPOINTS = *sum
	}
	return *sum >= MINPOINTS
}

const defaultLastDelayPeriod = 500 * time.Millisecond
const defaultDelayPeriod = 50 * time.Millisecond

func defaultDelayPrintf(format string, args ...interface{}) {
	delayPrintf(defaultDelayPeriod, format, args...)
}

func defaultDelayPrintln(content string) {
	delayPrintf(defaultDelayPeriod, content+"\n")
}

func delayPrintln(period time.Duration, content string) {
	delayPrintf(period, content+"\n")
}

func delayPrintf(period time.Duration, format string, args ...interface{}) {
	for _, c := range fmt.Sprintf(format, args...) {
		fmt.Printf("%c", c)
		time.Sleep(period)
	}
	time.Sleep(defaultDelayPeriod)
}

func main() {
	run()
}
