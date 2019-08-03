package svg

import (
	"fmt"
	"io"
	cf "maths"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCenterX     = 150
	clockCenterY     = 150
)

func Write(w io.Writer, t time.Time)  {
	_, err := io.WriteString(w, svgStart)
	if err != nil {
		panic(err)
	}
	_, err = io.WriteString(w, bezel)
	if err != nil {
		panic(err)
	}

	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)

	_, err = io.WriteString(w, svgEnd)
	if err != nil {
		panic(err)
	}
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(cf.SecondHandPoint(t), secondHandLength)

	_, err := fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
	if err != nil {
		panic(err)
	}
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(cf.MinuteHandPoint(t), minuteHandLength)

	_, err := fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
	if err != nil {
		panic(err)
	}
}


func hourHand(w io.Writer, t time.Time) {
	p := makeHand(cf.HourHandPoint(t), hourHandLength)

	_, err := fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
	if err != nil {
		panic(err)
	}

}

func makeHand(p cf.Point, length float64) cf.Point {
	p = cf.Point{X: p.X * length, Y: p.Y * length}
	p = cf.Point{X: p.X, Y: -p.Y}
	return cf.Point{X: p.X + clockCenterX, Y: p.Y + clockCenterY}
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd  = `</svg>`
