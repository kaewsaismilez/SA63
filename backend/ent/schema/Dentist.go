package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// Dentist holds the schema definition for the Dentist entity.
type Dentist struct {
   ent.Schema
}
 
// Fields of the Dentist.
func (Dentist) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
 }
 
// Edges of the Dentist.
func (Dentist) Edges() []ent.Edge {
   return []ent.Edge{
	edge.To("medicalfiles", Medicalfile.Type).StorageKey(edge.Column("dentist_id")),
	}	
}
