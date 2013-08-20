package colour

import (
  "strconv"
)

func Black(text string)   { colour(text, 0) }
func Red(text string)     { colour(text, 1) }
func Green(text string)   { colour(text, 2) }
func Yellow(text string)  { colour(text, 3) }
func Blue(text string)    { colour(text, 4) }
func Magenta(text string) { colour(text, 5) }
func Cyan(text string)    { colour(text, 6) }
func White(text string)   { colour(text, 7) }
func Default(text string) { colour(text, 9) }

func colour(text string, colour int) string {
  return "\x1b[3"+strconv.Itoa(colour)+";1m" + text + "\x1b[0m"
}
