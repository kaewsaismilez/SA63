// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kaews/app/ent/medicalfile"
	"github.com/kaews/app/ent/patient"
	"github.com/kaews/app/ent/predicate"
)

// PatientUpdate is the builder for updating Patient entities.
type PatientUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientMutation
	predicates []predicate.Patient
}

// Where adds a new predicate for the builder.
func (pu *PatientUpdate) Where(ps ...predicate.Patient) *PatientUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the name field.
func (pu *PatientUpdate) SetName(s string) *PatientUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetAge sets the age field.
func (pu *PatientUpdate) SetAge(i int) *PatientUpdate {
	pu.mutation.ResetAge()
	pu.mutation.SetAge(i)
	return pu
}

// AddAge adds i to age.
func (pu *PatientUpdate) AddAge(i int) *PatientUpdate {
	pu.mutation.AddAge(i)
	return pu
}

// AddMedicalfileIDs adds the medicalfiles edge to Medicalfile by ids.
func (pu *PatientUpdate) AddMedicalfileIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddMedicalfileIDs(ids...)
	return pu
}

// AddMedicalfiles adds the medicalfiles edges to Medicalfile.
func (pu *PatientUpdate) AddMedicalfiles(m ...*Medicalfile) *PatientUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pu.AddMedicalfileIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pu *PatientUpdate) Mutation() *PatientMutation {
	return pu.mutation
}

// RemoveMedicalfileIDs removes the medicalfiles edge to Medicalfile by ids.
func (pu *PatientUpdate) RemoveMedicalfileIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemoveMedicalfileIDs(ids...)
	return pu
}

// RemoveMedicalfiles removes medicalfiles edges to Medicalfile.
func (pu *PatientUpdate) RemoveMedicalfiles(m ...*Medicalfile) *PatientUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pu.RemoveMedicalfileIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Name(); ok {
		if err := patient.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Age(); ok {
		if err := patient.AgeValidator(v); err != nil {
			return 0, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
	}
	if value, ok := pu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
	}
	if value, ok := pu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
	}
	if nodes := pu.mutation.RemovedMedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.MedicalfilesTable,
			Columns: []string{patient.MedicalfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.MedicalfilesTable,
			Columns: []string{patient.MedicalfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientUpdateOne is the builder for updating a single Patient entity.
type PatientUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// SetName sets the name field.
func (puo *PatientUpdateOne) SetName(s string) *PatientUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetAge sets the age field.
func (puo *PatientUpdateOne) SetAge(i int) *PatientUpdateOne {
	puo.mutation.ResetAge()
	puo.mutation.SetAge(i)
	return puo
}

// AddAge adds i to age.
func (puo *PatientUpdateOne) AddAge(i int) *PatientUpdateOne {
	puo.mutation.AddAge(i)
	return puo
}

// AddMedicalfileIDs adds the medicalfiles edge to Medicalfile by ids.
func (puo *PatientUpdateOne) AddMedicalfileIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddMedicalfileIDs(ids...)
	return puo
}

// AddMedicalfiles adds the medicalfiles edges to Medicalfile.
func (puo *PatientUpdateOne) AddMedicalfiles(m ...*Medicalfile) *PatientUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return puo.AddMedicalfileIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (puo *PatientUpdateOne) Mutation() *PatientMutation {
	return puo.mutation
}

// RemoveMedicalfileIDs removes the medicalfiles edge to Medicalfile by ids.
func (puo *PatientUpdateOne) RemoveMedicalfileIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemoveMedicalfileIDs(ids...)
	return puo
}

// RemoveMedicalfiles removes medicalfiles edges to Medicalfile.
func (puo *PatientUpdateOne) RemoveMedicalfiles(m ...*Medicalfile) *PatientUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return puo.RemoveMedicalfileIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PatientUpdateOne) Save(ctx context.Context) (*Patient, error) {
	if v, ok := puo.mutation.Name(); ok {
		if err := patient.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Age(); ok {
		if err := patient.AgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}

	var (
		err  error
		node *Patient
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientUpdateOne) SaveX(ctx context.Context) *Patient {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientUpdateOne) sqlSave(ctx context.Context) (pa *Patient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patient.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
	}
	if value, ok := puo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
	}
	if value, ok := puo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldAge,
		})
	}
	if nodes := puo.mutation.RemovedMedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.MedicalfilesTable,
			Columns: []string{patient.MedicalfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MedicalfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.MedicalfilesTable,
			Columns: []string{patient.MedicalfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Patient{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
