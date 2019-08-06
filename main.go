package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mesuutt/claps/config"
	"github.com/mesuutt/claps/db"
	"github.com/mesuutt/claps/migration"
	"github.com/mesuutt/claps/server"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Str("pkg", "claps").Logger()

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)
	db.Init()
	migration.Migrate()

	defer db.GetDB().Close()

	server.Init()
}
