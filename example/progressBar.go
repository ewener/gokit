package main

import (
	"time"

	"github.com/ewener/gokit"
)

func main() {
	var bar gokit.Bar
	bar.NewOption(0, 543)
	for i := 0; i <= 543; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}
