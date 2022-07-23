// See https://github.com/golang/go/issues/26366.
package lib

import (
	_ "github.com/moontrade/nogc/alloc/rpmalloc/lib/darwin_amd64"
	_ "github.com/moontrade/nogc/alloc/rpmalloc/lib/darwin_arm64"
	_ "github.com/moontrade/nogc/alloc/rpmalloc/lib/linux_amd64"
)
