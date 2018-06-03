package pruner

import (
	"github.com/whosonfirst/go-whosonfirst-readwrite/bytes"
	"io"
)

type NullPruner struct {
	Pruner
}

func NewNullPruner() (Pruner, error) {

	r := NullPruner{}
	return &r, nil
}

func (r *NullPruner) Prune(uri string) error {
     return nil
}
