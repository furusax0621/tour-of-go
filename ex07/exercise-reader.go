// Exercise: Readers
// ASCII文字 'A' の無限ストリームを出力する Reader 型を実装してください。
package main

import "golang.org/x/tour/reader"

// MyReader は io.Reader を実装する演習用の独自構造体です。
type MyReader struct{}

func (z MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
