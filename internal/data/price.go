package data

import (
	"fmt"
	"strconv"
)

type Price int64

func (r Price) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d тг", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
