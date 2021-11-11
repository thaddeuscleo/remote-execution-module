package utils

import "fmt"

// var (
// 	colorReset  = "\033[0m"
// 	colorRed    = "\033[31m"
// 	colorGreen  = "\033[32m"
// 	colorYellow = "\033[33m"
// 	colorCyan   = "\033[36m"
// )

func LogInfo(content string) {
	fmt.Printf("[info] %s", content)
}

func LogWarning(content string) {
	fmt.Printf("[waring] %s", content)
}

func LogError(content string) {
	fmt.Printf("[error] %s", content)
}

func LogSuccess(content string) {
	fmt.Printf("[success] %s", content)
}
