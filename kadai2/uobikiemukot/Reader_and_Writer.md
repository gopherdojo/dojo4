# io.Readerとio.Writer

## 標準パッケージでどのように使われているか

GOROOT/src以下をgrepを使って検索し，io.Reader/io.Writerを実装している型を調査した．

bufio.Reader，bufio.Writer，strings.Readr，io.Fileなど，  
良く知られているもの以外にも多くの型がio.Reader/io.Writerを実装していることがわかった．

特にexportされていない非公開の型についてもio.Reader/io.Writerを実装しているものが多数あった．  
MockTerminalやFakeFileなど，テスト用と思われるio.Reader/io.Writerに準拠した型も多い．

### Writeインタフェースを実装している型 (exportされているもの)

<details><summary>egrep -R "func \(.*\) Write\([a-z]+ \[\]byte\)" * | egrep "\([a-z]+ \*?[A-Z].*\)</summary>

~~~go
archive/tar/writer.go:func (tw *Writer) Write(b []byte) (int, error) {
bufio/bufio.go:func (b *Writer) Write(p []byte) (nn int, err error) {
bytes/buffer.go:func (b *Buffer) Write(p []byte) (n int, err error) {
cmd/go/internal/cache/hash.go:func (h *Hash) Write(b []byte) (int, error) {
cmd/go/internal/list/list.go:func (t *TrackingWriter) Write(p []byte) (n int, err error) {
cmd/link/internal/ld/outbuf.go:func (out *OutBuf) Write(v []byte) (int, error) {
cmd/vendor/golang.org/x/crypto/ssh/terminal/terminal.go:func (t *Terminal) Write(buf []byte) (n int, err error) {
cmd/vendor/golang.org/x/crypto/ssh/terminal/terminal_test.go:func (c *MockTerminal) Write(data []byte) (n int, err error) {
compress/flate/deflate.go:func (w *Writer) Write(data []byte) (n int, err error) {
compress/gzip/gzip.go:func (z *Writer) Write(p []byte) (int, error) {
compress/zlib/writer.go:func (z *Writer) Write(p []byte) (n int, err error) {
crypto/cipher/io.go:func (w StreamWriter) Write(src []byte) (n int, err error) {
crypto/tls/conn.go:func (c *Conn) Write(b []byte) (int, error) {
internal/poll/fd_windows.go:func (fd *FD) Write(buf []byte) (int, error) {
internal/poll/fd_unix.go:func (fd *FD) Write(p []byte) (int, error) {
io/pipe.go:func (w *PipeWriter) Write(data []byte) (n int, err error) {
log/syslog/syslog.go:func (w *Writer) Write(b []byte) (int, error) {
mime/quotedprintable/writer.go:func (w *Writer) Write(p []byte) (n int, err error) {
net/http/httptest/recorder.go:func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
os/file.go:func (f *File) Write(b []byte) (n int, err error) {
runtime/race/testdata/mop_test.go:func (d DummyWriter) Write(p []byte) (n int) {
strings/builder.go:func (b *Builder) Write(p []byte) (int, error) {
text/tabwriter/tabwriter.go:func (b *Writer) Write(buf []byte) (n int, err error) {
text/template/exec_test.go:func (e ErrorWriter) Write(p []byte) (int, error) {
vendor/golang_org/x/net/http2/hpack/hpack.go:func (d *Decoder) Write(p []byte) (n int, err error) {
vendor/golang_org/x/text/transform/transform.go:func (w *Writer) Write(data []byte) (n int, err error) {
~~~
</details>

### Readインタフェースを実装している型 (exportされているもの)

<details><summary>egrep -R "func \(.*\) Read\([a-z]+ \[\]byte\)" * | egrep "\([a-z]+ \*?[A-Z].*\)</summary>

~~~go
archive/tar/reader.go:func (tr *Reader) Read(b []byte) (int, error) {
bufio/bufio_test.go:func (r *StringReader) Read(p []byte) (n int, err error) {
bufio/bufio.go:func (b *Reader) Read(p []byte) (n int, err error) {
bytes/buffer.go:func (b *Buffer) Read(p []byte) (n int, err error) {
bytes/reader.go:func (r *Reader) Read(b []byte) (n int, err error) {
cmd/pack/pack_test.go:func (f *FakeFile) Read(p []byte) (int, error) {
cmd/vendor/golang.org/x/crypto/ssh/terminal/terminal_test.go:func (c *MockTerminal) Read(data []byte) (n int, err error) {
compress/gzip/gunzip.go:func (z *Reader) Read(p []byte) (n int, err error) {
crypto/cipher/io.go:func (r StreamReader) Read(dst []byte) (n int, err error) {
crypto/tls/conn.go:func (c *Conn) Read(b []byte) (n int, err error) {
internal/poll/fd_windows.go:func (fd *FD) Read(buf []byte) (int, error) {
internal/poll/fd_unix.go:func (fd *FD) Read(p []byte) (int, error) {
io/io.go:func (l *LimitedReader) Read(p []byte) (n int, err error) {
io/io.go:func (s *SectionReader) Read(p []byte) (n int, err error) {
io/pipe.go:func (r *PipeReader) Read(data []byte) (n int, err error) {
math/rand/rand.go:func (r *Rand) Read(p []byte) (n int, err error) {
mime/multipart/multipart.go:func (p *Part) Read(d []byte) (n int, err error) {
mime/quotedprintable/reader.go:func (r *Reader) Read(p []byte) (n int, err error) {
net/net.go:func (v *Buffers) Read(p []byte) (n int, err error) {
os/file.go:func (f *File) Read(b []byte) (n int, err error) {
strings/reader.go:func (r *Reader) Read(b []byte) (n int, err error) {
text/scanner/scanner_test.go:func (r *StringReader) Read(p []byte) (n int, err error) {
vendor/golang_org/x/text/transform/transform.go:func (r *Reader) Read(p []byte) (int, error) {
~~~
</details>

### Writeインタフェースを実装している型 (全て)

<details><summary>egrep -R "func \(.*\) Write\([a-z]+ \[\]byte\)" *</summary>

~~~go
archive/zip/register.go:func (w *pooledFlateWriter) Write(p []byte) (n int, err error) {
archive/zip/zip_test.go:func (r *rleBuffer) Write(p []byte) (n int, err error) {
archive/zip/zip_test.go:func (fakeHash32) Write(p []byte) (int, error) { return len(p), nil }
archive/zip/zip_test.go:func (ss *suffixSaver) Write(p []byte) (n int, err error) {
archive/zip/writer.go:func (dirWriter) Write(b []byte) (int, error) {
archive/zip/writer.go:func (w *fileWriter) Write(p []byte) (int, error) {
archive/zip/writer.go:func (w *countWriter) Write(p []byte) (int, error) {
archive/tar/writer_test.go:func (w *failOnceWriter) Write(b []byte) (int, error) {
archive/tar/writer_test.go:func (w testNonEmptyWriter) Write(b []byte) (int, error) {
archive/tar/writer.go:func (tw *Writer) Write(b []byte) (int, error) {
archive/tar/writer.go:func (fw *regFileWriter) Write(b []byte) (n int, err error) {
archive/tar/writer.go:func (sw *sparseFileWriter) Write(b []byte) (n int, err error) {
archive/tar/writer.go:func (zeroWriter) Write(b []byte) (int, error) {
archive/tar/tar_test.go:func (f *testFile) Write(b []byte) (int, error) {
bufio/bufio_test.go:func (w errorWriterTest) Write(p []byte) (int, error) {
bufio/bufio_test.go:func (w errorWriterToTest) Write(p []byte) (int, error) {
bufio/bufio_test.go:func (w errorReaderFromTest) Write(p []byte) (int, error) {
bufio/bufio_test.go:func (w *writeCountingDiscard) Write(p []byte) (int, error) {
bufio/bufio.go:func (b *Writer) Write(p []byte) (nn int, err error) {
bytes/buffer.go:func (b *Buffer) Write(p []byte) (n int, err error) {
cmd/go/internal/cache/hash.go:func (h *Hash) Write(b []byte) (int, error) {
cmd/go/internal/test/test.go:func (lockedStdout) Write(b []byte) (int, error) {
cmd/go/internal/list/list.go:func (t *TrackingWriter) Write(p []byte) (n int, err error) {
cmd/go/internal/help/help.go:func (c *commentWriter) Write(p []byte) (int, error) {
cmd/go/internal/help/help.go:func (w *errWriter) Write(b []byte) (int, error) {
cmd/trace/trace.go:func (cw *countingWriter) Write(data []byte) (int, error) {
cmd/internal/bio/must.go:func (w mustWriter) Write(b []byte) (int, error) {
cmd/internal/test2json/test2json.go:func (c *converter) Write(b []byte) (int, error) {
cmd/compile/internal/syntax/dumper.go:func (p *dumper) Write(data []byte) (n int, err error) {
cmd/link/internal/ld/outbuf.go:func (out *OutBuf) Write(v []byte) (int, error) {
cmd/test2json/main.go:func (w *countWriter) Write(b []byte) (int, error) {
cmd/vendor/golang.org/x/crypto/ssh/terminal/terminal.go:func (t *Terminal) Write(buf []byte) (n int, err error) {
cmd/vendor/golang.org/x/crypto/ssh/terminal/terminal_test.go:func (c *MockTerminal) Write(data []byte) (n int, err error) {
compress/flate/writer_test.go:func (e *errorWriter) Write(b []byte) (int, error) {
compress/flate/deflate.go:func (w *dictWriter) Write(b []byte) (n int, err error) {
compress/flate/deflate.go:func (w *Writer) Write(data []byte) (n int, err error) {
compress/flate/deflate_test.go:func (b *syncBuffer) Write(p []byte) (n int, err error) {
compress/flate/deflate_test.go:func (w *failWriter) Write(b []byte) (int, error) {
compress/gzip/gzip.go:func (z *Writer) Write(p []byte) (int, error) {
compress/gzip/gzip_test.go:func (l *limitedWriter) Write(p []byte) (n int, err error) {
compress/lzw/writer.go:func (e *encoder) Write(p []byte) (n int, err error) {
compress/zlib/writer.go:func (z *Writer) Write(p []byte) (n int, err error) {
crypto/sha512/sha512.go:func (d *digest) Write(p []byte) (nn int, err error) {
crypto/hmac/hmac.go:func (h *hmac) Write(p []byte) (n int, err error) {
crypto/md5/md5.go:func (d *digest) Write(p []byte) (nn int, err error) {
crypto/cipher/io.go:func (w StreamWriter) Write(src []byte) (n int, err error) {
crypto/tls/handshake_test.go:func (r *recordingConn) Write(b []byte) (n int, err error) {
crypto/tls/prf.go:func (h *finishedHash) Write(msg []byte) (n int, err error) {
crypto/tls/handshake_client_test.go:func (o *opensslOutputSink) Write(data []byte) (n int, err error) {
crypto/tls/handshake_client_test.go:func (b *brokenConn) Write(data []byte) (int, error) {
crypto/tls/handshake_client_test.go:func (wcc *writeCountingConn) Write(data []byte) (int, error) {
crypto/tls/cipher_suites.go:func (c *cthWrapper) Write(p []byte) (int, error) { return c.h.Write(p) }
crypto/tls/tls_test.go:func (w *changeImplConn) Write(p []byte) (n int, err error) {
crypto/tls/tls_test.go:func (c *slowConn) Write(p []byte) (int, error) {
crypto/tls/conn.go:func (c *Conn) Write(b []byte) (int, error) {
crypto/sha256/sha256.go:func (d *digest) Write(p []byte) (nn int, err error) {
crypto/sha1/sha1.go:func (d *digest) Write(p []byte) (nn int, err error) {
encoding/pem/pem.go:func (l *lineBreaker) Write(b []byte) (n int, err error) {
encoding/hex/hex.go:func (e *encoder) Write(p []byte) (n int, err error) {
encoding/hex/hex.go:func (h *dumper) Write(data []byte) (n int, err error) {
encoding/ascii85/ascii85.go:func (e *encoder) Write(p []byte) (n int, err error) {
encoding/base64/base64.go:func (e *encoder) Write(p []byte) (n int, err error) {
encoding/xml/xml_test.go:func (errWriter) Write(p []byte) (n int, err error) { return 0, fmt.Errorf("unwritable") }
encoding/xml/marshal_test.go:func (lw *limitedBytesWriter) Write(p []byte) (n int, err error) {
encoding/csv/writer_test.go:func (e errorWriter) Write(b []byte) (int, error) {
encoding/base32/base32.go:func (e *encoder) Write(p []byte) (n int, err error) {
encoding/gob/encode.go:func (e *encBuffer) Write(p []byte) (int, error) {
fmt/print.go:func (b *buffer) Write(p []byte) {
fmt/print.go:func (p *pp) Write(b []byte) (ret int, err error) {
go/printer/printer_test.go:func (l *limitWriter) Write(buf []byte) (n int, err error) {
go/printer/printer.go:func (p *trimmer) Write(data []byte) (n int, err error) {
go/ast/print.go:func (p *printer) Write(data []byte) (n int, err error) {
hash/adler32/adler32.go:func (d *digest) Write(p []byte) (nn int, err error) {
hash/crc32/crc32.go:func (d *digest) Write(p []byte) (n int, err error) {
hash/fnv/fnv.go:func (s *sum32) Write(data []byte) (int, error) {
hash/fnv/fnv.go:func (s *sum32a) Write(data []byte) (int, error) {
hash/fnv/fnv.go:func (s *sum64) Write(data []byte) (int, error) {
hash/fnv/fnv.go:func (s *sum64a) Write(data []byte) (int, error) {
hash/fnv/fnv.go:func (s *sum128) Write(data []byte) (int, error) {
hash/fnv/fnv.go:func (s *sum128a) Write(data []byte) (int, error) {
hash/crc64/crc64.go:func (d *digest) Write(p []byte) (n int, err error) {
image/png/writer.go:func (e *encoder) Write(b []byte) (int, error) {
image/gif/writer.go:func (b blockWriter) Write(data []byte) (int, error) {
internal/poll/fd_windows.go:func (fd *FD) Write(buf []byte) (int, error) {
internal/poll/fd_unix.go:func (fd *FD) Write(p []byte) (int, error) {
io/ioutil/ioutil.go:func (devNull) Write(p []byte) (int, error) {
io/multi_test.go:func (c *writeStringChecker) Write(p []byte) (n int, err error) {
io/multi_test.go:func (f writerFunc) Write(p []byte) (int, error) {
io/multi.go:func (t *multiWriter) Write(p []byte) (n int, err error) {
io/pipe.go:func (p *pipe) Write(b []byte) (n int, err error) {
io/pipe.go:func (w *PipeWriter) Write(data []byte) (n int, err error) {
io/io_test.go:func (w *noReadFrom) Write(p []byte) (n int, err error) {
log/syslog/syslog.go:func (w *Writer) Write(b []byte) (int, error) {
mime/multipart/writer.go:func (p *part) Write(d []byte) (n int, err error) {
mime/quotedprintable/writer.go:func (w *Writer) Write(p []byte) (n int, err error) {
net/fd_windows.go:func (fd *netFD) Write(buf []byte) (int, error) {
net/fd_plan9.go:func (fd *netFD) Write(b []byte) (n int, err error) {
net/fd_unix.go:func (fd *netFD) Write(p []byte) (nn int, err error) {
net/textproto/writer.go:func (d *dotWriter) Write(b []byte) (n int, err error) {
net/net_fake.go:func (fd *netFD) Write(p []byte) (nn int, err error) {
net/net_fake.go:func (p *bufferedPipe) Write(b []byte) (int, error) {
net/net.go:func (c *conn) Write(b []byte) (int, error) {
net/http/transport.go:func (w persistConnWriter) Write(p []byte) (n int, err error) {
net/http/httptest/recorder.go:func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
net/http/httputil/reverseproxy.go:func (m *maxLatencyWriter) Write(p []byte) (int, error) {
net/http/requestwrite_test.go:func (f writerFunc) Write(p []byte) (int, error) { return f(p) }
net/http/server.go:func (cw *chunkWriter) Write(p []byte) (n int, err error) {
net/http/server.go:func (w *response) Write(data []byte) (n int, err error) {
net/http/server.go:func (tw *timeoutWriter) Write(p []byte) (int, error) {
net/http/server.go:func (c *loggingConn) Write(p []byte) (n int, err error) {
net/http/server.go:func (w checkConnErrorWriter) Write(p []byte) (n int, err error) {
net/http/h2_bundle.go:func (b *http2dataBuffer) Write(p []byte) (int, error) {
net/http/h2_bundle.go:func (w *http2bufferedWriter) Write(p []byte) (n int, err error) {
net/http/h2_bundle.go:func (p *http2pipe) Write(d []byte) (n int, err error) {
net/http/h2_bundle.go:func (cw http2chunkWriter) Write(p []byte) (n int, err error) { return cw.rws.writeChunk(p) }
net/http/h2_bundle.go:func (w *http2responseWriter) Write(p []byte) (n int, err error) {
net/http/h2_bundle.go:func (sew http2stickyErrWriter) Write(p []byte) (n int, err error) {
net/http/request_test.go:func (l logWrites) Write(p []byte) (n int, err error) {
net/http/fcgi/child.go:func (r *response) Write(data []byte) (int, error) {
net/http/fcgi/fcgi.go:func (w *streamWriter) Write(p []byte) (int, error) {
net/http/fcgi/fcgi_test.go:func (c *writeOnlyConn) Write(p []byte) (int, error) {
net/http/serve_test.go:func (c *testConn) Write(b []byte) (int, error) {
net/http/serve_test.go:func (c *slowTestConn) Write(b []byte) (int, error) {
net/http/serve_test.go:func (w terrorWriter) Write(p []byte) (int, error) {
net/http/internal/chunked.go:func (cw *chunkedWriter) Write(data []byte) (n int, err error) {
net/http/cgi/matryoshka_test.go:func (r *customWriterRecorder) Write(p []byte) (n int, err error) {
net/http/cgi/matryoshka_test.go:func (w *limitWriter) Write(p []byte) (n int, err error) {
net/http/cgi/child.go:func (r *response) Write(p []byte) (n int, err error) {
net/http/clientserver_test.go:func (b *lockedBytesBuffer) Write(p []byte) (int, error) {
net/http/fs.go:func (w *countingWriter) Write(p []byte) (n int, err error) {
net/http/transport_test.go:func (c writerFuncConn) Write(p []byte) (n int, err error) { return c.write(p) }
net/http/transport_test.go:func (c *logWritesConn) Write(p []byte) (n int, err error) {
net/http/transport_test.go:func (c funcConn) Write(p []byte) (int, error) { return c.write(p) }
net/http/transport_test.go:func (f funcWriter) Write(p []byte) (int, error) { return f(p) }
net/http/client_test.go:func (w chanWriter) Write(p []byte) (n int, err error) {
net/http/client_test.go:func (c *writeCountingConn) Write(p []byte) (int, error) {
net/http/filetransport.go:func (pr *populateResponse) Write(p []byte) (n int, err error) {
net/rpc/server_test.go:func (writeCrasher) Write(p []byte) (int, error) {
net/pipe.go:func (p *pipe) Write(b []byte) (int, error) {
net/splice_test.go:func (srv *spliceTestServer) Write(b []byte) (int, error) {
net/dnsclient_unix_test.go:func (f *fakeDNSConn) Write(b []byte) (int, error) {
os/file.go:func (f *File) Write(b []byte) (n int, err error) {
os/exec/exec_test.go:func (w *badWriter) Write(data []byte) (int, error) {
os/exec/exec.go:func (w *prefixSuffixSaver) Write(p []byte) (n int, err error) {
runtime/race/testdata/mop_test.go:func (d DummyWriter) Write(p []byte) (n int) {
strings/replace_test.go:func (errWriter) Write(p []byte) (n int, err error) {
strings/builder.go:func (b *Builder) Write(p []byte) (int, error) {
strings/replace.go:func (w *appendSliceWriter) Write(p []byte) (int, error) {
testing/benchmark.go:func (discard) Write(b []byte) (n int, err error) { return len(b), nil }
testing/sub_test.go:func (fw funcWriter) Write(b []byte) (int, error) { return fw(b) }
testing/iotest/logger.go:func (l *writeLogger) Write(p []byte) (n int, err error) {
testing/iotest/writer.go:func (t *truncateWriter) Write(p []byte) (n int, err error) {
testing/testing.go:func (w indenter) Write(b []byte) (n int, err error) {
text/tabwriter/tabwriter_test.go:func (b *buffer) Write(buf []byte) (written int, err error) {
text/tabwriter/tabwriter.go:func (b *Writer) Write(buf []byte) (n int, err error) {
text/template/exec_test.go:func (e ErrorWriter) Write(p []byte) (int, error) {
vendor/golang_org/x/net/http2/hpack/hpack.go:func (d *Decoder) Write(p []byte) (n int, err error) {
vendor/golang_org/x/text/unicode/norm/readwriter.go:func (w *normWriter) Write(data []byte) (n int, err error) {
vendor/golang_org/x/text/transform/transform.go:func (w *Writer) Write(data []byte) (n int, err error) {
~~~
</details>

### Readインタフェースを実装している型 (全て)

<details><summary>egrep -R "func \(.*\) Read\([a-z]+ \[\]byte\)" *</summary>

~~~go
archive/zip/register.go:func (r *pooledFlateReader) Read(p []byte) (n int, err error) {
archive/zip/zip_test.go:func (zeros) Read(p []byte) (int, error) {
archive/zip/reader.go:func (r *checksumReader) Read(b []byte) (n int, err error) {
archive/zip/reader_test.go://	func (zeros) Read(b []byte) (int, error) {
archive/tar/tar_test.go:func (f *testFile) Read(b []byte) (int, error) {
archive/tar/reader.go:func (tr *Reader) Read(b []byte) (int, error) {
archive/tar/reader.go:func (fr *regFileReader) Read(b []byte) (n int, err error) {
archive/tar/reader.go:func (sr *sparseFileReader) Read(b []byte) (n int, err error) {
archive/tar/reader.go:func (zeroReader) Read(b []byte) (int, error) {
archive/tar/reader_test.go:func (r testNonEmptyReader) Read(b []byte) (int, error) {
bufio/bufio_test.go:func (r13 *rot13Reader) Read(p []byte) (int, error) {
bufio/bufio_test.go:func (zeroReader) Read(p []byte) (int, error) {
bufio/bufio_test.go:func (r *StringReader) Read(p []byte) (n int, err error) {
bufio/bufio_test.go:func (r dataAndEOFReader) Read(p []byte) (int, error) {
bufio/bufio_test.go:func (t *testReader) Read(buf []byte) (n int, err error) {
bufio/bufio_test.go:func (r errorWriterToTest) Read(p []byte) (int, error) {
bufio/bufio_test.go:func (r errorReaderFromTest) Read(p []byte) (int, error) {
bufio/bufio_test.go:func (r *errorThenGoodReader) Read(p []byte) (int, error) {
bufio/bufio_test.go:func (r *emptyThenNonEmptyReader) Read(p []byte) (int, error) {
bufio/bufio_test.go:func (sr *scriptedReader) Read(p []byte) (n int, err error) {
bufio/bufio.go:func (b *Reader) Read(p []byte) (n int, err error) {
bufio/scan_test.go:func (sr *slowReader) Read(p []byte) (n int, err error) {
bufio/scan_test.go:func (alwaysError) Read(p []byte) (int, error) {
bufio/scan_test.go:func (endlessZeros) Read(p []byte) (int, error) {
bytes/buffer_test.go:func (r panicReader) Read(p []byte) (int, error) {
bytes/buffer.go:func (b *Buffer) Read(p []byte) (n int, err error) {
bytes/reader.go:func (r *Reader) Read(b []byte) (n int, err error) {
cmd/pack/pack_test.go:func (f *FakeFile) Read(p []byte) (int, error) {
cmd/vendor/golang.org/x/crypto/ssh/terminal/terminal_test.go:func (c *MockTerminal) Read(data []byte) (n int, err error) {
cmd/vendor/golang.org/x/crypto/ssh/terminal/util.go:func (r passwordReader) Read(buf []byte) (int, error) {
compress/flate/inflate.go:func (f *decompressor) Read(b []byte) (int, error) {
compress/flate/deflate_test.go:func (r *sparseReader) Read(b []byte) (n int, err error) {
compress/flate/deflate_test.go:func (b *syncBuffer) Read(p []byte) (n int, err error) {
compress/gzip/gunzip.go:func (z *Reader) Read(p []byte) (n int, err error) {
compress/lzw/reader.go:func (d *decoder) Read(b []byte) (int, error) {
compress/lzw/reader_test.go:func (devZero) Read(p []byte) (int, error) {
compress/zlib/reader.go:func (z *reader) Read(p []byte) (int, error) {
compress/bzip2/bzip2.go:func (bz2 *reader) Read(buf []byte) (n int, err error) {
crypto/ecdsa/ecdsa.go:func (z *zr) Read(dst []byte) (n int, err error) {
crypto/cipher/io.go:func (r StreamReader) Read(dst []byte) (n int, err error) {
crypto/tls/example_test.go:func (zeroSource) Read(b []byte) (n int, err error) {
crypto/tls/handshake_test.go:func (r *recordingConn) Read(b []byte) (n int, err error) {
crypto/tls/handshake_server_test.go:func (zeroSource) Read(b []byte) (n int, err error) {
crypto/tls/handshake_client_test.go:func (i opensslInput) Read(buf []byte) (n int, err error) {
crypto/tls/conn.go:func (b *block) Read(p []byte) (n int, err error) {
crypto/tls/conn.go:func (c *Conn) Read(b []byte) (n int, err error) {
crypto/rand/rand_unix.go:func (r *devReader) Read(b []byte) (n int, err error) {
crypto/rand/rand_unix.go:func (hr hideAgainReader) Read(p []byte) (n int, err error) {
crypto/rand/rand_unix.go:func (r *reader) Read(b []byte) (n int, err error) {
crypto/rand/util_test.go:func (r *countingReader) Read(p []byte) (n int, err error) {
crypto/rand/rand_js.go:func (r *reader) Read(b []byte) (int, error) {
crypto/rand/rand_windows.go:func (r *rngReader) Read(b []byte) (n int, err error) {
debug/elf/reader.go:func (r errorReader) Read(p []byte) (n int, err error) {
debug/elf/reader.go:func (r *readSeekerFromReader) Read(p []byte) (n int, err error) {
encoding/hex/hex.go:func (d *decoder) Read(p []byte) (n int, err error) {
encoding/ascii85/ascii85.go:func (d *decoder) Read(p []byte) (n int, err error) {
encoding/base64/base64.go:func (d *decoder) Read(p []byte) (n int, err error) {
encoding/base64/base64.go:func (r *newlineFilteringReader) Read(p []byte) (int, error) {
encoding/base64/base64_test.go:func (r *faultInjectReader) Read(p []byte) (int, error) {
encoding/xml/xml_test.go:func (d *downCaser) Read(p []byte) (int, error) {
encoding/csv/reader_test.go:func (r *nTimes) Read(p []byte) (n int, err error) {
encoding/binary/binary_test.go:func (br *byteSliceReader) Read(p []byte) (int, error) {
encoding/base32/base32_test.go:func (b *badReader) Read(p []byte) (int, error) {
encoding/base32/base32.go:func (d *decoder) Read(p []byte) (n int, err error) {
encoding/base32/base32.go:func (r *newlineFilteringReader) Read(p []byte) (int, error) {
encoding/gob/timing_test.go:func (b *benchmarkBuf) Read(p []byte) (n int, err error) {
encoding/gob/debug.go:func (p *peekReader) Read(b []byte) (n int, err error) {
encoding/gob/decode.go:func (d *decBuffer) Read(p []byte) (int, error) {
fmt/scan.go:func (r *stringReader) Read(b []byte) (n int, err error) {
fmt/scan.go:func (s *ss) Read(buf []byte) (n int, err error) {
fmt/scan_test.go:func (ec *eofCounter) Read(b []byte) (n int, err error) {
go/internal/gccgoimporter/testdata/escapeinfo.go:func (*T) Read(p []byte) {}
image/png/reader.go:func (d *decoder) Read(p []byte) (int, error) {
image/jpeg/reader_test.go:func (r *eofReader) Read(b []byte) (n int, err error) {
image/gif/reader.go:func (b *blockReader) Read(p []byte) (int, error) {
internal/poll/fd_windows.go:func (fd *FD) Read(buf []byte) (int, error) {
internal/poll/fd_unix.go:func (fd *FD) Read(p []byte) (int, error) {
io/multi_test.go:func (f readerFunc) Read(p []byte) (int, error) {
io/multi_test.go:func (b byteAndEOFReader) Read(p []byte) (n int, err error) {
io/io.go:func (l *LimitedReader) Read(p []byte) (n int, err error) {
io/io.go:func (s *SectionReader) Read(p []byte) (n int, err error) {
io/io.go:func (t *teeReader) Read(p []byte) (n int, err error) {
io/multi.go:func (mr *multiReader) Read(p []byte) (n int, err error) {
io/pipe.go:func (p *pipe) Read(b []byte) (n int, err error) {
io/pipe.go:func (r *PipeReader) Read(data []byte) (n int, err error) {
io/io_test.go:func (r zeroErrReader) Read(p []byte) (int, error) {
io/io_test.go:func (wantedAndErrReader) Read(p []byte) (int, error) {
io/io_test.go:func (r *dataAndErrorBuffer) Read(p []byte) (n int, err error) {
math/rand/rand.go:func (r *Rand) Read(p []byte) (n int, err error) {
mime/multipart/multipart_test.go:func (mr *maliciousReader) Read(b []byte) (n int, err error) {
mime/multipart/multipart_test.go:func (s *slowReader) Read(p []byte) (int, error) {
mime/multipart/multipart.go:func (r *stickyErrorReader) Read(p []byte) (n int, _ error) {
mime/multipart/multipart.go:func (p *Part) Read(d []byte) (n int, err error) {
mime/multipart/multipart.go:func (pr partReader) Read(d []byte) (int, error) {
mime/multipart/formdata_test.go:func (r *failOnReadAfterErrorReader) Read(p []byte) (n int, err error) {
mime/quotedprintable/reader.go:func (r *Reader) Read(p []byte) (n int, err error) {
net/fd_windows.go:func (fd *netFD) Read(buf []byte) (int, error) {
net/fd_plan9.go:func (fd *netFD) Read(b []byte) (n int, err error) {
net/fd_unix.go:func (fd *netFD) Read(p []byte) (n int, err error) {
net/textproto/reader.go:func (d *dotReader) Read(b []byte) (n int, err error) {
net/net_fake.go:func (fd *netFD) Read(p []byte) (n int, err error) {
net/net_fake.go:func (p *bufferedPipe) Read(b []byte) (int, error) {
net/timeout_test.go:func (b neverEnding) Read(p []byte) (int, error) {
net/net.go:func (c *conn) Read(b []byte) (int, error) {
net/net.go:func (v *Buffers) Read(p []byte) (n int, err error) {
net/http/transport.go:func (pc *persistConn) Read(p []byte) (n int, err error) {
net/http/transport.go:func (es *bodyEOFSignal) Read(p []byte) (n int, err error) {
net/http/transport.go:func (gz *gzipReader) Read(p []byte) (n int, err error) {
net/http/httputil/dump.go:func (b neverEnding) Read(p []byte) (n int, err error) {
net/http/httputil/dump.go:func (r *delegateReader) Read(p []byte) (int, error) {
net/http/httputil/reverseproxy_test.go:func (cc *checkCloser) Read(b []byte) (int, error) {
net/http/requestwrite_test.go:func (r *delegateReader) Read(p []byte) (int, error) {
net/http/server.go:func (cr *connReader) Read(p []byte) (n int, err error) {
net/http/server.go:func (ecr *expectContinueReader) Read(p []byte) (n int, err error) {
net/http/server.go:func (c *loggingConn) Read(p []byte) (n int, err error) {
net/http/h2_bundle.go:func (b *http2dataBuffer) Read(p []byte) (int, error) {
net/http/h2_bundle.go:func (p *http2pipe) Read(d []byte) (n int, err error) {
net/http/h2_bundle.go:func (b *http2requestBody) Read(p []byte) (n int, err error) {
net/http/h2_bundle.go:func (b http2transportResponseBody) Read(p []byte) (n int, err error) {
net/http/h2_bundle.go:func (gz *http2gzipReader) Read(p []byte) (n int, err error) {
net/http/h2_bundle.go:func (r http2errorReader) Read(p []byte) (int, error) { return 0, r.err }
net/http/request_test.go:func (dr delayedEOFReader) Read(p []byte) (n int, err error) {
net/http/request_test.go:func (r *infiniteReader) Read(b []byte) (int, error) {
net/http/request.go:func (l *maxBytesReader) Read(p []byte) (n int, err error) {
net/http/fcgi/fcgi_test.go:func (c *writeOnlyConn) Read(p []byte) (int, error) {
net/http/serve_test.go:func (c *testConn) Read(b []byte) (int, error) {
net/http/serve_test.go:func (c *slowTestConn) Read(b []byte) (n int, err error) {
net/http/serve_test.go:func (b neverEnding) Read(p []byte) (n int, err error) {
net/http/serve_test.go:func (cr countReader) Read(p []byte) (n int, err error) {
net/http/serve_test.go:func (r *repeatReader) Read(p []byte) (n int, err error) {
net/http/client.go:func (b *cancelTimerBody) Read(p []byte) (n int, err error) {
net/http/roundtrip_js.go:func (r *streamReader) Read(p []byte) (n int, err error) {
net/http/roundtrip_js.go:func (r *arrayReader) Read(p []byte) (n int, err error) {
net/http/cgi/matryoshka_test.go:func (b neverEnding) Read(p []byte) (n int, err error) {
net/http/clientserver_test.go:func (r testErrorReader) Read(p []byte) (n int, err error) {
net/http/transfer.go:func (r errorReader) Read(p []byte) (n int, err error) {
net/http/transfer.go:func (br *byteReader) Read(p []byte) (n int, err error) {
net/http/transfer.go:func (br transferBodyReader) Read(p []byte) (n int, err error) {
net/http/transfer.go:func (b *body) Read(p []byte) (n int, err error) {
net/http/transfer.go:func (bl bodyLocked) Read(p []byte) (n int, err error) {
net/http/transfer.go:func (fr finishAsyncByteRead) Read(p []byte) (n int, err error) {
net/http/transport_test.go:func (c byteFromChanReader) Read(p []byte) (n int, err error) {
net/http/transport_test.go:func (e errorReader) Read(p []byte) (int, error) { return 0, e.err }
net/http/transport_test.go:func (c *logWritesConn) Read(p []byte) (n int, err error) {
net/http/transport_test.go:func (c funcConn) Read(p []byte) (int, error)  { return c.read(p) }
net/http/client_test.go:func (f eofReaderFunc) Read(p []byte) (n int, err error) {
net/rpc/server_test.go:func (w *writeCrasher) Read(p []byte) (int, error) {
net/pipe.go:func (p *pipe) Read(b []byte) (int, error) {
net/splice_test.go:func (srv *spliceTestServer) Read(b []byte) (int, error) {
net/dnsclient_unix_test.go:func (f *fakeDNSConn) Read(b []byte) (int, error) {
os/timeout_test.go:func (b neverEnding) Read(p []byte) (int, error) {
os/file.go:func (f *File) Read(b []byte) (n int, err error) {
os/exec/exec_test.go:func (delayedInfiniteReader) Read(b []byte) (int, error) {
strings/reader.go:func (r *Reader) Read(b []byte) (n int, err error) {
testing/iotest/logger.go:func (l *readLogger) Read(p []byte) (n int, err error) {
testing/iotest/reader.go:func (r *oneByteReader) Read(p []byte) (int, error) {
testing/iotest/reader.go:func (r *halfReader) Read(p []byte) (int, error) {
testing/iotest/reader.go:func (r *dataErrReader) Read(p []byte) (n int, err error) {
testing/iotest/reader.go:func (r *timeoutReader) Read(p []byte) (int, error) {
text/scanner/scanner_test.go:func (r *StringReader) Read(p []byte) (n int, err error) {
text/scanner/scanner_test.go:func (errReader) Read(b []byte) (int, error) {
vendor/golang_org/x/text/unicode/norm/readwriter.go:func (r *normReader) Read(p []byte) (int, error) {
vendor/golang_org/x/text/transform/transform.go:func (r *Reader) Read(p []byte) (int, error) {
~~~
</details>

## io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

io.Readerやio.Writerといったインタフェースを引数に取る関数を作ることで，  
インタフェースに準拠した色々な型でその関数を利用できるようになる．

例えばos.Fileとbufio.ReadWriterはどちらもio.Reader/io.Writerに準拠しているため，  
io.Copy(Writer, Reader)などに共通して渡すことができる．

また，テストのために元々の型の性質を変更して，新たな型として定義している利用例があった．  
(下記のslowReaderやlockedStdoutなど)

### ioutil.devNull

ioutilにはdevNullという型が定義されている．  
io.Writerを実装しているが，実質的には何もしない(入力を捨てる)．

io.Copyと一緒に使って入力を読み捨てるのに使われるが，  
io.Writerを引数に取る関数には何にでも渡すことができる．

色々と使い道がありそう．

~~~go
type devNull int

// devNull implements ReaderFrom as an optimization so io.Copy to
// ioutil.Discard can avoid doing unnecessary work.
var _ io.ReaderFrom = devNull(0)

func (devNull) Write(p []byte) (int, error) {
        return len(p), nil
}

func (devNull) WriteString(s string) (int, error) {
        return len(s), nil
}

var blackHolePool = sync.Pool{
        New: func() interface{} {
                b := make([]byte, 8192)
                return &b
        },
}

func (devNull) ReadFrom(r io.Reader) (n int64, err error) {
        bufp := blackHolePool.Get().(*[]byte)
        readSize := 0
        for {
                readSize, err = r.Read(*bufp)
                n += int64(readSize)
                if err != nil {
                        blackHolePool.Put(bufp)
                        if err == io.EOF {
                                return n, nil
                        }
                        return
                }
        }
}

// Discard is an io.Writer on which all Write calls succeed
// without doing anything.
var Discard io.Writer = devNull(0)
~~~

### bufio/scan_test.go

bufio/scan_test.goでは，io.Readerを実装したslowReaderを新たに定義しており，  
意図的に遅い読み込みをシミュレートしたテストを書いている．

元々のio.Readerに少し手を加えるだけで実装できていて，インタフェースの利点が良くわかる．

~~~go
// slowReader is a reader that returns only a few bytes at a time, to test the incremental
// reads in Scanner.Scan.
type slowReader struct {
        max int
        buf io.Reader
}

func (sr *slowReader) Read(p []byte) (n int, err error) {
        if len(p) > sr.max {
                p = p[0:sr.max]
        }
        return sr.buf.Read(p)
}
~~~

### cmd/go/internal/test/test.go

os.Stdoutに関する書き込みを同期的にするlockedStdoutという型が用意されている．  
元々あるos.Stdoutの性質をテストのために少し変更している．

~~~go
// stdoutMu and lockedStdout provide a locked standard output
// that guarantees never to interlace writes from multiple
// goroutines, so that we can have multiple JSON streams writing
// to a lockedStdout simultaneously and know that events will
// still be intelligible.
var stdoutMu sync.Mutex

type lockedStdout struct{}

func (lockedStdout) Write(b []byte) (int, error) {
        stdoutMu.Lock()
        defer stdoutMu.Unlock()
        return os.Stdout.Write(b)
}
~~~
