package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// Employee holds the schema definition for the Employee entity.
type Employee struct {
   ent.Schema
}
 
// Fields of the Employee.
func (Employee) Fields() []ent.Field {
   return []ent.Field{
      field.String("name").NotEmpty(),
      field.String("email").NotEmpty(),
		field.String("password").NotEmpty(),
   }

}
 
// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
   return []ent.Edge{
      edge.To("medicalfiles", Medicalfile.Type).StorageKey(edge.Column("employee_id")),
  }
}
