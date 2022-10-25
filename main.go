package main

import (
	"fmt"
	_bpjsRefrepo "github.com/dedinirtadinata/bpjs-helper/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config3m.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	r := gin.Default()
	_bpjsRefrepo.NewSIMRSHandler(r)
	r.Run(viper.GetString("server.address"))

}
