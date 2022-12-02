package main

import (
	"gameserver/game/router"
	"path/filepath"

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
	server.Server()
}
