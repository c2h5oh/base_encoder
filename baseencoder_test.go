package baseencoder

import (
	"testing"
)

func Test_Base62Encode(t *testing.T) {
	testValues := map[string]int64{
		"0":    0,
		"1B":   99,
		"1MmC": 424242,
	}

	for token, value := range testValues {
		encoded := Base62Encode(value)
		if encoded != token {
			t.Logf("Base62 encoding of %d failed. Expected %s, got %s", value, token, encoded)
			t.Fail()
		}
	}
}

func Test_MimeBase64Encode(t *testing.T) {
	testValues := map[string]int64{
		"A":    0,
		"Bj":   99,
		"Bnky": 424242,
	}

	for token, value := range testValues {
		encoded := MimeBase64Encode(value)
		if encoded != token {
			t.Logf("MimeBase64 encoding of %d failed. Expected %s, got %s", value, token, encoded)
			t.Fail()
		}
	}
}

func Test_Base64UrlEncode(t *testing.T) {
	testValues := map[string]int64{
		"A":    0,
		"Bj":   99,
		"Bnky": 424242,
	}

	for token, value := range testValues {
		encoded := Base64UrlEncode(value)
		if encoded != token {
			t.Logf("Base64Url encoding of %d failed. Expected %s, got %s", value, token, encoded)
			t.Fail()
		}
	}
}

func Test_Base62Decode(t *testing.T) {
	testValues := map[string]int64{
		"0":    0,
		"1B":   99,
		"1MmC": 424242,
	}

	for token, value := range testValues {
		decoded := Base62Decode(token)
		if decoded != value {
			t.Logf("Base62 decoding of %s failed. Expected %d, got %d", token, value, decoded)
			t.Fail()
		}
	}
}

func Test_MimeBase64Decode(t *testing.T) {
	testValues := map[string]int64{
		"A":    0,
		"Bj":   99,
		"Bnky": 424242,
	}

	for token, value := range testValues {
		decoded := MimeBase64Decode(token)
		if decoded != value {
			t.Logf("MimeBase64 decoding of %s failed. Expected %d, got %d", token, value, decoded)
			t.Fail()
		}
	}
}

func Test_Base64UrlDecode(t *testing.T) {
	testValues := map[string]int64{
		"A":    0,
		"Bj":   99,
		"Bnky": 424242,
	}

	for token, value := range testValues {
		decoded := Base64UrlDecode(token)
		if decoded != value {
			t.Logf("Base64Url decoding of %s failed. Expected %d, got %d", token, value, decoded)
			t.Fail()
		}
	}
}
