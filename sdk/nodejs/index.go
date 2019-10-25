package nodejs

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/zhaopengme/sdkvm/mlog"
	"github.com/zhaopengme/sdkvm/sdk"
	"github.com/zhaopengme/sdkvm/util"
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

func (this *NodeSdk) Versions() {
	this.Init()
	content := ghttp.GetContent(this.host + "/index.json")
	versions, e := gjson.LoadContent(content)
	if e != nil {
		mlog.Fatal(e)
	}
	for _,version := range versions.ToArray(){
		v:=version.(map[string]interface{})
		mlog.Printf("%s",v["version"])
	}
}
