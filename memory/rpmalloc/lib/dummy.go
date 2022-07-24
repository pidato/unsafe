// See https://github.com/golang/go/issues/26366.
package lib

import (
	_ "github.com/pidato/unsafe/memory/rpmalloc/lib/darwin_amd64"
	_ "github.com/pidato/unsafe/memory/rpmalloc/lib/darwin_arm64"
	_ "github.com/pidato/unsafe/memory/rpmalloc/lib/linux_amd64"
)
