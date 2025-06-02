/*
Author: Fu Huizhong <fuhuizn@163.com>
2025-06-02
*/
package goZbarScanner

/*
#cgo pkg-config: zbar
#include <stdlib.h>
extern char *zbar_scan(void *image_data,int width,int height, char *fmt);
*/
import "C"

import (
	"errors"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"unsafe"
)

var Fail = errors.New("Fail")

func ScanFile(name string) (string, error) {
	fp, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	img, _, err := image.Decode(fp)
	if err != nil {
		return "", err
	}

	gray := getGray(img)

	p := gray.Pix
	w := gray.Bounds().Size().X
	h := gray.Bounds().Size().Y

	f := C.CString("Y800")
	defer C.free(unsafe.Pointer(f))
	ret := C.zbar_scan(unsafe.Pointer(&p[0]), C.int(w), C.int(h), f)
	if uintptr(unsafe.Pointer(ret)) == 0 {
		return "", Fail
	}
	defer C.free(unsafe.Pointer(ret))
	return C.GoString(ret), nil
}

func getGray(img image.Image) *image.Gray {
	W := img.Bounds().Size().X
	H := img.Bounds().Size().Y
	ret := image.NewGray(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: W, Y: H}})
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			r, g, b, _ := img.At(j, i).RGBA()
			// 使用加权平均方法计算灰度值
			gray := uint8((0.2989*float64(r) + 0.5870*float64(g) + 0.1140*float64(b)) / 256)
			ret.Set(j, i, color.Gray{Y: gray})
		}
	}
	return ret

}
