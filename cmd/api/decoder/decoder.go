package decoder

import "io"

type IDecoder[T any] interface {
	Decode(io.Reader) (T, error)
}
