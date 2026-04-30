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

// Info holds the personal card data.
var Info = struct {
	Name     string
	Handle   string
	Company  string
	Location string
	Website  string
	Twitter  string
	GitHub   string
	NPM      string
}{
	Name:     "Müslüm Barış Korkmazer",
	Handle:   "babico",
	Company:  "NGX Storage",
	Location: "Ankara, Turkey",
	Website:  "https://babico.tr",
	Twitter:  "https://twitter.com/babicoq",
	GitHub:   "https://github.com/babico",
	NPM:      "https://www.npmjs.com/package/babico",
}

func line(label, value, color string) string {
	return fmt.Sprintf("  %s%-10s%s%s%s%s", gray, label, reset, color, value, reset)
}

// Print prints the personal card to stdout.
func Print() {
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
	fmt.Printf("  %s%s%s  %s(%s)%s\n", bold+yellow, Info.Name, reset, gray, Info.Handle, reset)
	fmt.Println()
	fmt.Println(line("Company", Info.Company, cyan))
	fmt.Println(line("Location", Info.Location, cyan))
	fmt.Println(line("Web", Info.Website, green))
	fmt.Println(line("Twitter", Info.Twitter, green))
	fmt.Println(line("GitHub", Info.GitHub, green))
	fmt.Println(line("npm", Info.NPM, green))
	fmt.Println()
}

func main() {
	Print()
}
