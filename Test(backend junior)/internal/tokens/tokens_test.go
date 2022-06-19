package tokens

import (
	"testing"
)

type Case struct {
	tk   *Tokens
	fail bool
}

type Cases struct {
	c []Case
}

func TestSecretKey(t *testing.T) {
	testCases := Cases{
		c: []Case{
			{
				tk: &Tokens{
					SecretKey: []byte("test"),
				},
				fail: false,
			},
			{
				tk: &Tokens{
					SecretKey: []byte(" "),
				},
				fail: false},
			{
				tk: &Tokens{
					SecretKey: []byte(""),
				},
				fail: false},
		},
	}

	for _, c := range testCases.c {
		_, _, err := c.tk.generateTokens("id")
		if err != nil && !c.fail {
			t.Errorf("unexpected error: %#v", err)
		}
	}
}
