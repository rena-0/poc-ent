package main

import (
    "context"
    "encoding/csv"
    "log"
    "os"
    "strconv"

    "github.com/rena-0/poc-ent/ent"
    "github.com/rena-0/poc-ent/ent/airport"
    _ "github.com/mattn/go-sqlite3"
)

type AirportData struct {
    ID              int
    Ident          string
    Type           string
    Name           string
    LatitudeDeg    float64
    LongitudeDeg   float64
    ElevationFt    int
    Continent      string
    IsoCountry     string
    IsoRegion      string
    Municipality   string
    ScheduledService string
    GpsCode        string
    LocalCode      string
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
    airports, err := readCSV("data/airports.csv")
    if err != nil {
        log.Fatalf("failed to read CSV: %v", err)
    }

    // Import CSV data into SQLite
    if err := importCSVToEnt(ctx, client, airports); err != nil {
        log.Fatalf("failed to import CSV: %v", err)
    }

    log.Println("CSV data imported successfully!")
}

// Reads a CSV file and returns a slice of AirportData structs
func readCSV(filename string) ([]AirportData, error) {
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

    var airports []AirportData
    for _, record := range records[1:] { // Skip the header row
        id, _ := strconv.Atoi(record[0])         // ID to int
        latitude, _ := strconv.ParseFloat(record[4], 64) // Latitude to float
        longitude, _ := strconv.ParseFloat(record[5], 64) // Longitude to float
        elevation, _ := strconv.Atoi(record[6])   // Elevation to int

        airports = append(airports, AirportData{
            ID:              id,
            Ident:           record[1],
            Type:            record[2],
            Name:            record[3],
            LatitudeDeg:     latitude,
            LongitudeDeg:    longitude,
            ElevationFt:     elevation,
            Continent:       record[7],
            IsoCountry:      record[8],
            IsoRegion:       record[9],
            Municipality:    record[10],
            ScheduledService: record[11],
            GpsCode:         record[12],
            LocalCode:       record[13],
        })
    }
    return airports, nil
}

// Imports CSV data into SQLite using Ent
func importCSVToEnt(ctx context.Context, client *ent.Client, airports []AirportData) error {
    for _, a := range airports {
        _, err := client.Airport.
            Create().
            SetID(a.ID).
            SetIdent(a.Ident).
            SetType(a.Type).
            SetName(a.Name).
            SetLatitudeDeg(a.LatitudeDeg).
            SetLongitudeDeg(a.LongitudeDeg).
            SetElevationFt(a.ElevationFt).
            SetContinent(a.Continent).
            SetIsoCountry(a.IsoCountry).
            SetIsoRegion(a.IsoRegion).
            SetMunicipality(a.Municipality).
            SetScheduledService(a.ScheduledService).
            SetGpsCode(a.GpsCode).
            SetLocalCode(a.LocalCode).
            Save(ctx)
        if err != nil {
            return err
        }
    }
    return nil
}