// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/match"
	"entgo.io/bug/ent/team"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeamCreate is the builder for creating a Team entity.
type TeamCreate struct {
	config
	mutation *TeamMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TeamCreate) SetName(s string) *TeamCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TeamCreate) SetID(i int) *TeamCreate {
	tc.mutation.SetID(i)
	return tc
}

// AddHomeMatchIDs adds the "home_matches" edge to the Match entity by IDs.
func (tc *TeamCreate) AddHomeMatchIDs(ids ...int) *TeamCreate {
	tc.mutation.AddHomeMatchIDs(ids...)
	return tc
}

// AddHomeMatches adds the "home_matches" edges to the Match entity.
func (tc *TeamCreate) AddHomeMatches(m ...*Match) *TeamCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tc.AddHomeMatchIDs(ids...)
}

// AddAwayMatchIDs adds the "away_matches" edge to the Match entity by IDs.
func (tc *TeamCreate) AddAwayMatchIDs(ids ...int) *TeamCreate {
	tc.mutation.AddAwayMatchIDs(ids...)
	return tc
}

// AddAwayMatches adds the "away_matches" edges to the Match entity.
func (tc *TeamCreate) AddAwayMatches(m ...*Match) *TeamCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tc.AddAwayMatchIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tc *TeamCreate) Mutation() *TeamMutation {
	return tc.mutation
}

// Save creates the Team in the database.
func (tc *TeamCreate) Save(ctx context.Context) (*Team, error) {
	var (
		err  error
		node *Team
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeamCreate) SaveX(ctx context.Context) *Team {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeamCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeamCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeamCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Team.name"`)}
	}
	return nil
}

func (tc *TeamCreate) sqlSave(ctx context.Context) (*Team, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (tc *TeamCreate) createSpec() (*Team, *sqlgraph.CreateSpec) {
	var (
		_node = &Team{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: team.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: team.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: team.FieldName,
		})
		_node.Name = value
	}
	if nodes := tc.mutation.HomeMatchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.HomeMatchesTable,
			Columns: []string{team.HomeMatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: match.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.AwayMatchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.AwayMatchesTable,
			Columns: []string{team.AwayMatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: match.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TeamCreateBulk is the builder for creating many Team entities in bulk.
type TeamCreateBulk struct {
	config
	builders []*TeamCreate
}

// Save creates the Team entities in the database.
func (tcb *TeamCreateBulk) Save(ctx context.Context) ([]*Team, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Team, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeamMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeamCreateBulk) SaveX(ctx context.Context) []*Team {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeamCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeamCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
