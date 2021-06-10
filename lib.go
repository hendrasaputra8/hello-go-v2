package main 

import (
	"context"
	"database/sql"
	"os"
	
	_ "github.com/go-sql-driver/mysql"
)

func con() (*sql.conn, error) {
	connString := os.gotenv("MYSQL_CONN_STRING")
	db, err := sql.Open("mysql", connString)
	if err!= nil {
		return nil, err
	}

	conn, err := db.conn(context.Baground())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func wrideData(w http.ResponseWriter, data interface{}) {
	w.header().Set("Content-type", "application/json")
	
	result := map[string]interface{}{
		"Status": http.StatusOK,
		"Data":	data,
		"Message": "",
	} 

	err := json.NewEncoder(w).Encoder(result)
	if err!= nil {
		log.Println("ERROR", err.error())
		http.error(w, err.error(), http.StatusInternalServerError)
	}
}

func wrideError(w http.ResponseWriter, err error{}) {
	log.Println("ERROR", err.error())
	
	result := map[string]interface{}{
		"Status": http.StatusInternalServerError,
		"Data":	nill,
		"Message": err.Error(),
	} 

	err := json.NewEncoder(w).Encoder(result)
	if err!= nil {
		log.Println("ERROR", err.error())
		http.error(w, err.error(), http.StatusInternalServerError)
	}
}