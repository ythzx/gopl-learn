//将之前的lissajous 输出到web
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black} // 调色板 初始化数组  复合类型，生成切片

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/test", handler)
	// 下面是通过匿名函数的方式
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	lissajous(w)
	// })
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// fmt.Fprintf(w, "form is %v\n", r.Form) // 需要注释 会导致生成文件

	if r.Form["cycle"] != nil {
		cycle, err := strconv.Atoi(r.Form["cycle"][0])
		if err != nil {
			fmt.Println("convert err")
		}
		lissajous(w, float64(cycle))

	} else {
		lissajous(w, float64(0))
	}
}

func lissajous(out io.Writer, newCycle float64) {
	// number of complete x oscillator revolutions x轴谐波数量
	const (
		cycle   = 5
		res     = 0.001 //angular resolutions  角分辨率
		size    = 100   // image canvas 图像画布打小
		nframes = 100   // nubmer of animations frames 动画帧
		delay   = 8
	)

	if newCycle == float64(0) {
		newCycle = cycle
	}

	freq := rand.Float64() * 3.0        // y轴频率
	anim := gif.GIF{LoopCount: nframes} // 复合类型，生成结构体
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.00; t < newCycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}
