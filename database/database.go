package database

import (
	"database/sql"
	"quote_wall/model"
	_"github.com/lib/pq"
)

func StoreQuote(db *sql.DB, quotes *model.Quote) (error) {
	_, queryErr := db.Exec("INSERT INTO Quote_Wall (quote,name) VALUES($1,$2)", quotes.Quote, quotes.Name)
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}

func GetAllQuote(db *sql.DB) ([]*model.Quote, error) {
	rows, err := db.Query("SELECT id,name,quote from Quote_Wall")
	if (err != nil) {
		return [] *model.Quote{}, err;
	}
	var quotes []*model.Quote
	for rows.Next() {
		var name string
		var quote string
		var id int
		rows.Scan(&id, &name, &quote)
		quotes = append(quotes, &model.Quote{Name:name, Quote:quote})
	}
	return quotes, nil
}
