package pruner

import (

)

type Pruner interface {
	Pruner(string) error
}
