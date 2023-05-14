package main

import (
	_ "jk-goframe-shop/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"jk-goframe-shop/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "jk-goframe-shop/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
