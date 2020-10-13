package main

import (
	"e-commerce/config/db"
	"fmt"
)

func main() {
	err := db.Migrate()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("oke")
}

// for migrate db --> open comments
// func main ()  {
// 	db.Migrate()
// }
