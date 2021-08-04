package main

import (
    "hr-face-free/routers"
    "fmt"
    "flag"
)

func main () {
   r := routers.SetupRouter()

   //dev: go run main.go -port=8081
   port := flag.Int("port", 8080,"端口号，默认为8080")
   flag.Parse()

   if err := r.Run(fmt.Sprintf(":%d", *port)); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
