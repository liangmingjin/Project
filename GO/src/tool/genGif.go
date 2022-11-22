// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.RGBA{255, 0, 0, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	//genPng()
	lissajous()
}
func getRbg(a uint) uint8 {
	b := a % 510
	if b > 255 {
		return uint8(510 - b)
	}
	return uint8(b)
}
func lissajous() {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 98    // image canvas covers [-size..+size]
		nframes = 255   // number of animation frames
		delay   = 1     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	red := uint8(rand.Int())
	green := uint8(rand.Int())
	//blue := uint8(rand.Int())
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		palette = []color.Color{color.White, color.RGBA{getRbg(uint(red) + uint(i)*2), getRbg(uint(green) + uint(i)*2), 0, 0xff}}
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	//outImage(anim.Image)
	buf := &bytes.Buffer{}
	_ = gif.EncodeAll(buf, &anim)
	if err := ioutil.WriteFile("test.gif", buf.Bytes(), 0666); err != nil {
		fmt.Println(err)
	}
}
func outImage(Image []*image.Paletted) {
	for k, v := range Image {
		buf := &bytes.Buffer{}
		_ = png.Encode(buf, v)
		if err := ioutil.WriteFile(fmt.Sprintf("image%d.png", k), buf.Bytes(), 0666); err != nil {
			fmt.Println(err)
			continue
		}

	}

}

func genPng() {
	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()
}
