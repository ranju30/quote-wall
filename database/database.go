package database

import (
	"database/sql"
	"quote_wall/model"
	_"github.com/lib/pq"
)

func StoreQuote(db *sql.DB, quotes *model.Quote) (error) {
	_, queryErr := db.Exec("INSERT INTO Quote_Wall (quote,name) VALUES($1,$2)", quotes.Quote,quotes.Name)
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}
