// Exercise: Images
// 前に解いた、画像ジェネレーターを覚えていますか？ 今回は、データのスライスの代わりに image.Image インタフェースの実装を返すようにしてみましょう。
// 自分の Image 型を定義し、 インタフェースを満たすのに必要なメソッド を実装し、 pic.ShowImage を呼び出してみてください。
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image はimage.Image インタフーフェースを実装する演習用の構造体です。
type Image struct{}

// ColorModel は img のカラーモデルを返します。
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds は img の画素サイズを返します。サイズは 200 * 200 固定です。
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

// At は img 内の任意の座標における色(color.RGBA)を返します。
func (img Image) At(x, y int) color.Color {
	// x と y の値から色を決定する
	v := uint8(((x + y) / 2) % 256)
	//v := uint8((x * y) % 256)
	//v := uint8((x ^ y) % 256)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
