package gol

import (
	"fmt"
	"time"
)

const (
	BG_BLACK = "\u001b[40m"
	BG_RED = "\u001b[41m"
	BG_GREEN = "\u001b[42m"
	BG_YELLOW = "\u001b[43m"
	BG_BLUE = "\u001b[44m"
	BG_MAGENTA = "\u001b[45m"
	BG_CYAN = "\u001b[46m"
	BG_WHITE = "\u001b[47m"
	FG_BLACK = "\u001b[30m"
	FG_RED = "\u001b[31m"
	FG_GREEN = "\u001b[32m"
	FG_YELLOW = "\u001b[33m"
	FG_BLUE = "\u001b[34m"
	FG_MAGENTA = "\u001b[35m"
	FG_CYAN = "\u001b[36m"
	FG_WHITE = "\u001b[37m"
	RESET = "\u001b[0m"
)

func LogInfo(args ...interface{}) {
	fmt.Printf("%s INFO %s", BG_GREEN + FG_BLACK, RESET)
	fmt.Printf("%s %s %s %s ", BG_BLACK + FG_GREEN, time.Now().Format(time.RFC3339), BG_GREEN, RESET)
	fmt.Print(args...)
	fmt.Print("\n")
}
