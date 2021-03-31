package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalUint(u uint) graphql.Marshaler {
	return MarshalUint64(uint64(u))
}

func UnmarshalUint(v interface{}) (uint, error) {
	i, err := UnmarshalUint64(v)
	return uint(i), err
}

func MarshalUint64(u uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.FormatUint(u, 10))
		if err != nil {
			return
		}
	})
}

func UnmarshalUint64(v interface{}) (uint64, error) {
	switch t := v.(type) {
	case string:
		return strconv.ParseUint(t, 10, 64)
	case int:
		return uint64(t), nil
	case int64:
		return uint64(t), nil
	case json.Number:
		i, err := t.Int64()
		return uint64(i), err
	case float64:
		return uint64(t), nil
	default:
		return 0, fmt.Errorf("%v %T is not an uint64", v, v)
	}
}
