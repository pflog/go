package main

import (
	"time"

	"github.com/mrcrgl/pflog/log"

	"github.com/mrcrgl/pflog/container"
)

func main() {
	for {
		log.With(container.NewCtx("4b560c51-358b-4908-b87c-79c0613c457a", "request")).Info("Yeah")

		log.Info(`This is
a multiline message.

Yeah.`)
		log.Info("Hallo")
		log.Infof("Hallo %s! Es ist %d Stunde nach 12", "Heinz", 12)

		time.Sleep(time.Second)
	}
}
