package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

//value receiver is used in the method, bcs we don't want to edit original Runtime
func (r Runtime) MarshalJSON() ([]byte, error) {
	//you can put int32 to Sprintf to change it to string
	jsonValue := fmt.Sprintf("%d mins", r)
	//needs to be quoted, instead  102 mins it should be "102 mins"
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}