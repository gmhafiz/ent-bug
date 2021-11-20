// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/bug/ent/match"
	"entgo.io/bug/ent/team"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MatchCreate is the builder for creating a Match entity.
type MatchCreate struct {
	config
	mutation *MatchMutation
	hooks    []Hook
}

// SetStartDate sets the "start_date" field.
func (mc *MatchCreate) SetStartDate(t time.Time) *MatchCreate {
	mc.mutation.SetStartDate(t)
	return mc
}

// SetID sets the "id" field.
func (mc *MatchCreate) SetID(i int) *MatchCreate {
	mc.mutation.SetID(i)
	return mc
}

// SetHomeTeamID sets the "home_team" edge to the Team entity by ID.
func (mc *MatchCreate) SetHomeTeamID(id int) *MatchCreate {
	mc.mutation.SetHomeTeamID(id)
	return mc
}

// SetNillableHomeTeamID sets the "home_team" edge to the Team entity by ID if the given value is not nil.
func (mc *MatchCreate) SetNillableHomeTeamID(id *int) *MatchCreate {
	if id != nil {
		mc = mc.SetHomeTeamID(*id)
	}
	return mc
}

// SetHomeTeam sets the "home_team" edge to the Team entity.
func (mc *MatchCreate) SetHomeTeam(t *Team) *MatchCreate {
	return mc.SetHomeTeamID(t.ID)
}

// SetAwayTeamID sets the "away_team" edge to the Team entity by ID.
func (mc *MatchCreate) SetAwayTeamID(id int) *MatchCreate {
	mc.mutation.SetAwayTeamID(id)
	return mc
}

// SetNillableAwayTeamID sets the "away_team" edge to the Team entity by ID if the given value is not nil.
func (mc *MatchCreate) SetNillableAwayTeamID(id *int) *MatchCreate {
	if id != nil {
		mc = mc.SetAwayTeamID(*id)
	}
	return mc
}

// SetAwayTeam sets the "away_team" edge to the Team entity.
func (mc *MatchCreate) SetAwayTeam(t *Team) *MatchCreate {
	return mc.SetAwayTeamID(t.ID)
}

// Mutation returns the MatchMutation object of the builder.
func (mc *MatchCreate) Mutation() *MatchMutation {
	return mc.mutation
}

// Save creates the Match in the database.
func (mc *MatchCreate) Save(ctx context.Context) (*Match, error) {
	var (
		err  error
		node *Match
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MatchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MatchCreate) SaveX(ctx context.Context) *Match {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MatchCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MatchCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MatchCreate) check() error {
	if _, ok := mc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "Match.start_date"`)}
	}
	return nil
}

func (mc *MatchCreate) sqlSave(ctx context.Context) (*Match, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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

func (mc *MatchCreate) createSpec() (*Match, *sqlgraph.CreateSpec) {
	var (
		_node = &Match{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: match.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: match.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.StartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: match.FieldStartDate,
		})
		_node.StartDate = value
	}
	if nodes := mc.mutation.HomeTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   match.HomeTeamTable,
			Columns: []string{match.HomeTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.match_home_team = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.AwayTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   match.AwayTeamTable,
			Columns: []string{match.AwayTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.match_away_team = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MatchCreateBulk is the builder for creating many Match entities in bulk.
type MatchCreateBulk struct {
	config
	builders []*MatchCreate
}

// Save creates the Match entities in the database.
func (mcb *MatchCreateBulk) Save(ctx context.Context) ([]*Match, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Match, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MatchMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MatchCreateBulk) SaveX(ctx context.Context) []*Match {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MatchCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MatchCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
