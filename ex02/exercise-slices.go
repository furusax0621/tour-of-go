// Exercise: Slices
// Pic 関数を実装してみましょう。このプログラムを実行すると、生成した画像が表示されるはずです。
// この関数は、長さ dy のsliceに、各要素が8bitのunsigned int型で長さ dx のsliceを割り当てたものを返すように実装する必要があります。
// 画像は、整数値をグレースケール(実際はブルースケール)として解釈したものです。
package main

import "golang.org/x/tour/pic"

// Pic は長さdx * dyの画像を生成します。
func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := range slice {
		slice[i] = make([]uint8, dx)
	}

	for y := range slice {
		for x := range slice[y] {
			// x と y の値を使って画像を生成する
			slice[y][x] = uint8((x + y) / 2)
			//slice[y][x] = uint8(x * y)
			//slice[y][x] = uint8(x^y)
		}
	}

	return slice
}

func main() {
	pic.Show(Pic)
}
