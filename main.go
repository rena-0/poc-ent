package main

import (
    "context"
    "log"

    "github.com/rena-0/poc-ent/ent"
    "entgo.io/ent/dialect/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to SQLite
    client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer client.Close()

    // Run migration
    ctx := context.Background()
    if err := client.Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    log.Println("Database schema created successfully!")
}
