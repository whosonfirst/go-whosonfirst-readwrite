package bytes

import (
	gobytes "bytes"
	"io/ioutil"
)

func ReadCloserFromBytes(b []byte) (io.ReadCloser, error) {
	body := gobytes.NewReader(b)
	return ioutil.NopCloser(body)
}
