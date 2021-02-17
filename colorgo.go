package colorgo

import (
	"colorgo/colors"
	"fmt"
)

type colorgo struct {
	selectedColor string
}

func NewColorGo() *colorgo {
	result := new(colorgo)
	result.SetColor(colors.DebugColor)
	return result
}

func (c *colorgo) SetColor(color string) {
	c.selectedColor = color
}

func (c *colorgo) Println(str string) {
	fmt.Println(fmt.Printf(c.selectedColor, str))
}
