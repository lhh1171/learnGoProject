package main

import (
	"flag"
	"fmt"
	"micro/kit/pkg/configs"
	"micro/kit/pkg/databases"
)

var confFile = flag.String("f", "library-user-service/user.yaml", "user config file")

func main() {
	flag.Parse()

	err := configs.Init(*confFile)
	if err != nil {
		panic(err)
	}

	err = databases.InitMySql(configs.Conf.MySQLConfig)
	if err != nil {
		fmt.Println("load mysql failed")
	}
	_, ok := databases.DB.Get("library")
	fmt.Println(ok)
}
