package main

import (
	"fmt"
)

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
	green  = "\033[32m"
	gray   = "\033[90m"
)

func main() {
	art := []string{
		"  ╔══════════════════════════════╗",
		"  ║  (⌐■_■)  sudo make me cool  ║",
		"  ╚══════════════════════════════╝",
	}
	fmt.Println()
	for _, l := range art {
		fmt.Println(cyan + l + reset)
	}
	fmt.Println()
	fmt.Printf("  %s%s%s  %s(%s)%s\n", bold+yellow, "Müslüm Barış Korkmazer", reset, gray, "babico", reset)
	fmt.Println()
	line := func(label, value, color string) {
		fmt.Printf("  %s%-10s%s%s%s%s\n", gray, label, reset, color, value, reset)
	}
	line("Company",  "NGX Storage",                       cyan)
	line("Location", "Ankara, Turkey",                    cyan)
	line("Web",      "https://babico.tr",                 green)
	line("Twitter",  "https://twitter.com/babicoq",       green)
	line("GitHub",   "https://github.com/babico",         green)
	line("npm",      "https://www.npmjs.com/package/babico", green)
	fmt.Println()
}
