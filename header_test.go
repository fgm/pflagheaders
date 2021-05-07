package pflagheaders_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/spf13/pflag"

	"github.com/fgm/pflagheaders"
)

func TestHappy(t *testing.T) {
	h := &pflagheaders.Header{}
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	fs.VarP(h, pflagheaders.NameLong, pflagheaders.NameShort, pflagheaders.Help)
	fs.Parse([]string{
		"-H",
		"content-type: text/plain",
		"-H",
		"Authorization: bearer foo",
	})
	expected := "Authorization: bearer foo\r\nContent-Type: text/plain\r\n"
	actual := h.String()
	if actual != expected {
		t.Errorf("Expected (len %d):\n%s\n%s\nGot (len %d):\n%s\n%s",
			len(expected), expected, hex.Dump([]byte(expected)),
			len(actual), actual, hex.Dump([]byte(actual)))
	}
}

func TestSad_BadFormat(t *testing.T) {
	h := &pflagheaders.Header{}
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	fs.VarP(h, pflagheaders.NameLong, pflagheaders.NameShort, pflagheaders.Help)
	err := fs.Parse([]string{
		"-H",
		"content-type text/plain",
	})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	// pflag does not (yet ?) support errors.Is().
	// See https://github.com/spf13/pflag/issues/227
	actual := err.Error()
	expectedSub := pflagheaders.ErrFormat.Error()
	if !strings.Contains(actual, expectedSub) {
		t.Fatalf("Expected string containing:\n%20s\ngot:\n%20s", expectedSub, actual)
	}
}
