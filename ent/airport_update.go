// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rena-0/poc-ent/ent/airport"
	"github.com/rena-0/poc-ent/ent/predicate"
)

// AirportUpdate is the builder for updating Airport entities.
type AirportUpdate struct {
	config
	hooks    []Hook
	mutation *AirportMutation
}

// Where appends a list predicates to the AirportUpdate builder.
func (au *AirportUpdate) Where(ps ...predicate.Airport) *AirportUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetIdent sets the "ident" field.
func (au *AirportUpdate) SetIdent(s string) *AirportUpdate {
	au.mutation.SetIdent(s)
	return au
}

// SetNillableIdent sets the "ident" field if the given value is not nil.
func (au *AirportUpdate) SetNillableIdent(s *string) *AirportUpdate {
	if s != nil {
		au.SetIdent(*s)
	}
	return au
}

// SetType sets the "type" field.
func (au *AirportUpdate) SetType(s string) *AirportUpdate {
	au.mutation.SetType(s)
	return au
}

// SetNillableType sets the "type" field if the given value is not nil.
func (au *AirportUpdate) SetNillableType(s *string) *AirportUpdate {
	if s != nil {
		au.SetType(*s)
	}
	return au
}

// SetName sets the "name" field.
func (au *AirportUpdate) SetName(s string) *AirportUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AirportUpdate) SetNillableName(s *string) *AirportUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetLatitudeDeg sets the "latitude_deg" field.
func (au *AirportUpdate) SetLatitudeDeg(f float64) *AirportUpdate {
	au.mutation.ResetLatitudeDeg()
	au.mutation.SetLatitudeDeg(f)
	return au
}

// SetNillableLatitudeDeg sets the "latitude_deg" field if the given value is not nil.
func (au *AirportUpdate) SetNillableLatitudeDeg(f *float64) *AirportUpdate {
	if f != nil {
		au.SetLatitudeDeg(*f)
	}
	return au
}

// AddLatitudeDeg adds f to the "latitude_deg" field.
func (au *AirportUpdate) AddLatitudeDeg(f float64) *AirportUpdate {
	au.mutation.AddLatitudeDeg(f)
	return au
}

// SetLongitudeDeg sets the "longitude_deg" field.
func (au *AirportUpdate) SetLongitudeDeg(f float64) *AirportUpdate {
	au.mutation.ResetLongitudeDeg()
	au.mutation.SetLongitudeDeg(f)
	return au
}

// SetNillableLongitudeDeg sets the "longitude_deg" field if the given value is not nil.
func (au *AirportUpdate) SetNillableLongitudeDeg(f *float64) *AirportUpdate {
	if f != nil {
		au.SetLongitudeDeg(*f)
	}
	return au
}

// AddLongitudeDeg adds f to the "longitude_deg" field.
func (au *AirportUpdate) AddLongitudeDeg(f float64) *AirportUpdate {
	au.mutation.AddLongitudeDeg(f)
	return au
}

// SetElevationFt sets the "elevation_ft" field.
func (au *AirportUpdate) SetElevationFt(i int) *AirportUpdate {
	au.mutation.ResetElevationFt()
	au.mutation.SetElevationFt(i)
	return au
}

// SetNillableElevationFt sets the "elevation_ft" field if the given value is not nil.
func (au *AirportUpdate) SetNillableElevationFt(i *int) *AirportUpdate {
	if i != nil {
		au.SetElevationFt(*i)
	}
	return au
}

// AddElevationFt adds i to the "elevation_ft" field.
func (au *AirportUpdate) AddElevationFt(i int) *AirportUpdate {
	au.mutation.AddElevationFt(i)
	return au
}

// SetContinent sets the "continent" field.
func (au *AirportUpdate) SetContinent(s string) *AirportUpdate {
	au.mutation.SetContinent(s)
	return au
}

// SetNillableContinent sets the "continent" field if the given value is not nil.
func (au *AirportUpdate) SetNillableContinent(s *string) *AirportUpdate {
	if s != nil {
		au.SetContinent(*s)
	}
	return au
}

// SetIsoCountry sets the "iso_country" field.
func (au *AirportUpdate) SetIsoCountry(s string) *AirportUpdate {
	au.mutation.SetIsoCountry(s)
	return au
}

// SetNillableIsoCountry sets the "iso_country" field if the given value is not nil.
func (au *AirportUpdate) SetNillableIsoCountry(s *string) *AirportUpdate {
	if s != nil {
		au.SetIsoCountry(*s)
	}
	return au
}

// SetIsoRegion sets the "iso_region" field.
func (au *AirportUpdate) SetIsoRegion(s string) *AirportUpdate {
	au.mutation.SetIsoRegion(s)
	return au
}

// SetNillableIsoRegion sets the "iso_region" field if the given value is not nil.
func (au *AirportUpdate) SetNillableIsoRegion(s *string) *AirportUpdate {
	if s != nil {
		au.SetIsoRegion(*s)
	}
	return au
}

// SetMunicipality sets the "municipality" field.
func (au *AirportUpdate) SetMunicipality(s string) *AirportUpdate {
	au.mutation.SetMunicipality(s)
	return au
}

// SetNillableMunicipality sets the "municipality" field if the given value is not nil.
func (au *AirportUpdate) SetNillableMunicipality(s *string) *AirportUpdate {
	if s != nil {
		au.SetMunicipality(*s)
	}
	return au
}

// SetScheduledService sets the "scheduled_service" field.
func (au *AirportUpdate) SetScheduledService(s string) *AirportUpdate {
	au.mutation.SetScheduledService(s)
	return au
}

// SetNillableScheduledService sets the "scheduled_service" field if the given value is not nil.
func (au *AirportUpdate) SetNillableScheduledService(s *string) *AirportUpdate {
	if s != nil {
		au.SetScheduledService(*s)
	}
	return au
}

// SetGpsCode sets the "gps_code" field.
func (au *AirportUpdate) SetGpsCode(s string) *AirportUpdate {
	au.mutation.SetGpsCode(s)
	return au
}

// SetNillableGpsCode sets the "gps_code" field if the given value is not nil.
func (au *AirportUpdate) SetNillableGpsCode(s *string) *AirportUpdate {
	if s != nil {
		au.SetGpsCode(*s)
	}
	return au
}

// SetLocalCode sets the "local_code" field.
func (au *AirportUpdate) SetLocalCode(s string) *AirportUpdate {
	au.mutation.SetLocalCode(s)
	return au
}

// SetNillableLocalCode sets the "local_code" field if the given value is not nil.
func (au *AirportUpdate) SetNillableLocalCode(s *string) *AirportUpdate {
	if s != nil {
		au.SetLocalCode(*s)
	}
	return au
}

// Mutation returns the AirportMutation object of the builder.
func (au *AirportUpdate) Mutation() *AirportMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AirportUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AirportUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AirportUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AirportUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AirportUpdate) check() error {
	if v, ok := au.mutation.Ident(); ok {
		if err := airport.IdentValidator(v); err != nil {
			return &ValidationError{Name: "ident", err: fmt.Errorf(`ent: validator failed for field "Airport.ident": %w`, err)}
		}
	}
	if v, ok := au.mutation.GetType(); ok {
		if err := airport.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Airport.type": %w`, err)}
		}
	}
	if v, ok := au.mutation.Name(); ok {
		if err := airport.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Airport.name": %w`, err)}
		}
	}
	if v, ok := au.mutation.Continent(); ok {
		if err := airport.ContinentValidator(v); err != nil {
			return &ValidationError{Name: "continent", err: fmt.Errorf(`ent: validator failed for field "Airport.continent": %w`, err)}
		}
	}
	if v, ok := au.mutation.IsoCountry(); ok {
		if err := airport.IsoCountryValidator(v); err != nil {
			return &ValidationError{Name: "iso_country", err: fmt.Errorf(`ent: validator failed for field "Airport.iso_country": %w`, err)}
		}
	}
	if v, ok := au.mutation.IsoRegion(); ok {
		if err := airport.IsoRegionValidator(v); err != nil {
			return &ValidationError{Name: "iso_region", err: fmt.Errorf(`ent: validator failed for field "Airport.iso_region": %w`, err)}
		}
	}
	if v, ok := au.mutation.Municipality(); ok {
		if err := airport.MunicipalityValidator(v); err != nil {
			return &ValidationError{Name: "municipality", err: fmt.Errorf(`ent: validator failed for field "Airport.municipality": %w`, err)}
		}
	}
	if v, ok := au.mutation.ScheduledService(); ok {
		if err := airport.ScheduledServiceValidator(v); err != nil {
			return &ValidationError{Name: "scheduled_service", err: fmt.Errorf(`ent: validator failed for field "Airport.scheduled_service": %w`, err)}
		}
	}
	if v, ok := au.mutation.GpsCode(); ok {
		if err := airport.GpsCodeValidator(v); err != nil {
			return &ValidationError{Name: "gps_code", err: fmt.Errorf(`ent: validator failed for field "Airport.gps_code": %w`, err)}
		}
	}
	if v, ok := au.mutation.LocalCode(); ok {
		if err := airport.LocalCodeValidator(v); err != nil {
			return &ValidationError{Name: "local_code", err: fmt.Errorf(`ent: validator failed for field "Airport.local_code": %w`, err)}
		}
	}
	return nil
}

func (au *AirportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Ident(); ok {
		_spec.SetField(airport.FieldIdent, field.TypeString, value)
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.SetField(airport.FieldType, field.TypeString, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.LatitudeDeg(); ok {
		_spec.SetField(airport.FieldLatitudeDeg, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedLatitudeDeg(); ok {
		_spec.AddField(airport.FieldLatitudeDeg, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.LongitudeDeg(); ok {
		_spec.SetField(airport.FieldLongitudeDeg, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedLongitudeDeg(); ok {
		_spec.AddField(airport.FieldLongitudeDeg, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.ElevationFt(); ok {
		_spec.SetField(airport.FieldElevationFt, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedElevationFt(); ok {
		_spec.AddField(airport.FieldElevationFt, field.TypeInt, value)
	}
	if value, ok := au.mutation.Continent(); ok {
		_spec.SetField(airport.FieldContinent, field.TypeString, value)
	}
	if value, ok := au.mutation.IsoCountry(); ok {
		_spec.SetField(airport.FieldIsoCountry, field.TypeString, value)
	}
	if value, ok := au.mutation.IsoRegion(); ok {
		_spec.SetField(airport.FieldIsoRegion, field.TypeString, value)
	}
	if value, ok := au.mutation.Municipality(); ok {
		_spec.SetField(airport.FieldMunicipality, field.TypeString, value)
	}
	if value, ok := au.mutation.ScheduledService(); ok {
		_spec.SetField(airport.FieldScheduledService, field.TypeString, value)
	}
	if value, ok := au.mutation.GpsCode(); ok {
		_spec.SetField(airport.FieldGpsCode, field.TypeString, value)
	}
	if value, ok := au.mutation.LocalCode(); ok {
		_spec.SetField(airport.FieldLocalCode, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AirportUpdateOne is the builder for updating a single Airport entity.
type AirportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AirportMutation
}

// SetIdent sets the "ident" field.
func (auo *AirportUpdateOne) SetIdent(s string) *AirportUpdateOne {
	auo.mutation.SetIdent(s)
	return auo
}

// SetNillableIdent sets the "ident" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableIdent(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetIdent(*s)
	}
	return auo
}

// SetType sets the "type" field.
func (auo *AirportUpdateOne) SetType(s string) *AirportUpdateOne {
	auo.mutation.SetType(s)
	return auo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableType(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetType(*s)
	}
	return auo
}

// SetName sets the "name" field.
func (auo *AirportUpdateOne) SetName(s string) *AirportUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableName(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetLatitudeDeg sets the "latitude_deg" field.
func (auo *AirportUpdateOne) SetLatitudeDeg(f float64) *AirportUpdateOne {
	auo.mutation.ResetLatitudeDeg()
	auo.mutation.SetLatitudeDeg(f)
	return auo
}

// SetNillableLatitudeDeg sets the "latitude_deg" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableLatitudeDeg(f *float64) *AirportUpdateOne {
	if f != nil {
		auo.SetLatitudeDeg(*f)
	}
	return auo
}

// AddLatitudeDeg adds f to the "latitude_deg" field.
func (auo *AirportUpdateOne) AddLatitudeDeg(f float64) *AirportUpdateOne {
	auo.mutation.AddLatitudeDeg(f)
	return auo
}

// SetLongitudeDeg sets the "longitude_deg" field.
func (auo *AirportUpdateOne) SetLongitudeDeg(f float64) *AirportUpdateOne {
	auo.mutation.ResetLongitudeDeg()
	auo.mutation.SetLongitudeDeg(f)
	return auo
}

// SetNillableLongitudeDeg sets the "longitude_deg" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableLongitudeDeg(f *float64) *AirportUpdateOne {
	if f != nil {
		auo.SetLongitudeDeg(*f)
	}
	return auo
}

// AddLongitudeDeg adds f to the "longitude_deg" field.
func (auo *AirportUpdateOne) AddLongitudeDeg(f float64) *AirportUpdateOne {
	auo.mutation.AddLongitudeDeg(f)
	return auo
}

// SetElevationFt sets the "elevation_ft" field.
func (auo *AirportUpdateOne) SetElevationFt(i int) *AirportUpdateOne {
	auo.mutation.ResetElevationFt()
	auo.mutation.SetElevationFt(i)
	return auo
}

// SetNillableElevationFt sets the "elevation_ft" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableElevationFt(i *int) *AirportUpdateOne {
	if i != nil {
		auo.SetElevationFt(*i)
	}
	return auo
}

// AddElevationFt adds i to the "elevation_ft" field.
func (auo *AirportUpdateOne) AddElevationFt(i int) *AirportUpdateOne {
	auo.mutation.AddElevationFt(i)
	return auo
}

// SetContinent sets the "continent" field.
func (auo *AirportUpdateOne) SetContinent(s string) *AirportUpdateOne {
	auo.mutation.SetContinent(s)
	return auo
}

// SetNillableContinent sets the "continent" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableContinent(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetContinent(*s)
	}
	return auo
}

// SetIsoCountry sets the "iso_country" field.
func (auo *AirportUpdateOne) SetIsoCountry(s string) *AirportUpdateOne {
	auo.mutation.SetIsoCountry(s)
	return auo
}

// SetNillableIsoCountry sets the "iso_country" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableIsoCountry(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetIsoCountry(*s)
	}
	return auo
}

// SetIsoRegion sets the "iso_region" field.
func (auo *AirportUpdateOne) SetIsoRegion(s string) *AirportUpdateOne {
	auo.mutation.SetIsoRegion(s)
	return auo
}

// SetNillableIsoRegion sets the "iso_region" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableIsoRegion(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetIsoRegion(*s)
	}
	return auo
}

// SetMunicipality sets the "municipality" field.
func (auo *AirportUpdateOne) SetMunicipality(s string) *AirportUpdateOne {
	auo.mutation.SetMunicipality(s)
	return auo
}

// SetNillableMunicipality sets the "municipality" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableMunicipality(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetMunicipality(*s)
	}
	return auo
}

// SetScheduledService sets the "scheduled_service" field.
func (auo *AirportUpdateOne) SetScheduledService(s string) *AirportUpdateOne {
	auo.mutation.SetScheduledService(s)
	return auo
}

// SetNillableScheduledService sets the "scheduled_service" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableScheduledService(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetScheduledService(*s)
	}
	return auo
}

// SetGpsCode sets the "gps_code" field.
func (auo *AirportUpdateOne) SetGpsCode(s string) *AirportUpdateOne {
	auo.mutation.SetGpsCode(s)
	return auo
}

// SetNillableGpsCode sets the "gps_code" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableGpsCode(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetGpsCode(*s)
	}
	return auo
}

// SetLocalCode sets the "local_code" field.
func (auo *AirportUpdateOne) SetLocalCode(s string) *AirportUpdateOne {
	auo.mutation.SetLocalCode(s)
	return auo
}

// SetNillableLocalCode sets the "local_code" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableLocalCode(s *string) *AirportUpdateOne {
	if s != nil {
		auo.SetLocalCode(*s)
	}
	return auo
}

// Mutation returns the AirportMutation object of the builder.
func (auo *AirportUpdateOne) Mutation() *AirportMutation {
	return auo.mutation
}

// Where appends a list predicates to the AirportUpdate builder.
func (auo *AirportUpdateOne) Where(ps ...predicate.Airport) *AirportUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AirportUpdateOne) Select(field string, fields ...string) *AirportUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Airport entity.
func (auo *AirportUpdateOne) Save(ctx context.Context) (*Airport, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AirportUpdateOne) SaveX(ctx context.Context) *Airport {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AirportUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AirportUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AirportUpdateOne) check() error {
	if v, ok := auo.mutation.Ident(); ok {
		if err := airport.IdentValidator(v); err != nil {
			return &ValidationError{Name: "ident", err: fmt.Errorf(`ent: validator failed for field "Airport.ident": %w`, err)}
		}
	}
	if v, ok := auo.mutation.GetType(); ok {
		if err := airport.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Airport.type": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Name(); ok {
		if err := airport.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Airport.name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Continent(); ok {
		if err := airport.ContinentValidator(v); err != nil {
			return &ValidationError{Name: "continent", err: fmt.Errorf(`ent: validator failed for field "Airport.continent": %w`, err)}
		}
	}
	if v, ok := auo.mutation.IsoCountry(); ok {
		if err := airport.IsoCountryValidator(v); err != nil {
			return &ValidationError{Name: "iso_country", err: fmt.Errorf(`ent: validator failed for field "Airport.iso_country": %w`, err)}
		}
	}
	if v, ok := auo.mutation.IsoRegion(); ok {
		if err := airport.IsoRegionValidator(v); err != nil {
			return &ValidationError{Name: "iso_region", err: fmt.Errorf(`ent: validator failed for field "Airport.iso_region": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Municipality(); ok {
		if err := airport.MunicipalityValidator(v); err != nil {
			return &ValidationError{Name: "municipality", err: fmt.Errorf(`ent: validator failed for field "Airport.municipality": %w`, err)}
		}
	}
	if v, ok := auo.mutation.ScheduledService(); ok {
		if err := airport.ScheduledServiceValidator(v); err != nil {
			return &ValidationError{Name: "scheduled_service", err: fmt.Errorf(`ent: validator failed for field "Airport.scheduled_service": %w`, err)}
		}
	}
	if v, ok := auo.mutation.GpsCode(); ok {
		if err := airport.GpsCodeValidator(v); err != nil {
			return &ValidationError{Name: "gps_code", err: fmt.Errorf(`ent: validator failed for field "Airport.gps_code": %w`, err)}
		}
	}
	if v, ok := auo.mutation.LocalCode(); ok {
		if err := airport.LocalCodeValidator(v); err != nil {
			return &ValidationError{Name: "local_code", err: fmt.Errorf(`ent: validator failed for field "Airport.local_code": %w`, err)}
		}
	}
	return nil
}

func (auo *AirportUpdateOne) sqlSave(ctx context.Context) (_node *Airport, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Airport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, airport.FieldID)
		for _, f := range fields {
			if !airport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != airport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Ident(); ok {
		_spec.SetField(airport.FieldIdent, field.TypeString, value)
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.SetField(airport.FieldType, field.TypeString, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.LatitudeDeg(); ok {
		_spec.SetField(airport.FieldLatitudeDeg, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedLatitudeDeg(); ok {
		_spec.AddField(airport.FieldLatitudeDeg, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.LongitudeDeg(); ok {
		_spec.SetField(airport.FieldLongitudeDeg, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedLongitudeDeg(); ok {
		_spec.AddField(airport.FieldLongitudeDeg, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.ElevationFt(); ok {
		_spec.SetField(airport.FieldElevationFt, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedElevationFt(); ok {
		_spec.AddField(airport.FieldElevationFt, field.TypeInt, value)
	}
	if value, ok := auo.mutation.Continent(); ok {
		_spec.SetField(airport.FieldContinent, field.TypeString, value)
	}
	if value, ok := auo.mutation.IsoCountry(); ok {
		_spec.SetField(airport.FieldIsoCountry, field.TypeString, value)
	}
	if value, ok := auo.mutation.IsoRegion(); ok {
		_spec.SetField(airport.FieldIsoRegion, field.TypeString, value)
	}
	if value, ok := auo.mutation.Municipality(); ok {
		_spec.SetField(airport.FieldMunicipality, field.TypeString, value)
	}
	if value, ok := auo.mutation.ScheduledService(); ok {
		_spec.SetField(airport.FieldScheduledService, field.TypeString, value)
	}
	if value, ok := auo.mutation.GpsCode(); ok {
		_spec.SetField(airport.FieldGpsCode, field.TypeString, value)
	}
	if value, ok := auo.mutation.LocalCode(); ok {
		_spec.SetField(airport.FieldLocalCode, field.TypeString, value)
	}
	_node = &Airport{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
