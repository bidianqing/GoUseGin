package idgen

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node
var err error

func NewIdGenerator(generatorId int64) {
	node, err = snowflake.NewNode(generatorId)
	if err != nil {
		panic(err)
	}
}

func Generate() int64 {
	// Generate a snowflake ID.
	id := node.Generate()

	return id.Int64()
}
