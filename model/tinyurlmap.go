package model

import (
	"time"
)

type TinyUrlMap struct {
	ID          int64     `ddb:"id"`
	OriginUrl   string    `ddb:"origin_url"`
	ShortUrl    string    `ddb:"short_url"`
	CreatedTime time.Time `ddb:"created_time"`
}
