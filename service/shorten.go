package service

type Shorten interface {
	Transform(id int64) string
}

type Encoder interface {
	Encode(number int64) string
}

func GenToken() string {
	enc := NewHashId(SetSalt("5ABt4tlmAAueED9"))
	token, _ := enc.Encode(NewSnowflakeIdGen().GetID())
	return token
}
