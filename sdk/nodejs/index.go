package nodejs

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/gfile"
	"github.com/zhaopengme/sdkvm/mlog"
	"github.com/zhaopengme/sdkvm/sdk"
	"github.com/zhaopengme/sdkvm/util"
	"github.com/zhaopengme/sdkvm/util/gziputil"
	"path/filepath"
)

type NodeSdk struct {
	sdk.Sdk
	host string
}

func (this *NodeSdk) Init() {
	urls := g.MapStrStr{"host": "https://npm.taobao.org/mirrors/node", "chinaHost": "https://nodejs.org/dist"}
	host := util.CheckHostSpeed(urls)
	if host == "" {
		mlog.Fatal("timeout,please retry.")
	}
	this.host = host
}

func (this *NodeSdk) Install(version string) {
	this.Init()
	filename := fmt.Sprintf("node-%s-darwin-x64.tar.gz", version)
	url := fmt.Sprintf("%s/%s/%s", this.host, version, filename)
	bytes := ghttp.GetBytes(url)
	tmpFile := "/tmp/" + filename
	e := gfile.PutBytes(tmpFile, bytes)
	if e != nil {
		mlog.Fatal(e)
	}
	e = gziputil.Decompress(tmpFile, "tmp")
	if e != nil {
		mlog.Fatal(e)
	}
	tmpSdkDir := fmt.Sprintf("node-%s-darwin-x64", version)
	home, _ := gfile.Home()
	sdkHome := filepath.Join(home, ".sdkvm", "node", version)
	e = gfile.Mkdir(sdkHome)
	if e != nil {
		mlog.Fatal(e)
	}
	if gfile.Exists(sdkHome) {
		gfile.Remove(sdkHome)
	}
	e = gfile.Move(filepath.Join("/tmp/", tmpSdkDir), sdkHome)
	if e != nil {
		mlog.Fatal(e)
	}
}

func (this *NodeSdk) Versions() {
	this.Init()
	content := ghttp.GetContent(this.host + "/index.json")
	versions, e := gjson.LoadContent(content)
	if e != nil {
		mlog.Fatal(e)
	}
	for _, version := range versions.ToArray() {
		v := version.(map[string]interface{})
		mlog.Printf("%s", v["version"])
	}
}

func (this *NodeSdk) LocalVersions() {
	this.Init()
	home, _ := gfile.Home()
	sdkHome := filepath.Join(home, ".sdkvm", "node")
	dirs, e := gfile.ScanDir(sdkHome, "v*")
	if e != nil {
		mlog.Fatal(e)
	}
	for _, v := range dirs {
		mlog.Print(gfile.Basename(v))
	}
}

func (this *NodeSdk) UseVersion(version string) {
	this.Init()
	home, _ := gfile.Home()
	sdkHome := filepath.Join(home, ".sdkvm", "node", version)
	if !gfile.Exists(sdkHome) {
		mlog.Fatalf("%s not exists ", version)
	}
	gfile.Move(sdkHome, sdkHome+"_default")
}

func (this *NodeSdk) SetEnv() {
	this.Init()
	genv.Set("sdkvm","hello")
}