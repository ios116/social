package main

import (
	flag "github.com/spf13/pflag"
	"log"
	"os/user"
	"social/cmd/data"
	"social/cmd/rpc"
	"social/cmd/web"
	"syscall"
)

func main() {
	maxOpenFiles()
	var service = flag.StringP("service", "s", "rpc", "kind of service")
	flag.Parse()
	switch *service {
	case "rpc":
		rpc.Server()
	case "web":
		web.Server()
	case "data":
		data.DataForTest()
	}
}

func maxOpenFiles() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	log.Println("Hi " + user.Name + " (id: " + user.Uid + ")")
	log.Println("Username: " + user.Username)
	log.Println("Home Dir: " + user.HomeDir)

	var rLimit syscall.Rlimit
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Println("Error Getting Rlimit ", err)
	}
	log.Println("rLimit=> ", rLimit)

	if rLimit.Cur < rLimit.Max {
		rLimit.Cur = rLimit.Max
		err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err != nil {
			log.Println("Error Setting Rlimit ", err)
		}
	}
	log.Println("new limit =>", rLimit.Cur)
}
