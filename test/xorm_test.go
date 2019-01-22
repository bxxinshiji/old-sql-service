package test

import (
	"log"
	"os"
	"testing"

	_ "gitee.com/xsjcs/old-sql-service/providers/database"
)

func TestXorm(t *testing.T) {
	if os.Getenv("DB_PASSWORD") == "" {
		log.Println("12")
	}
	log.Println(os.Getenv("DB_PASSWORD"))
}
