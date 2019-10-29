package main

import (
	flag "github.com/spf13/pflag"
	"social/cmd/data"
	"social/cmd/web"
)

func main() {
	var service = flag.StringP("service", "s", "rpc", "kind of service")
	flag.Parse()
	switch *service {
	case "web":
		web.Server()
	case "data":
		data.DataForTest()
	}
}

