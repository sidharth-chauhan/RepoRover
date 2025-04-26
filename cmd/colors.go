	package cmd

	import (
		"github.com/fatih/color"
	)

	// Regular colors
	var (
		// Red returns a red colored string
		Red = color.New(color.FgRed).SprintFunc()
		// Green returns a green colored string
		Green = color.New(color.FgGreen).SprintFunc()
		// Yellow returns a yellow colored string
		Yellow = color.New(color.FgYellow).SprintFunc()
		// Blue returns a blue colored string
		Blue = color.New(color.FgBlue).SprintFunc()
		// Magenta returns a magenta colored string
		Magenta = color.New(color.FgMagenta).SprintFunc()
		// Cyan returns a cyan colored string
		Cyan = color.New(color.FgCyan).SprintFunc()
		// White returns a white colored string
		White = color.New(color.FgWhite).SprintFunc()
	)

	// Faint colors - lower intensity versions of regular colors
	var (
		// FaintRed returns a faint red colored string
		FaintRed = color.New(color.FgRed).Add(color.Faint).SprintFunc()
		// FaintGreen returns a faint green colored string
		FaintGreen = color.New(color.FgGreen).Add(color.Faint).SprintFunc()
		// FaintYellow returns a faint yellow colored string
		FaintYellow = color.New(color.FgYellow).Add(color.Faint).SprintFunc()
		// FaintBlue returns a faint blue colored string
		FaintBlue = color.New(color.FgBlue).Add(color.Faint).SprintFunc()
		// FaintMagenta returns a faint magenta colored string
		FaintMagenta = color.New(color.FgMagenta).Add(color.Faint).SprintFunc()
		// FaintCyan returns a faint cyan colored string
		FaintCyan = color.New(color.FgCyan).Add(color.Faint).SprintFunc()
		// FaintWhite returns a faint white colored string
		FaintWhite = color.New(color.FgWhite).Add(color.Faint).SprintFunc()
	)

	// Bold colors
	var (
		// BoldRed returns a bold red colored string
		BoldRed = color.New(color.FgRed, color.Bold).SprintFunc()
		// BoldGreen returns a bold green colored string
		BoldGreen = color.New(color.FgGreen, color.Bold).SprintFunc()
		// BoldYellow returns a bold yellow colored string
		BoldYellow = color.New(color.FgYellow, color.Bold).SprintFunc()
		// BoldBlue returns a bold blue colored string
		BoldBlue = color.New(color.FgBlue, color.Bold).SprintFunc()
		// BoldMagenta returns a bold magenta colored string
		BoldMagenta = color.New(color.FgMagenta, color.Bold).SprintFunc()
		// BoldCyan returns a bold cyan colored string
		BoldCyan = color.New(color.FgCyan, color.Bold).SprintFunc()
		// BoldWhite returns a bold white colored string
		BoldWhite = color.New(color.FgWhite, color.Bold).SprintFunc()
	)

