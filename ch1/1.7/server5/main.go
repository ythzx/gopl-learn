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
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	if r.Form["cycle"] != nil {
		// cycle := 5
		// nweCycle, err := strconv.Atoi(r.Form["cycle"][0])
		// if err != nil {
		// 	fmt.Println("convert err")
		// }
		// cycle = nweCycle
		lissajous(w, r.Form["cycle"][0])
	}

}

func lissajous(out io.Writer, cycle_num string) {
	const (
		res     = 0.001 //angular resolutions  角分辨率
		size    = 1000  // image canvas 图像画布打小
		nframes = 100   // nubmer of animations frames 动画帧
		delay   = 8
	)

	nweCycle, err := strconv.Atoi(cycle_num)
	if err != nil {
		fmt.Println("convert err")
	}

	freq := rand.Float64() * 3.0        // y轴频率
	anim := gif.GIF{LoopCount: nframes} // 复合类型，生成结构体
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.00; t < float64(nweCycle)*2*math.Pi; t += res {
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
