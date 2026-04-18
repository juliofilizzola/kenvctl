package utils

import (
	"fmt"
	"time"
)

// TimeSleep pauses the execution for 5 seconds, displaying a countdown message during the delay.
func TimeSleep() {
	for i := 5; i > 0; i-- {
		fmt.Printf("O programa irá fechar em %d segundos...\n", i)
		time.Sleep(1 * time.Second)
	}
}
