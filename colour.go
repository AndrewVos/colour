package colour

import (
	"strconv"
)

var Enabled = true

func Black(text string) string   { return colour(text, 0) }
func Red(text string) string     { return colour(text, 1) }
func Green(text string) string   { return colour(text, 2) }
func Yellow(text string) string  { return colour(text, 3) }
func Blue(text string) string    { return colour(text, 4) }
func Magenta(text string) string { return colour(text, 5) }
func Cyan(text string) string    { return colour(text, 6) }
func White(text string) string   { return colour(text, 7) }
func Default(text string) string { return colour(text, 9) }

func colour(text string, colour int) string {
	if Enabled {
		return "\x1b[3" + strconv.Itoa(colour) + ";1m" + text + "\x1b[0m"
	}
	return text
}
