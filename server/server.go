package server

import (
	"github.com/mesuutt/claps/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()

	r.Run(config.GetString("server.port"))
}
