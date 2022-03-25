package main

import "fmt"

func Run()  {
	fmt.Println("Running our application")
}

func main()  {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}