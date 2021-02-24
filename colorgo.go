package colorgo

import (
	"fmt"

	"github.com/Lornzo/colorgo/colors"
)

type colorgo struct {
	selectedColor string
}

func NewColorGo() *colorgo {
	result := new(colorgo)
	result.SetColor(colors.DefaultColor)
	return result
}

func (c *colorgo) SetColor(color string) {
	c.selectedColor = color
}

func (c *colorgo) Println(str string) {
	fmt.Println(fmt.Sprintf(c.selectedColor, str))

}

func (c *colorgo) Printf(str string) {
	fmt.Printf(c.selectedColor, str)
}
