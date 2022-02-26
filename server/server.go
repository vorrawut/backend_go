package server

import "safebsc/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.host") + ":" + config.GetString("server.port"))
}
