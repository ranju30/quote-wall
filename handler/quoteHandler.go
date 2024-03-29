package handler

import (
	"strings"
	"quote_wall/model"
	"quote_wall/database"
	"log"
	"net/http"
	"database/sql"
	"fmt"
)

func openDB() *sql.DB {
	db, err := database.OpenDatabase()
	if (err != nil) {
		log.Fatal(err.Error())
	}
	db.Ping();
	return db;
}

func SaveQuote() http.HandlerFunc {
	db := openDB();
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm();
		name := strings.Join(req.Form["name"], "");
		quote := strings.Join(req.Form["quote"], "");
		quote_to_db := model.Quote{Quote:quote, Name:name}
		err := database.StoreQuote(db, &quote_to_db)
		if (err != nil) {
			log.Fatal(err.Error())
			res.Write([]byte("got error.."))
			return
		}
		res.Write([]byte("Quote has stored.."));
	}
}

func GetQuote() http.HandlerFunc {
	db := openDB();
	return func(res http.ResponseWriter, req *http.Request) {
		data, err := database.GetAllQuote(db);
		if (err != nil) {
			res.Write([]byte("got error.."))
			return
		}
		var quotes string;
		for _, each := range data {
			fmt.Println(each)
			quotes += "<b>" + each.Name + "</b>" + " : " + "<span>" + each.Quote + "</span>" + "<br/>"
		}
		res.Write([]byte(quotes));
	}
}
