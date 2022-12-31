package main

import (
	"context"
	"fmt"
	"go-test/ent"
	"go-test/ent/migrate"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		"db", "5432", "db", "root", "root"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(
		ctx,
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	log.Print("ent migrate done.")
}
