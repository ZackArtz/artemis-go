package cmd

import (
	"github.com/zackartz/artemis-go/framework"
)

func PingCommand(ctx framework.Context) {
	ctx.Reply("pong!")
}
