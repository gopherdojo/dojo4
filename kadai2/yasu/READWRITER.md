## io.Writerとio.Readerについて

### io.Writer
#### どのように使われている？
io.Writerはfmtパッケージに置いて、FrintlnやFrintfなどに使われており、次のようになっています。

```print.go
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintln(a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

ここでは受け取ったinterface{}の配列をinterfaceの持つ情報の文字列に変換し、その文字列をbyte型のスライスにしたものをw(io.Writer)に書き込んでいます。

では、このio.Writerが何者かと言うと
Printlnの実装を見てみると

```print.go
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

os.Stdout、つまり標準出力が渡されています。
os.StdoutにはWriteメソッドが存在しておりWriteを発火させる事によって標準出力に渡されたbyte配列を書き込んでいるのです。
もちろん、Fprintlnにはos.Fileなどを渡すこともでき、io.Writerは書き込む先を抽象化するために使われてると言えます。

### io.Reader
#### どのように使われている？

fmtパッケージのScanlnを例にとって考えてみると　

```scan.go
func Fscan(r io.Reader, a ...interface{}) (n int, err error) {
	s, old := newScanState(r, true, false)
	n, err = s.doScan(a)
	s.free(old)
	return
}
```

io.Readerはここで使われています。
こちらでのReaderの使い方はFprintlnとほぼ一緒で

```scan.go
func Scan(a ...interface{}) (n int, err error) {
	return Fscan(os.Stdin, a...)
}
```

```scan.go
func newScanState(r io.Reader, nlIsSpace, nlIsEnd bool) (s *ss, old ssave) {
	s = ssFree.Get().(*ss)
	if rs, ok := r.(io.RuneScanner); ok {
		s.rs = rs
	} else {
		s.rs = &readRune{reader: r, peekRune: -1}
	}
	s.nlIsSpace = nlIsSpace
	s.nlIsEnd = nlIsEnd
	s.atEOF = false
	s.limit = hugeWid
	s.argLimit = hugeWid
	s.maxWid = hugeWid
	s.validSave = true
	s.count = 0
	return
}

```

標準入力から受け取った文字列をssという構造体に書き込んでいます。

こちらもos.Stdinやos.Fileなどの構造体を抽象化するために用いられています。


### 利点
利点としては抽象化する事により実際に受け取っている値に依存する事なく関数を実装できる事だと思います。
抽象化する事により受け取った具象系の変更の影響を受けないので変更にも強いです。
