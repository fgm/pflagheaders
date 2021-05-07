# HTTP Headers for spf13/pflag

This package allows use of repeated CLI flags defining HTTP headers, like:

## CLI usage

```bash
# Long format
mycommand --header "Accept: text/plain" --header "Authorization: bearer sometoken"

# Short format
mycommand -H "Accept: text/plain" -H "Authorization: bearer sometoken"

# Repeated headers are supported and combined
mycommand -H "X-Array-Header: value1" -H "X-Array-Header: value2"
# Will return a slice value with value1 and value2 for key  X-Array-Header

# Headers are canonicalized
mycommand -H "content-type: application/json"
# Will have key Content-Type
```

## Code usage

The simplest

```go
package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/fgm/pflagheaders"
)

func main() {
	// HeaderFlag provides a preconfigured default flag
	h := pflagheaders.HeaderFlag()
	pflag.Parse()

	fmt.Printf("Headers:\n%s\n", h)
	// The resulting http.Header is available after Parse:
	fmt.Printf("Inner header:\n%#v\n", h.Header)
}
```
