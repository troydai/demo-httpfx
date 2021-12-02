package database

import (
	"encoding/base64"

	"github.com/troydai/demo-httpfx/types"
)

func NewDataSource() types.DataRetriver {
	return &database{}
}

var _ types.DataRetriver = (*database)(nil)

type database struct{}

func (d database) Get(key []byte) string {
	ec := base64.StdEncoding
	return ec.EncodeToString(key)
}
