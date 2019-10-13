package main

import (
	flag "github.com/spf13/pflag"
	"social/cmd/rpc"
	"social/cmd/web"
)

func main() {
	var service = flag.StringP("service", "s", "rpc", "kind of service")
	flag.Parse()
	switch *service {
	case "rpc":
		rpc.Server()
	case "web":
		web.Server()
	}
}
