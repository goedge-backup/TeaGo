package files

import (
	"os"
	"sync"
	"io"
)

type Writer struct {
	file   *os.File
	locker *sync.Mutex
}

func NewWriter(path string) (*Writer, error) {
	return NewFile(path).Writer()
}

func (this *Writer) WriteString(s string) (n int64, err error) {
	written, err := this.file.WriteString(s)
	n = int64(written)
	return
}

func (this *Writer) Write(b []byte) (n int64, err error) {
	written, err := this.file.Write(b)
	n = int64(written)
	return
}

func (this *Writer) WriteIOReader(reader io.Reader) (n int64, err error) {
	return io.Copy(this.file, reader)
}

func (this *Writer) Truncate(size ... int64) error {
	if len(size) > 0 {
		return this.file.Truncate(size[0])
	}
	return this.file.Truncate(0)
}

func (this *Writer) Seek(offset int64, whence ... int) (ret int64, err error) {
	if len(whence) > 0 {
		return this.file.Seek(offset, whence[0])
	}
	return this.file.Seek(offset, 0)
}

func (this *Writer) Sync() error {
	return this.file.Sync()
}

func (this *Writer) Lock() {
	this.locker.Lock()
}

func (this *Writer) Unlock() {
	this.locker.Unlock()
}

func (this *Writer) Close() error {
	return this.file.Close()
}
