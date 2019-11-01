package nodejs

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gview"
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
	tmpFile := gfile.Join(gfile.TempDir(), filename)
	e := gfile.PutBytes(tmpFile, bytes)
	if e != nil {
		mlog.Fatal(e)
	}
	e = gziputil.Decompress(tmpFile, gfile.TempDir())
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
	e = gfile.Move(filepath.Join(gfile.TempDir(), tmpSdkDir), sdkHome)
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

func (this *NodeSdk) GenerateEnv(version string) {
	home, _ := gfile.Home()
	sdkHome := filepath.Join(home, ".sdkvm", "node", version)
	sdkShell := filepath.Join(home, ".sdkvm", "env.sh")

	content, e := gview.ParseContent(`#!/bin/sh
export PATH="{{.sdkHome}}/bin:$PATH"
	`, g.Map{"sdkHome": sdkHome})
	if e != nil {
		mlog.Fatal(e)
	}
	e = gfile.PutContents(sdkShell, content)
	if e != nil {
		mlog.Fatal(e)
	}
}

func (this *NodeSdk) RunCmd() {
	command := gcmd.GetArg(2)
	switch command {
	case "ls-remote":
		this.Versions()
	case "ls":
		this.LocalVersions()
	case "use":
		version := gcmd.GetArg(3)
		//this.UseVersion(version)
		this.GenerateEnv(version)
	case "install":
		version := gcmd.GetArg(3)
		this.Install(version)
	case "uninstall":
		version := gcmd.GetArg(3)
		this.Uninstall(version)
	default:
		fmt.Println("not support")
	}
}
