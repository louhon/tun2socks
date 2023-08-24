package main

import (
	"C"
	"fmt"

	"go.uber.org/automaxprocs/maxprocs"

	_ "github.com/xjasonlyu/tun2socks/v2/dns"
	"github.com/xjasonlyu/tun2socks/v2/engine"
	"github.com/xjasonlyu/tun2socks/v2/internal/version"
)

var (
	key = new(engine.Key)

	configFile  string
	versionFlag bool
)

//export t2s_engine_start
func t2s_engine_start(device *C.char, proxy *C.char, iface *C.char) {
	maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))
	fmt.Println(version.String())
	fmt.Println(version.BuildString())
	var lk = new(engine.Key)
	lk.Mark = 0
	lk.MTU = 1400
	lk.UDPTimeout = 0
	lk.Device = C.GoString(device)
	lk.Proxy = C.GoString(proxy)
	lk.Interface = C.GoString(iface)
	lk.LogLevel = "info"
	engine.Insert(lk)
	engine.Start()
}

//export t2s_engine_stop
func t2s_engine_stop() {
	engine.Stop()
}

func main() {

}

