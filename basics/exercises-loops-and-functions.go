// Exercise: Loops and Functions
// 関数とループを使った簡単な練習として、平方根の実装をしてみましょう
package main

import (
	"fmt"
	"math"
)

// ニュートン法の繰り返しを中断するしきい値
const epsilon = 1.0e-5

// Sqrt は引数 x の平方根を求めます。
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		ztmp := z
		z -= (z*z - x) / (2 * z)
		// ニュートン法の変化量がしきい値を下回ったら繰り返しを中断し、その回数を出力する
		if math.Abs(z-ztmp) < epsilon {
			fmt.Printf("繰り返し回数: %v\n", i)
			break
		}
	}
	return z
}

func main() {
	x := float64(2)
	// mathライブラリのSqrt関数と精度を比較
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
