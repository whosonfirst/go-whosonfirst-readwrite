package reader

import (
	"github.com/whosonfirst/go-whosonfirst-readwrite/cache"
	"io"
)

type CacheReader struct {
	Reader
	reader Reader
	cache  cache.Cache
}

func NewCacheReader(r Reader, c cache.Cache) (Reader, error) {

	cr := CacheReader{
		reader: r,
		cache:  c,
	}

	return &cr, nil
}

func (r *CacheReader) Read(key string) (io.ReadCloser, error) {

	fh, err := r.cache.Get(key)

	if err == nil {
		return fh, nil
	}

	if err != nil && err.Error() != "CACHE MISS" {
		return nil, err
	}

	fh, err = r.reader.Read(key)

	if err != nil {
		return nil, err
	}

	fh, err = r.cache.Set(key, fh)

	if err != nil {
		return nil, err
	}

	return fh, nil
}
