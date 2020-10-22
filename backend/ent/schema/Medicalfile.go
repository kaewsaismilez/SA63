package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// Medicalfile holds the schema definition for the Medicalfile entity.
type Medicalfile struct {
   ent.Schema
}
 
// Fields of the Medicalfile.
func (Medicalfile) Fields() []ent.Field {
   return []ent.Field{
		field.String("detail").NotEmpty(),
		field.Time("added_time"),
   }
}
 
// Edges of the Medicalfile.
func (Medicalfile) Edges() []ent.Edge {
   return []ent.Edge{
   edge.From("dentist", Dentist.Type).Ref("medicalfiles").Unique(),
   edge.From("patient", Patient.Type).Ref("medicalfiles").Unique(),
   edge.From("employee", Employee.Type).Ref("medicalfiles").Unique(),

   }
}

