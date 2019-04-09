// Exercise: Errors
// Sqrt 関数を 以前の演習 からコピーし、 error の値を返すように修正してみてください。
package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt は Sqrt に負の値が与えられた場合のエラーを規定します。
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + fmt.Sprint(float64(e))
}

// ニュートン法の繰り返しを中断するしきい値
const epsilon = 1.0e-5

// Sqrt は引数 x の平方根を求めます。
func Sqrt(x float64) (float64, error) {
	// 負の値が与えられた時、ErrNegativeSqrtを返す
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

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
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
