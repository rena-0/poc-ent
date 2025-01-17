// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/rena-0/poc-ent/ent/airport"
	"github.com/rena-0/poc-ent/ent/schema"
	"github.com/rena-0/poc-ent/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	airportFields := schema.Airport{}.Fields()
	_ = airportFields
	// airportDescIdent is the schema descriptor for ident field.
	airportDescIdent := airportFields[1].Descriptor()
	// airport.IdentValidator is a validator for the "ident" field. It is called by the builders before save.
	airport.IdentValidator = airportDescIdent.Validators[0].(func(string) error)
	// airportDescType is the schema descriptor for type field.
	airportDescType := airportFields[2].Descriptor()
	// airport.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	airport.TypeValidator = airportDescType.Validators[0].(func(string) error)
	// airportDescName is the schema descriptor for name field.
	airportDescName := airportFields[3].Descriptor()
	// airport.NameValidator is a validator for the "name" field. It is called by the builders before save.
	airport.NameValidator = airportDescName.Validators[0].(func(string) error)
	// airportDescContinent is the schema descriptor for continent field.
	airportDescContinent := airportFields[7].Descriptor()
	// airport.ContinentValidator is a validator for the "continent" field. It is called by the builders before save.
	airport.ContinentValidator = airportDescContinent.Validators[0].(func(string) error)
	// airportDescIsoCountry is the schema descriptor for iso_country field.
	airportDescIsoCountry := airportFields[8].Descriptor()
	// airport.IsoCountryValidator is a validator for the "iso_country" field. It is called by the builders before save.
	airport.IsoCountryValidator = airportDescIsoCountry.Validators[0].(func(string) error)
	// airportDescIsoRegion is the schema descriptor for iso_region field.
	airportDescIsoRegion := airportFields[9].Descriptor()
	// airport.IsoRegionValidator is a validator for the "iso_region" field. It is called by the builders before save.
	airport.IsoRegionValidator = airportDescIsoRegion.Validators[0].(func(string) error)
	// airportDescMunicipality is the schema descriptor for municipality field.
	airportDescMunicipality := airportFields[10].Descriptor()
	// airport.MunicipalityValidator is a validator for the "municipality" field. It is called by the builders before save.
	airport.MunicipalityValidator = airportDescMunicipality.Validators[0].(func(string) error)
	// airportDescScheduledService is the schema descriptor for scheduled_service field.
	airportDescScheduledService := airportFields[11].Descriptor()
	// airport.ScheduledServiceValidator is a validator for the "scheduled_service" field. It is called by the builders before save.
	airport.ScheduledServiceValidator = airportDescScheduledService.Validators[0].(func(string) error)
	// airportDescGpsCode is the schema descriptor for gps_code field.
	airportDescGpsCode := airportFields[12].Descriptor()
	// airport.GpsCodeValidator is a validator for the "gps_code" field. It is called by the builders before save.
	airport.GpsCodeValidator = airportDescGpsCode.Validators[0].(func(string) error)
	// airportDescLocalCode is the schema descriptor for local_code field.
	airportDescLocalCode := airportFields[13].Descriptor()
	// airport.LocalCodeValidator is a validator for the "local_code" field. It is called by the builders before save.
	airport.LocalCodeValidator = airportDescLocalCode.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[1].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
}
