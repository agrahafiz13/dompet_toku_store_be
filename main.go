package main

import (
	"log"

	"github.com/agrahafiz13/dompet_toku_store_be/config"
	"github.com/agrahafiz13/dompet_toku_store_be/database"
	"github.com/agrahafiz13/dompet_toku_store_be/routes"
)

func main() {
	cfg := config.Load()

	db := database.InitMySQL(cfg)
	rdb := database.InitRedis(cfg)
	firebaseApp := database.InitFirebase(cfg)

	router := routes.Setup(db, rdb, firebaseApp, cfg)

	log.Printf("Server running on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
