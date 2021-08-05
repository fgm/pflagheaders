package pflagheaders

import (
	"errors"
	"net/http"
	"strings"

	"github.com/spf13/pflag"
)

// NameLong is the recommanded long name for a header flag.
const NameLong = "header"

// NameShort is the recommended short name for a header flag.
const NameShort = "H"

// Help is the recommended help string for header flags.
const Help = `add a HTTP request Header. May be used multiple times, like -H "accept: text/plain" -H "Authorization : Bearer cn389ncoiwuencr"`

// Type is the type expected by spf13/pflag for these data.
const Type = "stringSlice"

// ErrFormat is returned on ill-formatted header flag values.
var ErrFormat = errors.New(`header value is not a "name: value" string`)

// Header supports parsing HTTP headers as CLI arguments using spf13/pflag.
//
// If the same header is passed multiple times on the CLI, the values are
// aggregated under the header key, not replaced.
type Header struct {
	http.Header
}

// String implements fmt.Stringer and part of pflag.Value.
//
// It needs to have a value receiver to work on both header and *header.
func (h Header) String() string {
	if h.Header == nil {
		return ""
	}
	sb := strings.Builder{}
	h.Write(&sb)
	return sb.String()
}

// Set is part of pflag.Value.
func (h *Header) Set(s string) error {
	if h.Header == nil {
		h.Header = make(http.Header)
	}
	kvp := strings.Split(s, ":")
	if len(kvp) != 2 {
		return ErrFormat
	}
	for i := 0; i < len(kvp); i++ {
		kvp[i] = strings.Trim(kvp[i], " ")
	}
	h.Header.Add(kvp[0], kvp[1])
	return nil
}

// Type is part of pflag.Value.
func (h Header) Type() string {
	return Type
}

// HeaderFlag returns a global default header flag.
func HeaderFlag() *Header {
	h := &Header{}
	HeaderFlagP(h)
	return h
}

// HeaderFlagP initializes a global default header flag.
func HeaderFlagP(header *Header) {
	pflag.CommandLine.VarP(header, NameLong, NameShort, Help)
}
