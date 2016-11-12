package main

import (
	"fmt"
	"image/png"
	"os"
)

func to_num(r, g, b, a uint32) byte {
	r = r >> 8
	g = g >> 8
	b = b >> 8
	//fmt.Printf("%d %d %d\n", r, g, b)
	if r < 128 { return 0x01; }
	return 0x03;
}

func main() {
	cnt := 0
	f, _ := os.Open("icons.png")
	defer f.Close()
//	f2, _ := os.Create("file.dat")
//	defer f2.Close()

	g, _ := png.Decode(f)
	// bounds := g.Bounds()

	buf := make([]byte, 1)

	fmt.Printf("data icons_img {\r\n")

	for y := 0; y < 112; y++ {
		for x := 0; x < 8; x+=4 {
			buf[0] = 0
			for ix := 0; ix < 4; ix++ {
				c := g.At(x+ix, y)
				cr, cg, cb, ca := c.RGBA()
				buf[0] = (buf[0] << 2) | to_num(cr,cg,cb,ca)
			}
//			f2.Write(buf)
			fmt.Printf("%d ", uint32(buf[0]) & 0xFF)
			cnt = cnt + 1
			if cnt >= 4000 {
				cnt = 0
				fmt.Printf("\r\n}\r\ndata _ {\r\n")
			}
		}
	}

	fmt.Printf("\r\n}\r\n")
}
