package service

import (
	"github.com/speps/go-hashids"
)

type Shorten interface {
	Transform(id int64) string
}

type HashidShorten struct {
	h *hashids.HashID
}

func NewHashidShorten() *HashidShorten {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	h, _ := hashids.NewWithData(hd)
	return &HashidShorten{h: h}
}

func (shorten *HashidShorten) Transform(id int64) string {
	str, _ := shorten.h.EncodeInt64([]int64{id})
	return str
}

func GenToken() string {
	return NewHashidShorten().Transform(NewSnowflakeIdGen().GetID())
}
