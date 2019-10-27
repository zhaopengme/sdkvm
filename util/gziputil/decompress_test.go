package gziputil

import "testing"

func TestDecompress(t *testing.T) {
	Decompress("/tmp/node-v9.9.0-darwin-x64.tar.gz", "/tmp/")
}
