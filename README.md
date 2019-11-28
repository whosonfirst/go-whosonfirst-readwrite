# go-whosonfirst-readwrite

Go package for reading and writing Who's On First documents from a variety of sources.

## Important

This package is officially deprecated and has been superseded by the [go-reader-*](https://github.com/whosonfirst?utf8=%E2%9C%93&q=go-reader&type=&language=) and [go-writer-*](https://github.com/whosonfirst?utf8=%E2%9C%93&q=go-writer&type=&language=) packages, as well as [go-whosonfirst-reader](https://github.com/whosonfirst/go-whosonfirst-reader) and [go-whosonfirst-writer](https://github.com/whosonfirst/go-whosonfirst-writer). In time this repository will be archived.

## Install

You will need to have both `Go` (specifically version [1.12](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make tools
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Interfaces

## Readers

```
type Reader interface {
	Read(string) (io.ReadCloser, error)
}
```

## Writers

```
type Writer interface {
	Write(string, io.ReadCloser) error
}
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-readwrite-http
* https://github.com/whosonfirst/go-whosonfirst-readwrite-mysql
* https://github.com/whosonfirst/go-whosonfirst-readwrite-github
* https://github.com/whosonfirst/go-whosonfirst-readwrite-sqlite
* https://github.com/whosonfirst/go-whosonfirst-readwrite-s3
* https://github.com/whosonfirst/go-whosonfirst-cache
* https://github.com/whosonfirst/go-whosonfirst-cache-bigcache
* https://github.com/whosonfirst/go-whosonfirst-cache-s3
