package main

import (
	"fmt"
	"strings"

	"github.com/Shiirookami/weather-app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		return
	}

	db, err := buildGormDB(cfg.Postgres)
	if err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}
	defer db.Close()
	splash()

}

func buildGormDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func splash() {
	colorReset := "\033[0m"
	SplashText := `
	__      __               __  .__                      _____                 
   /  \    /  \ ____ _____ _/  |_|  |__   ___________    /  _  \ ______ ______  
   \   \/\/   // __ \\__  \\   __\  |  \_/ __ \_  __ \  /  /_\  \\____ \\____ \ 
	\        /\  ___/ / __ \|  | |   Y  \  ___/|  | \/ /    |    \  |_> >  |_> >
	 \__/\  /  \___  >____  /__| |___|  /\___  >__|    \____|__  /   __/|   __/ 
		  \/       \/     \/          \/     \/                \/|__|   |__|    
   `
	fmt.Println(colorReset, strings.TrimSpace(SplashText))
}
