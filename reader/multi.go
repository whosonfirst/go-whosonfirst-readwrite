package reader

import (
	"errors"
)

type MultiReader struct {
	Reader
	readers []Reader
}

func NewMultiReader(readers ...Reader) (Reader, error) {

	mr := MultiReader{
		reader: readers,
	}

	return &mr, nil
}

func (mr *MultiReader) Read(uri string) (io.ReadCloser, error) {

	for _, r := range mr.readers {

		fh, err := r.Read(uri)

		if err != nil {
			continue
		}

		return fh, nil
	}

	return nil, errors.New("Unable to read URI")
}

func (r *MultiReader) URI(uri string) string {
	return uri
}
