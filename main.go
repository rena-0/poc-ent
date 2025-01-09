package main

import (
    "context"
    "encoding/csv"
    "log"
    "os"
    "strconv"

    "github.com/rena-0/poc-ent/ent"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    Name string
    Age  int
}

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

    // Read CSV
    users, err := readCSV("data/users.csv")
    if err != nil {
        log.Fatalf("failed to read CSV: %v", err)
    }

    // Import CSV data into SQLite
    if err := importCSVToEnt(ctx, client, users); err != nil {
        log.Fatalf("failed to import CSV: %v", err)
    }

    log.Println("CSV data imported successfully!")
}

// Reads a CSV file and returns a slice of User structs
func readCSV(filename string) ([]User, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var users []User
    for _, record := range records {
        age, _ := strconv.Atoi(record[1]) // Convert age to int
        users = append(users, User{Name: record[0], Age: age})
    }
    return users, nil
}

// Imports CSV data into SQLite using Ent
func importCSVToEnt(ctx context.Context, client *ent.Client, users []User) error {
    for _, u := range users {
        _, err := client.User.
            Create().
            SetName(u.Name).
            SetAge(u.Age).
            Save(ctx)
        if err != nil {
            return err
        }
    }
    return nil
}
