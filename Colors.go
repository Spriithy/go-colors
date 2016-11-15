package colors

import "fmt"

type style uint8

const (
	// None means neither bold nor underline
	None = style(iota)

	// Bold text
	Bold

	// Underline text - NOT IMPLEMENTED YET
	Underline
)

var prefix = "\033["

var (
	red        = prefix + "0;31m"
	green      = prefix + "0;32m"
	yellow     = prefix + "0;33m"
	blue       = prefix + "0;34m"
	purple     = prefix + "0;35m"
	cyan       = prefix + "0;36m"
	boldRed    = prefix + "1;31m"
	boldGreen  = prefix + "1;32m"
	boldYellow = prefix + "1;33m"
	boldBlue   = prefix + "1;34m"
	boldPurple = prefix + "1;35m"
	boldCyan   = prefix + "1;36m"
	black      = prefix + "0;30m"
	white      = prefix + "1;37m"
	lightGray  = prefix + "0;37m"
	darkGray   = prefix + "1;30m"
	none       = prefix + "0m"
)

func concat(a ...interface{}) string {
	str := ""
	for _, x := range a {
		str += fmt.Sprintf("%s", x)
	}
	return str
}

// Black returns an ASCII Black string made of all interfaces{} passed
func Black(w style, a ...interface{}) string {
	return black + concat(a...) + none
}

// Red returns an ASCII re string made of all interfaces{} passed
func Red(w style, a ...interface{}) string {
	c := red
	if w == Bold {
		c = boldRed
	}
	return c + concat(a...) + none
}

// Green returns an ASCII re string made of all interfaces{} passed
func Green(w style, a ...interface{}) string {
	c := green
	if w == Bold {
		c = boldGreen
	}
	return c + concat(a...) + none
}

// Yellow returns an ASCII re string made of all interfaces{} passed
func Yellow(w style, a ...interface{}) string {
	c := yellow
	if w == Bold {
		c = boldYellow
	}
	return c + concat(a...) + none
}

// Blue returns an ASCII re string made of all interfaces{} passed
func Blue(w style, a ...interface{}) string {
	c := blue
	if w == Bold {
		c = boldBlue
	}
	return c + concat(a...) + none
}

// Purple returns an ASCII re string made of all interfaces{} passed
func Purple(w style, a ...interface{}) string {
	c := purple
	if w == Bold {
		c = boldPurple
	}
	return c + concat(a...) + none
}

// Cyan returns an ASCII re string made of all interfaces{} passed
func Cyan(w style, a ...interface{}) string {
	c := cyan
	if w == Bold {
		c = boldCyan
	}
	return c + concat(a...) + none
}

// TODO : other colors
