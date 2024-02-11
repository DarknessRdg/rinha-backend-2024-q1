package decoder

import (
	"encoding/json"
	"io"
)

type jsonDecoder[T any] struct{}

func NewJsonDecoder[T any]() IDecoder[T] {
	return &jsonDecoder[T]{}
}

func (j *jsonDecoder[T]) Decode(reader io.Reader) (T, error) {
	var value T

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&value)
	return value, err
}
