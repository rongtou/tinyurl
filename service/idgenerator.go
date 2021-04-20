package service

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

type IdGenerator interface {
	GetID() (int64, error)
}

type SnowflakeIdGenerator struct {
	node *snowflake.Node
}

func NewSnowflakeIdGen() *SnowflakeIdGenerator {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalln(err)
	}
	return &SnowflakeIdGenerator{
		node: node,
	}
}

func (idGen *SnowflakeIdGenerator) GetID() int64 {
	id := idGen.node.Generate().Int64()
	return id
}
