package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// Airport holds the schema definition for the Airport entity.
type Airport struct {
    ent.Schema
}

// Fields of the Airport.
func (Airport) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique(),              
        field.String("ident").NotEmpty(),       
        field.String("type").NotEmpty(),        
        field.String("name").NotEmpty(),        
        field.Float("latitude_deg"),            
        field.Float("longitude_deg"),           
        field.Int("elevation_ft"),              
        field.String("continent").NotEmpty(),   
        field.String("iso_country").NotEmpty(), 
        field.String("iso_region").NotEmpty(),  
        field.String("municipality").NotEmpty(), 
        field.String("scheduled_service").NotEmpty(),
        field.String("gps_code").NotEmpty(),   
        field.String("local_code").NotEmpty(), 
    }
}
