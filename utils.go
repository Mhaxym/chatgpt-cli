package main

const (
	// Colors
	GREEN   = "\033[32m"
	RED     = "\033[31m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	// Styles
	bold   = "\033[1m"
	italig = "\033[3m"
	// Reset
	reset = "\033[0m"
)

type StylerConfig struct {
	Color  string
	Bold   bool
	Italic bool
}

func ConsoleStyler(message string, config *StylerConfig) string {
	var style string
	if config.Bold {
		style += bold
	}
	if config.Italic {
		style += italig
	}
	style += config.Color
	style += message
	style += reset
	return style
}
