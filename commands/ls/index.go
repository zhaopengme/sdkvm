package ls

import (
	"path/filepath"

	"github.com/gogf/gf/os/gfile"
	"github.com/zhaopengme/sdkvm/mlog"
)

func Run() {
	sdkDir, e := gfile.ScanDir(filepath.Join(gfile.Pwd(), "sdk"), "*")
	if e != nil {
		mlog.Fatal(e)
	}
	mlog.Print(sdkDir)
}
