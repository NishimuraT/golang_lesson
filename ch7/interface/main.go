package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"syscall"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}
// インターフェイスの埋め込み
type ReadWriter interface {
	Reader
	io.Writer
}

// ## nilポインタを含むインターフェイスは、nilではない
func f(out io.Writer) {
	if out != nil {
		// nilポインタだと実行してしまう
		out.Write([]byte("done!\n"))// パニック、nilポインタを参照
	}
}

func main() {
	var w io.Writer
	w = os.Stdout
	// func (b *Buffer) Write(p []byte) (n int, err error) {
	//w = bytes.Buffer{} // errorになる。*BufferがWriteメソッドを持っており、BufferはWriteメソッドを持っていないから。
	w = &bytes.Buffer{}
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)

	// ## nilポインタを含むインターフェイスは、nilではない
	var buf1 *bytes.Buffer// nilポインタ
	f(buf1)
	// 解決策、nilポインタを使わない
	var buf2 io.Writer
	debug := true
	if debug {
		buf2 = new(bytes.Buffer)
	}
	f(buf2)

	// ## 型アサーション(x.(T)この時、xはインターフェイス型、Tは断定(アサーション)型)
	// パターン１、Tが具象型の場合
	var w2 io.Writer
	w2 = os.Stdout
	f := w2.(*os.File)
	//c := w2.(*bytes.Buffer)// パニック、wは、bytes.Bufferではない
	fmt.Printf("%T %T \n", w2, f)
	// パターン２、Tがインターフェイスの場合、具象型のxがTを満足するか確認します
	var w3 io.Writer
	w3 = os.Stdout
	rw := w3.(io.ReadWriter)
	fmt.Printf("%T \n", rw)
	// okの使用
	if f2, ok := w2.(*os.File); ok {
		fmt.Println(f2)
	}
	// 型アサーションを使って、より信頼性が高いエラー確認
	_, err := os.Open("")
	fmt.Println(IsNotExist(err)) // true or false

	// ## 型アサーションによる、振る舞いの問い合わせ
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w2.(stringWriter); ok {
		sw.WriteString("")
	} else {
		w2.Write([]byte(""))// 文字列リテラルは、メモリを割り当てますが、その直後に破棄される。効率的ではない場合がある
	}

	// ## 型アサーションによる、判別共用体
	sqlQuote("a")
	sqlQuote(1)
	// 例（encoding/xml）
	xmlParse()
}

var ErrNotExist = errors.New("file does not exist")

type PathError struct {
	Op string
	Path string
	Err error
}
func (e *PathError)Error() string {
	return e.Op + " " + e.Path + ": "+ e.Err.Error()
}

// 型アサーションを使って、より信頼性が高いエラー確認
func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func xmlParse() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		tok, err := dec.Token()
		if err != io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)// プッシュ
		case xml.EndElement:
			stack = stack[:len(stack) - 1]// ポップ
		case xml.CharData:
			//if containsAll()割愛
		}
	}
}