package writer

import (
	"io"
	"io/ioutil"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

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

func (w *MultiWriter) Writer(path string, fh io.Reader) error {

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return err
	}

	// please make this concurrent with a cancel context
	
	for _, wr := range w.writers {

		buf := bytes.NewReader(body)
		reader := nopCloser{buf}

		err = wr.Write(path, reader)

		if err != nil {
			return err
		}
	}

	return nil
}
