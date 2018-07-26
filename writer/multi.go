package writer

import (
	"github.com/whosonfirst/go-whosonfirst-readwrite/bytes"
	"io"
)

type MultiWriter struct {
	Writer
	writers []Writer
}

func NewMultiWriter(writers ...Writer) (Writer, error) {

	w := MultiWriter{
		writers: writers,
	}

	return &w, nil
}

func (w *MultiWriter) Write(path string, fh io.ReadCloser) error {

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return err
	}

	// please make this concurrent with a cancel context

	for _, wr := range w.writers {

		reader := bytes.ReadCloserFromBytes(body)

		err = wr.Write(path, reader)

		if err != nil {
			return err
		}
	}

	return nil
}
