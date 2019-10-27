package util

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"math"
	"sync"
	"time"
)

const (
	gPROXY_CHECK_TIMEOUT = 5 * time.Second
)

func CheckHostSpeed(urls map[string]string) string {
	wg := sync.WaitGroup{}
	checkMap := gmap.NewIntStrMap(true)
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			checkMap.Set(checkProxyLatency(url), url)
		}()
	}
	garray.NewStrArray()
	wg.Wait()
	url := ""
	latency := math.MaxInt32
	checkMap.Iterator(func(k int, v string) bool {
		if k < latency {
			url = v
			latency = k
		}
		return true
	})
	return url
}

// checkProxyLatency checks the latency for specified url.
func checkProxyLatency(url string) int {
	httpClient := ghttp.NewClient()
	httpClient.SetTimeOut(gPROXY_CHECK_TIMEOUT)

	start := gtime.Millisecond()
	r, err := httpClient.Head(url)
	if err != nil || r.StatusCode != 200 {
		return math.MaxInt32
	}
	defer r.Close()
	return int(gtime.Millisecond() - start)
}
