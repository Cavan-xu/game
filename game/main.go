package main

import (
	"gameserver/game/router"
	"gameserver/game/service"
	"path/filepath"
	"time"

	"github.com/Cavan-xu/van/core/conf"
	"github.com/Cavan-xu/van/core/log"
	"github.com/Cavan-xu/van/vnet"
)

var (
	ConfigPath   = "config"
	ServerConfig = filepath.Join(ConfigPath, "conf.json")
)

func main() {
	c := &vnet.Config{}
	if err := conf.LoadConfig(ServerConfig, c); err != nil {
		panic(err)
	}
	server, err := vnet.NewServer(c, vnet.WithLog(&log.Log{}))
	if err != nil {
		panic(err)
	}

	server.AddRouter(router.NewLoginRouter())
	server.AddRouter(router.NewRegisterRouter())
	go func() {
		if err = server.Server(); err != nil {
			panic(err)
		}
	}()
	go Ticker()

	select {}
}

func Ticker() {
	defer func() {
		if err := recover(); err != nil {
			go Ticker()
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var tick int32
	for {
		select {
		case <-ticker.C:
			service.CheckOffLine(tick)
			tick++
		}
	}
}
