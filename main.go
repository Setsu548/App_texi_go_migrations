package main

import (
	"github.com/Setsu548/App_texi_go_migrations/config"
	"github.com/Setsu548/App_texi_go_migrations/db"
	"github.com/Setsu548/App_texi_go_migrations/migrate"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	dbConn, err := db.PostgresDBConnect(cfg)
	if err != nil {
		panic(err)
	}

	// Inicializar bootstrap con TokenManager singleton y migraciones
	mg := migrate.NewMigration(dbConn, cfg)
	mg.Migrate()
}
