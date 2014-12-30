package base_encoder

import (
	"bytes"
	"math"
)

const (
	Base62     = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MimeBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	Base64Url  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"
)

func Base62Encode(value int64) string {
	return Encode(value, Base62)
}

// This function doesn't care about padding - check MimeBase64 docs
func MimeBase64Encode(value int64) string {
	return Encode(value, MimeBase64)
}

func Base64UrlEncode(value int64) string {
	return Encode(value, Base64Url)
}

func Base62Decode(token string) int64 {
	return Decode(token, Base62)
}

// This function doesn't care about padding - check MimeBase64 docs
func MimeBase64Decode(token string) int64 {
	return Decode(token, MimeBase64)
}

func Base64UrlDecode(token string) int64 {
	return Decode(token, Base64Url)
}

func Encode(value int64, alphabet string) string {
	if value == 0 {
		return string(alphabet[0])
	}

	chars := make([]byte, 0)

	length := int64(len(alphabet))

	for value > 0 {
		result := value / length
		remainder := value % length
		chars = append(chars, alphabet[remainder])
		value = result
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

func Decode(token string, alphabet string) int64 {
	value := int64(0)
	idx := 0.0
	chars := []byte(alphabet)

	charsLength := float64(len(chars))
	tokenLength := float64(len(token))

	for _, c := range []byte(token) {
		power := tokenLength - (idx + 1)
		index := int64(bytes.IndexByte(chars, c))
		value += index * int64(math.Pow(charsLength, power))
		idx++
	}

	return value
}
