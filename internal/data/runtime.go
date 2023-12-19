package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// value receiver is used in the method, bcs we don't want to edit original Runtime
func (r *Runtime) MarshalJSON() ([]byte, error) {
	// you can put int32 to Sprintf to change it to string
	jsonValue := fmt.Sprintf("%d mins", r)
	// needs to be quoted, instead  102 mins it should be "102 mins"
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// this is a hack - the method changes the value of *r pointer, so it points to new Runtime value unmarshalled 
	*r = Runtime(i)
	return nil
}
