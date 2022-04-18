package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	cmd := rootCmd

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	cmd.SetArgs([]string{"encrypt", "--message", "Why did the chicken cross the road?"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	expectedString := "Jul qvq gur puvpxra pebff gur ebnq?"
	assert.Equal(t, expectedString, string(out))
}

func TestDecrypt(t *testing.T) {
	cmd := rootCmd

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	cmd.SetArgs([]string{"encrypt", "--message", "Jul qvq gur puvpxra pebff gur ebnq?"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	expectedString := "Why did the chicken cross the road?"
	assert.Equal(t, expectedString, string(out))
}
