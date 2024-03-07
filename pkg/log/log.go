package log

import (
	"fmt"
	"time"
)

func LogInfo(s string) {
	fmt.Printf("%s [+]Info: %s\n", time.Now().Format(time.RFC3339), s)
}

func LogWarning(s string) {
	fmt.Printf("%s [-]Warning: %s\n", time.Now().Format(time.RFC3339), s)
}

func LogError(s string) {
	fmt.Printf("%s [-]Error: %s\n", time.Now().Format(time.RFC3339), s)
}
