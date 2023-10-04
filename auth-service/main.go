package main

import (
	"fmt"
	"library-management/common/db"
)

func main() {

	_, err := db.InitializeDB()
	if err != nil {
		fmt.Printf("failed to connect database an error occured :", err)
	}
	defer func() {
		if sqldb, err := db.DB.DB(); err == nil {
			sqldb.Close()
		}
	}()
}
