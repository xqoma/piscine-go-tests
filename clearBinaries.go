package piscinegotests

import (
	"os"
	"path"
)

func ClearBinaries() {
	binariesDir := path.Join(os.TempDir(), "binaries")
	os.RemoveAll(binariesDir)
}
