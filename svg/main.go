package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/svgo"
)

func main() {
	width := 500
	height := 500
	file, err := os.Create("./out.svg")
	if err != nil {
		panic(err)
	}

	canvas := svg.New(file)
	canvas.Start(width, height)

	canvas.Def()
	fmt.Fprintf(canvas.Writer, "<style type=\"text/css\">\n<![CDATA[\n#test {fill:   #00cc00;}]]>\n</style>\n")
	canvas.DefEnd()

	canvas.Gid("test")
	canvas.Circle(width/2, height/2, 100)
	canvas.Gend()
	canvas.Roundrect(100, 100, 100, 200, 0, 0, "fill:red")
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}
