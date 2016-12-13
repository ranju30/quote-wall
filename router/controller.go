package router

import (
	"github.com/gorilla/mux"
	"net/http"
	quoteHandler "quote_wall/handler"
)

func HandleRequest() {
	handler := mux.NewRouter();

	handler.HandleFunc("/store", quoteHandler.SaveQuote()).Methods("POST")
	handler.HandleFunc("/store", quoteHandler.GetQuote()).Methods("GET")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")));

	http.Handle("/", handler);
}
