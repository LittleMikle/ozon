package main

import (
	"github.com/LittleMikle/ozon/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal()
	}

}
