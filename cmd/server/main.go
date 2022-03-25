package main

import (
	"fmt"

	"github.com/DanielMartin96/goddit/internal/db"
)

func Run() error {
	fmt.Println("Running our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	return nil
}

func main()  {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}