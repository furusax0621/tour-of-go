// Exercise: rot13Reader
// よくあるパターンは、別の io.Reader をラップし、ストリームの内容を何らかの方法で変換するio.Readerです。
// io.Reader を実装し、 io.Reader でROT13 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用して読み出すように rot13Reader を実装してみてください。
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (z rot13Reader) Read(p []byte) (int, error) {
	num, err := z.r.Read(p)
	if err != nil {
		return num, err
	}

	for i := range p {
		if 'a' <= p[i] && p[i] <= 'z' {
			p[i] = (p[i]-'a'+13)%26 + 'a'
		} else if 'A' <= p[i] && p[i] <= 'Z' {
			p[i] = (p[i]-'A'+13)%26 + 'A'
		}
	}

	return len(p), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
