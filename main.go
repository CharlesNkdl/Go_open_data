package main

import (
	"fmt"
	"go_data_gouv_client/config"
	"go_data_gouv_client/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	fmt.Printf("salut la base de donnée : %v\n", db)
}
