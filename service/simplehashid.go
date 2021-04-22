package service

import "errors"

const (
	// DefaultAlphabet is the default Alphabet used by go-hashids
	DefaultAlphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

type HashIdOption struct {
	Alphabet string
	Salt     string
}

type ApplyHashIdOptionFunc func(*HashIdOption)

type HashId struct {
	Option   HashIdOption
	alphabet []rune
	salt     []rune
}

func SetAlphabet(alphabet string) ApplyHashIdOptionFunc {
	return func(o *HashIdOption) {
		o.Alphabet = alphabet
	}
}

func SetSalt(salt string) ApplyHashIdOptionFunc {
	return func(o *HashIdOption) {
		o.Salt = salt
	}
}

func NewHashId(opts ...ApplyHashIdOptionFunc) *HashId {
	opt := HashIdOption{
		Alphabet: DefaultAlphabet,
	}

	for _, o := range opts {
		o(&opt)
	}

	alphabet := []rune(opt.Alphabet)
	salt := []rune(opt.Salt)

	consistentShuffleInPlace(alphabet, salt)

	return &HashId{
		Option:   opt,
		alphabet: alphabet,
		salt:     salt,
	}
}

func (h *HashId) Encode(number int64) (string, error) {
	if number < 0 {
		return "", errors.New("negative number not supported")
	}

	input := number
	var result []rune
	alen := len(h.alphabet)
	for {
		r := h.alphabet[input%int64(alen)]
		result = append(result, r)
		input /= int64(alen)
		if input == 0 {
			break
		}
	}
	return string(result), nil
}

func consistentShuffleInPlace(alphabet []rune, salt []rune) {
	if len(salt) == 0 {
		return
	}

	for i, v, p := len(alphabet)-1, 0, 0; i > 0; i-- {
		p += int(salt[v])
		j := (int(salt[v]) + v + p) % i
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
		v = (v + 1) % len(salt)
	}
	return
}
