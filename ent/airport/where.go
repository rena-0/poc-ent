// Code generated by ent, DO NOT EDIT.

package airport

import (
	"entgo.io/ent/dialect/sql"
	"github.com/rena-0/poc-ent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldID, id))
}

// Ident applies equality check predicate on the "ident" field. It's identical to IdentEQ.
func Ident(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIdent, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldType, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldName, v))
}

// LatitudeDeg applies equality check predicate on the "latitude_deg" field. It's identical to LatitudeDegEQ.
func LatitudeDeg(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLatitudeDeg, v))
}

// LongitudeDeg applies equality check predicate on the "longitude_deg" field. It's identical to LongitudeDegEQ.
func LongitudeDeg(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLongitudeDeg, v))
}

// ElevationFt applies equality check predicate on the "elevation_ft" field. It's identical to ElevationFtEQ.
func ElevationFt(v int) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldElevationFt, v))
}

// Continent applies equality check predicate on the "continent" field. It's identical to ContinentEQ.
func Continent(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldContinent, v))
}

// IsoCountry applies equality check predicate on the "iso_country" field. It's identical to IsoCountryEQ.
func IsoCountry(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIsoCountry, v))
}

// IsoRegion applies equality check predicate on the "iso_region" field. It's identical to IsoRegionEQ.
func IsoRegion(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIsoRegion, v))
}

// Municipality applies equality check predicate on the "municipality" field. It's identical to MunicipalityEQ.
func Municipality(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldMunicipality, v))
}

// ScheduledService applies equality check predicate on the "scheduled_service" field. It's identical to ScheduledServiceEQ.
func ScheduledService(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldScheduledService, v))
}

// GpsCode applies equality check predicate on the "gps_code" field. It's identical to GpsCodeEQ.
func GpsCode(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldGpsCode, v))
}

// LocalCode applies equality check predicate on the "local_code" field. It's identical to LocalCodeEQ.
func LocalCode(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLocalCode, v))
}

// IdentEQ applies the EQ predicate on the "ident" field.
func IdentEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIdent, v))
}

// IdentNEQ applies the NEQ predicate on the "ident" field.
func IdentNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldIdent, v))
}

// IdentIn applies the In predicate on the "ident" field.
func IdentIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldIdent, vs...))
}

// IdentNotIn applies the NotIn predicate on the "ident" field.
func IdentNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldIdent, vs...))
}

// IdentGT applies the GT predicate on the "ident" field.
func IdentGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldIdent, v))
}

// IdentGTE applies the GTE predicate on the "ident" field.
func IdentGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldIdent, v))
}

// IdentLT applies the LT predicate on the "ident" field.
func IdentLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldIdent, v))
}

// IdentLTE applies the LTE predicate on the "ident" field.
func IdentLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldIdent, v))
}

// IdentContains applies the Contains predicate on the "ident" field.
func IdentContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldIdent, v))
}

// IdentHasPrefix applies the HasPrefix predicate on the "ident" field.
func IdentHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldIdent, v))
}

// IdentHasSuffix applies the HasSuffix predicate on the "ident" field.
func IdentHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldIdent, v))
}

// IdentEqualFold applies the EqualFold predicate on the "ident" field.
func IdentEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldIdent, v))
}

// IdentContainsFold applies the ContainsFold predicate on the "ident" field.
func IdentContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldIdent, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldType, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldName, v))
}

// LatitudeDegEQ applies the EQ predicate on the "latitude_deg" field.
func LatitudeDegEQ(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLatitudeDeg, v))
}

// LatitudeDegNEQ applies the NEQ predicate on the "latitude_deg" field.
func LatitudeDegNEQ(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldLatitudeDeg, v))
}

// LatitudeDegIn applies the In predicate on the "latitude_deg" field.
func LatitudeDegIn(vs ...float64) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldLatitudeDeg, vs...))
}

// LatitudeDegNotIn applies the NotIn predicate on the "latitude_deg" field.
func LatitudeDegNotIn(vs ...float64) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldLatitudeDeg, vs...))
}

// LatitudeDegGT applies the GT predicate on the "latitude_deg" field.
func LatitudeDegGT(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldLatitudeDeg, v))
}

// LatitudeDegGTE applies the GTE predicate on the "latitude_deg" field.
func LatitudeDegGTE(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldLatitudeDeg, v))
}

// LatitudeDegLT applies the LT predicate on the "latitude_deg" field.
func LatitudeDegLT(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldLatitudeDeg, v))
}

// LatitudeDegLTE applies the LTE predicate on the "latitude_deg" field.
func LatitudeDegLTE(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldLatitudeDeg, v))
}

// LongitudeDegEQ applies the EQ predicate on the "longitude_deg" field.
func LongitudeDegEQ(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLongitudeDeg, v))
}

// LongitudeDegNEQ applies the NEQ predicate on the "longitude_deg" field.
func LongitudeDegNEQ(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldLongitudeDeg, v))
}

// LongitudeDegIn applies the In predicate on the "longitude_deg" field.
func LongitudeDegIn(vs ...float64) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldLongitudeDeg, vs...))
}

// LongitudeDegNotIn applies the NotIn predicate on the "longitude_deg" field.
func LongitudeDegNotIn(vs ...float64) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldLongitudeDeg, vs...))
}

// LongitudeDegGT applies the GT predicate on the "longitude_deg" field.
func LongitudeDegGT(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldLongitudeDeg, v))
}

// LongitudeDegGTE applies the GTE predicate on the "longitude_deg" field.
func LongitudeDegGTE(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldLongitudeDeg, v))
}

// LongitudeDegLT applies the LT predicate on the "longitude_deg" field.
func LongitudeDegLT(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldLongitudeDeg, v))
}

// LongitudeDegLTE applies the LTE predicate on the "longitude_deg" field.
func LongitudeDegLTE(v float64) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldLongitudeDeg, v))
}

// ElevationFtEQ applies the EQ predicate on the "elevation_ft" field.
func ElevationFtEQ(v int) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldElevationFt, v))
}

// ElevationFtNEQ applies the NEQ predicate on the "elevation_ft" field.
func ElevationFtNEQ(v int) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldElevationFt, v))
}

// ElevationFtIn applies the In predicate on the "elevation_ft" field.
func ElevationFtIn(vs ...int) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldElevationFt, vs...))
}

// ElevationFtNotIn applies the NotIn predicate on the "elevation_ft" field.
func ElevationFtNotIn(vs ...int) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldElevationFt, vs...))
}

// ElevationFtGT applies the GT predicate on the "elevation_ft" field.
func ElevationFtGT(v int) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldElevationFt, v))
}

// ElevationFtGTE applies the GTE predicate on the "elevation_ft" field.
func ElevationFtGTE(v int) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldElevationFt, v))
}

// ElevationFtLT applies the LT predicate on the "elevation_ft" field.
func ElevationFtLT(v int) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldElevationFt, v))
}

// ElevationFtLTE applies the LTE predicate on the "elevation_ft" field.
func ElevationFtLTE(v int) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldElevationFt, v))
}

// ContinentEQ applies the EQ predicate on the "continent" field.
func ContinentEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldContinent, v))
}

// ContinentNEQ applies the NEQ predicate on the "continent" field.
func ContinentNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldContinent, v))
}

// ContinentIn applies the In predicate on the "continent" field.
func ContinentIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldContinent, vs...))
}

// ContinentNotIn applies the NotIn predicate on the "continent" field.
func ContinentNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldContinent, vs...))
}

// ContinentGT applies the GT predicate on the "continent" field.
func ContinentGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldContinent, v))
}

// ContinentGTE applies the GTE predicate on the "continent" field.
func ContinentGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldContinent, v))
}

// ContinentLT applies the LT predicate on the "continent" field.
func ContinentLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldContinent, v))
}

// ContinentLTE applies the LTE predicate on the "continent" field.
func ContinentLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldContinent, v))
}

// ContinentContains applies the Contains predicate on the "continent" field.
func ContinentContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldContinent, v))
}

// ContinentHasPrefix applies the HasPrefix predicate on the "continent" field.
func ContinentHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldContinent, v))
}

// ContinentHasSuffix applies the HasSuffix predicate on the "continent" field.
func ContinentHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldContinent, v))
}

// ContinentEqualFold applies the EqualFold predicate on the "continent" field.
func ContinentEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldContinent, v))
}

// ContinentContainsFold applies the ContainsFold predicate on the "continent" field.
func ContinentContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldContinent, v))
}

// IsoCountryEQ applies the EQ predicate on the "iso_country" field.
func IsoCountryEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIsoCountry, v))
}

// IsoCountryNEQ applies the NEQ predicate on the "iso_country" field.
func IsoCountryNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldIsoCountry, v))
}

// IsoCountryIn applies the In predicate on the "iso_country" field.
func IsoCountryIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldIsoCountry, vs...))
}

// IsoCountryNotIn applies the NotIn predicate on the "iso_country" field.
func IsoCountryNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldIsoCountry, vs...))
}

// IsoCountryGT applies the GT predicate on the "iso_country" field.
func IsoCountryGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldIsoCountry, v))
}

// IsoCountryGTE applies the GTE predicate on the "iso_country" field.
func IsoCountryGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldIsoCountry, v))
}

// IsoCountryLT applies the LT predicate on the "iso_country" field.
func IsoCountryLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldIsoCountry, v))
}

// IsoCountryLTE applies the LTE predicate on the "iso_country" field.
func IsoCountryLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldIsoCountry, v))
}

// IsoCountryContains applies the Contains predicate on the "iso_country" field.
func IsoCountryContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldIsoCountry, v))
}

// IsoCountryHasPrefix applies the HasPrefix predicate on the "iso_country" field.
func IsoCountryHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldIsoCountry, v))
}

// IsoCountryHasSuffix applies the HasSuffix predicate on the "iso_country" field.
func IsoCountryHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldIsoCountry, v))
}

// IsoCountryEqualFold applies the EqualFold predicate on the "iso_country" field.
func IsoCountryEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldIsoCountry, v))
}

// IsoCountryContainsFold applies the ContainsFold predicate on the "iso_country" field.
func IsoCountryContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldIsoCountry, v))
}

// IsoRegionEQ applies the EQ predicate on the "iso_region" field.
func IsoRegionEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldIsoRegion, v))
}

// IsoRegionNEQ applies the NEQ predicate on the "iso_region" field.
func IsoRegionNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldIsoRegion, v))
}

// IsoRegionIn applies the In predicate on the "iso_region" field.
func IsoRegionIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldIsoRegion, vs...))
}

// IsoRegionNotIn applies the NotIn predicate on the "iso_region" field.
func IsoRegionNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldIsoRegion, vs...))
}

// IsoRegionGT applies the GT predicate on the "iso_region" field.
func IsoRegionGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldIsoRegion, v))
}

// IsoRegionGTE applies the GTE predicate on the "iso_region" field.
func IsoRegionGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldIsoRegion, v))
}

// IsoRegionLT applies the LT predicate on the "iso_region" field.
func IsoRegionLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldIsoRegion, v))
}

// IsoRegionLTE applies the LTE predicate on the "iso_region" field.
func IsoRegionLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldIsoRegion, v))
}

// IsoRegionContains applies the Contains predicate on the "iso_region" field.
func IsoRegionContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldIsoRegion, v))
}

// IsoRegionHasPrefix applies the HasPrefix predicate on the "iso_region" field.
func IsoRegionHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldIsoRegion, v))
}

// IsoRegionHasSuffix applies the HasSuffix predicate on the "iso_region" field.
func IsoRegionHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldIsoRegion, v))
}

// IsoRegionEqualFold applies the EqualFold predicate on the "iso_region" field.
func IsoRegionEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldIsoRegion, v))
}

// IsoRegionContainsFold applies the ContainsFold predicate on the "iso_region" field.
func IsoRegionContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldIsoRegion, v))
}

// MunicipalityEQ applies the EQ predicate on the "municipality" field.
func MunicipalityEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldMunicipality, v))
}

// MunicipalityNEQ applies the NEQ predicate on the "municipality" field.
func MunicipalityNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldMunicipality, v))
}

// MunicipalityIn applies the In predicate on the "municipality" field.
func MunicipalityIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldMunicipality, vs...))
}

// MunicipalityNotIn applies the NotIn predicate on the "municipality" field.
func MunicipalityNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldMunicipality, vs...))
}

// MunicipalityGT applies the GT predicate on the "municipality" field.
func MunicipalityGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldMunicipality, v))
}

// MunicipalityGTE applies the GTE predicate on the "municipality" field.
func MunicipalityGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldMunicipality, v))
}

// MunicipalityLT applies the LT predicate on the "municipality" field.
func MunicipalityLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldMunicipality, v))
}

// MunicipalityLTE applies the LTE predicate on the "municipality" field.
func MunicipalityLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldMunicipality, v))
}

// MunicipalityContains applies the Contains predicate on the "municipality" field.
func MunicipalityContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldMunicipality, v))
}

// MunicipalityHasPrefix applies the HasPrefix predicate on the "municipality" field.
func MunicipalityHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldMunicipality, v))
}

// MunicipalityHasSuffix applies the HasSuffix predicate on the "municipality" field.
func MunicipalityHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldMunicipality, v))
}

// MunicipalityEqualFold applies the EqualFold predicate on the "municipality" field.
func MunicipalityEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldMunicipality, v))
}

// MunicipalityContainsFold applies the ContainsFold predicate on the "municipality" field.
func MunicipalityContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldMunicipality, v))
}

// ScheduledServiceEQ applies the EQ predicate on the "scheduled_service" field.
func ScheduledServiceEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldScheduledService, v))
}

// ScheduledServiceNEQ applies the NEQ predicate on the "scheduled_service" field.
func ScheduledServiceNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldScheduledService, v))
}

// ScheduledServiceIn applies the In predicate on the "scheduled_service" field.
func ScheduledServiceIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldScheduledService, vs...))
}

// ScheduledServiceNotIn applies the NotIn predicate on the "scheduled_service" field.
func ScheduledServiceNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldScheduledService, vs...))
}

// ScheduledServiceGT applies the GT predicate on the "scheduled_service" field.
func ScheduledServiceGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldScheduledService, v))
}

// ScheduledServiceGTE applies the GTE predicate on the "scheduled_service" field.
func ScheduledServiceGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldScheduledService, v))
}

// ScheduledServiceLT applies the LT predicate on the "scheduled_service" field.
func ScheduledServiceLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldScheduledService, v))
}

// ScheduledServiceLTE applies the LTE predicate on the "scheduled_service" field.
func ScheduledServiceLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldScheduledService, v))
}

// ScheduledServiceContains applies the Contains predicate on the "scheduled_service" field.
func ScheduledServiceContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldScheduledService, v))
}

// ScheduledServiceHasPrefix applies the HasPrefix predicate on the "scheduled_service" field.
func ScheduledServiceHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldScheduledService, v))
}

// ScheduledServiceHasSuffix applies the HasSuffix predicate on the "scheduled_service" field.
func ScheduledServiceHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldScheduledService, v))
}

// ScheduledServiceEqualFold applies the EqualFold predicate on the "scheduled_service" field.
func ScheduledServiceEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldScheduledService, v))
}

// ScheduledServiceContainsFold applies the ContainsFold predicate on the "scheduled_service" field.
func ScheduledServiceContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldScheduledService, v))
}

// GpsCodeEQ applies the EQ predicate on the "gps_code" field.
func GpsCodeEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldGpsCode, v))
}

// GpsCodeNEQ applies the NEQ predicate on the "gps_code" field.
func GpsCodeNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldGpsCode, v))
}

// GpsCodeIn applies the In predicate on the "gps_code" field.
func GpsCodeIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldGpsCode, vs...))
}

// GpsCodeNotIn applies the NotIn predicate on the "gps_code" field.
func GpsCodeNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldGpsCode, vs...))
}

// GpsCodeGT applies the GT predicate on the "gps_code" field.
func GpsCodeGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldGpsCode, v))
}

// GpsCodeGTE applies the GTE predicate on the "gps_code" field.
func GpsCodeGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldGpsCode, v))
}

// GpsCodeLT applies the LT predicate on the "gps_code" field.
func GpsCodeLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldGpsCode, v))
}

// GpsCodeLTE applies the LTE predicate on the "gps_code" field.
func GpsCodeLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldGpsCode, v))
}

// GpsCodeContains applies the Contains predicate on the "gps_code" field.
func GpsCodeContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldGpsCode, v))
}

// GpsCodeHasPrefix applies the HasPrefix predicate on the "gps_code" field.
func GpsCodeHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldGpsCode, v))
}

// GpsCodeHasSuffix applies the HasSuffix predicate on the "gps_code" field.
func GpsCodeHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldGpsCode, v))
}

// GpsCodeEqualFold applies the EqualFold predicate on the "gps_code" field.
func GpsCodeEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldGpsCode, v))
}

// GpsCodeContainsFold applies the ContainsFold predicate on the "gps_code" field.
func GpsCodeContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldGpsCode, v))
}

// LocalCodeEQ applies the EQ predicate on the "local_code" field.
func LocalCodeEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEQ(FieldLocalCode, v))
}

// LocalCodeNEQ applies the NEQ predicate on the "local_code" field.
func LocalCodeNEQ(v string) predicate.Airport {
	return predicate.Airport(sql.FieldNEQ(FieldLocalCode, v))
}

// LocalCodeIn applies the In predicate on the "local_code" field.
func LocalCodeIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldIn(FieldLocalCode, vs...))
}

// LocalCodeNotIn applies the NotIn predicate on the "local_code" field.
func LocalCodeNotIn(vs ...string) predicate.Airport {
	return predicate.Airport(sql.FieldNotIn(FieldLocalCode, vs...))
}

// LocalCodeGT applies the GT predicate on the "local_code" field.
func LocalCodeGT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGT(FieldLocalCode, v))
}

// LocalCodeGTE applies the GTE predicate on the "local_code" field.
func LocalCodeGTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldGTE(FieldLocalCode, v))
}

// LocalCodeLT applies the LT predicate on the "local_code" field.
func LocalCodeLT(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLT(FieldLocalCode, v))
}

// LocalCodeLTE applies the LTE predicate on the "local_code" field.
func LocalCodeLTE(v string) predicate.Airport {
	return predicate.Airport(sql.FieldLTE(FieldLocalCode, v))
}

// LocalCodeContains applies the Contains predicate on the "local_code" field.
func LocalCodeContains(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContains(FieldLocalCode, v))
}

// LocalCodeHasPrefix applies the HasPrefix predicate on the "local_code" field.
func LocalCodeHasPrefix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasPrefix(FieldLocalCode, v))
}

// LocalCodeHasSuffix applies the HasSuffix predicate on the "local_code" field.
func LocalCodeHasSuffix(v string) predicate.Airport {
	return predicate.Airport(sql.FieldHasSuffix(FieldLocalCode, v))
}

// LocalCodeEqualFold applies the EqualFold predicate on the "local_code" field.
func LocalCodeEqualFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldEqualFold(FieldLocalCode, v))
}

// LocalCodeContainsFold applies the ContainsFold predicate on the "local_code" field.
func LocalCodeContainsFold(v string) predicate.Airport {
	return predicate.Airport(sql.FieldContainsFold(FieldLocalCode, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Airport) predicate.Airport {
	return predicate.Airport(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Airport) predicate.Airport {
	return predicate.Airport(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Airport) predicate.Airport {
	return predicate.Airport(sql.NotPredicates(p))
}