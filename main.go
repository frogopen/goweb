package main

import (
	"fmt"
	"goweb/admin"
	"log"
	"net/http"
)

func main() {
	//配置初始化
	fmt.Println("service started.")
	httpsServer := admin.NewHTTPServer()
	log.Fatal(http.ListenAndServe(":8080", httpsServer.Router))
}
