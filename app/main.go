package main

import (
	"net/http"
	"quote_wall/router"
	"fmt"
)


func main() {
	port := "8080";
	fmt.Println("Server is running at port : ",port);
	router.HandleRequest();
	http.ListenAndServe(":" + port, nil)
}
