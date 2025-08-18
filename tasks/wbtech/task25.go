package main

import "time"

// Задание 25
// Реализовать собственную функцию sleep.

func MySleep(d time.Duration) {
	now := time.Now()
	for time.Since(now) < d {
	}
}

func main() {
	MySleep(5 * time.Second)
}
