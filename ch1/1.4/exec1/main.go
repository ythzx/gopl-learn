package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var Red = color.RGBA{255, 0, 0, 255}
var Green = color.RGBA{0, 255, 0, 255}
var Blue = color.RGBA{0, 0, 255, 255}

// 调色板 初始化数组  复合类型，生成切片
// color.RGBA{R G B A uint8}
var palette = []color.Color{color.White, color.Black, Red, Green, Blue}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	redIndex   = 2
	greenIndex = 3
	blueIndex  = 4
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions x轴谐波数量
		res     = 0.001 //angular resolutions  角分辨率
		size    = 100   // image canvas 图像画布打小
		nframes = 100   // nubmer of animations frames 动画帧
		delay   = 8
	)

	freq := rand.Float64() * 3.0        // y轴频率
	anim := gif.GIF{LoopCount: nframes} // 复合类型，生成结构体
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.00; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
			// img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}
