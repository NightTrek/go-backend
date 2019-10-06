package sqlcontroller

import (
	"database/sql"
)



func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}


}