package service

type Shorten interface {
	Transform(id int64) string
}

type Encoder interface {
	Encode(number int64) string
}

func GenToken() string {
	enc := NewHashId()
	token, _ := enc.Encode(NewSnowflakeIdGen().GetID())
	return token
}
