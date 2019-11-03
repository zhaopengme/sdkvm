package gziputil

import (
	"github.com/gogf/gf/os/gfile"
	"testing"
)

func TestDecompress(t *testing.T) {
	e :=Decompress("/var/folders/d6/ttzx_qyd6y35bb9nx9nrdmvc0000gp/T/node-v13.0.1-darwin-x64.tar.gz", gfile.TempDir())
	print(e)
}
