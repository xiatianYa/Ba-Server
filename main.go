package main

import (
	_ "Ba-Server/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"Ba-Server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
