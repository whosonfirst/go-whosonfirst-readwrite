package reader

import (
	"bufio"
	"github.com/whosonfirst/go-whosonfirst-readwrite/bytes"
	"io"
	"os"
)

type StdinReader struct {
	Reader
}

func NewStdinReader() (Reader, error) {

	r := StdinReader{}
	return &r, nil
}

func (r *StdinReader) Read(uri string) (io.ReadCloser, error) {

	body := []byte()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		scanner.Bytes()
	}

	return bytes.ReadCloserFromBytes(body)
}

func (r *StdinReader) URI(uri string) string {
	return uri
}
