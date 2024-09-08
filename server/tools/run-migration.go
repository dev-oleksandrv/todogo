package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/dev-oleksandrv/config"
)

func main() {
	config.LoadConfig()
	dbUrl := os.Getenv("DATABASE_URL")
	cmd := exec.Command("migrate", "-path", "db/migrations", "-database", dbUrl, "up")
	cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
        log.Fatalf("Migration failed: %v", err)
    }
    fmt.Println("Migration completed successfully.")
}