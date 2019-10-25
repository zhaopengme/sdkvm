package util

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestCheckHostSpeed(t *testing.T) {
	urls := g.MapStrStr{"1": "https://nodejs.org/dist/", "2": "https://npm.taobao.org/mirrors/node/"}
	out := CheckHostSpeed(urls)
	t.Log(out)
}
