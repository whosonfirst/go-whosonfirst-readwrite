# go-whosonfirst-readwrite

Go package for reading and writing Who's On First documents from a variety of sources.

## Important

Two things:

1. It's probably too soon, for you. If nothing else the documentation is incomplete
2. While it is tempting to imagine that this could be used as a general purpose abstraction layer for reading and writing a variety of documents this is probably premature and likely incorrect. This is really about WOF documents. Once we get _that_ working we'll think about how this code applies more generally.

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
