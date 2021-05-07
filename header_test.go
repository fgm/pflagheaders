package pflagheaders_test

import (
	"encoding/hex"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/pflag"

	pfh "github.com/fgm/pflagheaders"
)

func TestHappy(t *testing.T) {
	h := &pfh.Header{}
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	fs.VarP(h, pfh.NameLong, pfh.NameShort, pfh.Help)
	fs.Parse([]string{
		"-H", "content-type: text/plain",
		"-H", "Authorization: bearer foo",
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
	h := &pfh.Header{}
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	fs.VarP(h, pfh.NameLong, pfh.NameShort, pfh.Help)
	err := fs.Parse([]string{
		"-H", "content-type text/plain",
	})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	// pflag does not (yet ?) support errors.Is().
	// See https://github.com/spf13/pflag/issues/227
	actual := err.Error()
	expectedSub := pfh.ErrFormat.Error()
	if !strings.Contains(actual, expectedSub) {
		t.Fatalf("Expected string containing:\n%20s\ngot:\n%20s", expectedSub, actual)
	}
}

func TestHeaderFlag(t *testing.T) {
	cl := pflag.CommandLine
	defer func() { pflag.CommandLine = cl }()

	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	pflag.CommandLine = fs
	f := pfh.HeaderFlag()
	if f == nil {
		t.Fatalf("Expected non-nil Header")
	}
	err := fs.Parse([]string{"-H", "key: value"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// Keys are normalized, so
	expected := http.Header{"Key": []string{"value"}}
	actual := f.Header
	if !cmp.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
