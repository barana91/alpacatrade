package main


import (
	"fmt"
	"github.com/baav91/alpacatrade/conf"
	"log"
)

func main() {
	fmt.Println("Hello, World!")
	c, err := conf.ReadConf("./resource/dev.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}