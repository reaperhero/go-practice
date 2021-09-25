package gohashid

import (
	hashids "github.com/speps/go-hashids"
)

const (
	salt = "jm3nxsitckyob89lup6eqg4wfra1d7"
)

func HashIDEncodeBySalt(number int64) (result string, err error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 8
	hd.Alphabet = "KtR7qal9FnTSDJIorwhPLjHAvcBp1G"

	id, _ := hashids.NewWithData(hd)

	result, err = id.EncodeInt64([]int64{number})
	return
}

func HashIDDecodeBySalt(hash string) (number int64, err error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 8
	hd.Alphabet = "KtR7qal9FnTSDJIorwhPLjHAvcBp1G"

	id, _ := hashids.NewWithData(hd)

	result, err := id.DecodeInt64WithError(hash)
	if err == nil {
		if len(result) != 0 {
			number = result[0]
		}
	}
	return
}
