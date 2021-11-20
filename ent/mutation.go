// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"entgo.io/bug/ent/match"
	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/team"
	"entgo.io/bug/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMatch = "Match"
	TypeTeam  = "Team"
	TypeUser  = "User"
)

// MatchMutation represents an operation that mutates the Match nodes in the graph.
type MatchMutation struct {
	config
	op               Op
	typ              string
	id               *int
	start_date       *time.Time
	clearedFields    map[string]struct{}
	home_team        *int
	clearedhome_team bool
	away_team        *int
	clearedaway_team bool
	done             bool
	oldValue         func(context.Context) (*Match, error)
	predicates       []predicate.Match
}

var _ ent.Mutation = (*MatchMutation)(nil)

// matchOption allows management of the mutation configuration using functional options.
type matchOption func(*MatchMutation)

// newMatchMutation creates new mutation for the Match entity.
func newMatchMutation(c config, op Op, opts ...matchOption) *MatchMutation {
	m := &MatchMutation{
		config:        c,
		op:            op,
		typ:           TypeMatch,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMatchID sets the ID field of the mutation.
func withMatchID(id int) matchOption {
	return func(m *MatchMutation) {
		var (
			err   error
			once  sync.Once
			value *Match
		)
		m.oldValue = func(ctx context.Context) (*Match, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Match.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMatch sets the old Match of the mutation.
func withMatch(node *Match) matchOption {
	return func(m *MatchMutation) {
		m.oldValue = func(context.Context) (*Match, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MatchMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MatchMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Match entities.
func (m *MatchMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MatchMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetStartDate sets the "start_date" field.
func (m *MatchMutation) SetStartDate(t time.Time) {
	m.start_date = &t
}

// StartDate returns the value of the "start_date" field in the mutation.
func (m *MatchMutation) StartDate() (r time.Time, exists bool) {
	v := m.start_date
	if v == nil {
		return
	}
	return *v, true
}

// OldStartDate returns the old "start_date" field's value of the Match entity.
// If the Match object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MatchMutation) OldStartDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStartDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStartDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStartDate: %w", err)
	}
	return oldValue.StartDate, nil
}

// ResetStartDate resets all changes to the "start_date" field.
func (m *MatchMutation) ResetStartDate() {
	m.start_date = nil
}

// SetHomeTeamID sets the "home_team" edge to the Team entity by id.
func (m *MatchMutation) SetHomeTeamID(id int) {
	m.home_team = &id
}

// ClearHomeTeam clears the "home_team" edge to the Team entity.
func (m *MatchMutation) ClearHomeTeam() {
	m.clearedhome_team = true
}

// HomeTeamCleared reports if the "home_team" edge to the Team entity was cleared.
func (m *MatchMutation) HomeTeamCleared() bool {
	return m.clearedhome_team
}

// HomeTeamID returns the "home_team" edge ID in the mutation.
func (m *MatchMutation) HomeTeamID() (id int, exists bool) {
	if m.home_team != nil {
		return *m.home_team, true
	}
	return
}

// HomeTeamIDs returns the "home_team" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// HomeTeamID instead. It exists only for internal usage by the builders.
func (m *MatchMutation) HomeTeamIDs() (ids []int) {
	if id := m.home_team; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetHomeTeam resets all changes to the "home_team" edge.
func (m *MatchMutation) ResetHomeTeam() {
	m.home_team = nil
	m.clearedhome_team = false
}

// SetAwayTeamID sets the "away_team" edge to the Team entity by id.
func (m *MatchMutation) SetAwayTeamID(id int) {
	m.away_team = &id
}

// ClearAwayTeam clears the "away_team" edge to the Team entity.
func (m *MatchMutation) ClearAwayTeam() {
	m.clearedaway_team = true
}

// AwayTeamCleared reports if the "away_team" edge to the Team entity was cleared.
func (m *MatchMutation) AwayTeamCleared() bool {
	return m.clearedaway_team
}

// AwayTeamID returns the "away_team" edge ID in the mutation.
func (m *MatchMutation) AwayTeamID() (id int, exists bool) {
	if m.away_team != nil {
		return *m.away_team, true
	}
	return
}

// AwayTeamIDs returns the "away_team" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AwayTeamID instead. It exists only for internal usage by the builders.
func (m *MatchMutation) AwayTeamIDs() (ids []int) {
	if id := m.away_team; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAwayTeam resets all changes to the "away_team" edge.
func (m *MatchMutation) ResetAwayTeam() {
	m.away_team = nil
	m.clearedaway_team = false
}

// Where appends a list predicates to the MatchMutation builder.
func (m *MatchMutation) Where(ps ...predicate.Match) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MatchMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Match).
func (m *MatchMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MatchMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.start_date != nil {
		fields = append(fields, match.FieldStartDate)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MatchMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case match.FieldStartDate:
		return m.StartDate()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MatchMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case match.FieldStartDate:
		return m.OldStartDate(ctx)
	}
	return nil, fmt.Errorf("unknown Match field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MatchMutation) SetField(name string, value ent.Value) error {
	switch name {
	case match.FieldStartDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStartDate(v)
		return nil
	}
	return fmt.Errorf("unknown Match field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MatchMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MatchMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MatchMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Match numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MatchMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MatchMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MatchMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Match nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MatchMutation) ResetField(name string) error {
	switch name {
	case match.FieldStartDate:
		m.ResetStartDate()
		return nil
	}
	return fmt.Errorf("unknown Match field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MatchMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.home_team != nil {
		edges = append(edges, match.EdgeHomeTeam)
	}
	if m.away_team != nil {
		edges = append(edges, match.EdgeAwayTeam)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MatchMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case match.EdgeHomeTeam:
		if id := m.home_team; id != nil {
			return []ent.Value{*id}
		}
	case match.EdgeAwayTeam:
		if id := m.away_team; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MatchMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MatchMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MatchMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedhome_team {
		edges = append(edges, match.EdgeHomeTeam)
	}
	if m.clearedaway_team {
		edges = append(edges, match.EdgeAwayTeam)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MatchMutation) EdgeCleared(name string) bool {
	switch name {
	case match.EdgeHomeTeam:
		return m.clearedhome_team
	case match.EdgeAwayTeam:
		return m.clearedaway_team
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MatchMutation) ClearEdge(name string) error {
	switch name {
	case match.EdgeHomeTeam:
		m.ClearHomeTeam()
		return nil
	case match.EdgeAwayTeam:
		m.ClearAwayTeam()
		return nil
	}
	return fmt.Errorf("unknown Match unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MatchMutation) ResetEdge(name string) error {
	switch name {
	case match.EdgeHomeTeam:
		m.ResetHomeTeam()
		return nil
	case match.EdgeAwayTeam:
		m.ResetAwayTeam()
		return nil
	}
	return fmt.Errorf("unknown Match edge %s", name)
}

// TeamMutation represents an operation that mutates the Team nodes in the graph.
type TeamMutation struct {
	config
	op                  Op
	typ                 string
	id                  *int
	name                *string
	clearedFields       map[string]struct{}
	home_matches        map[int]struct{}
	removedhome_matches map[int]struct{}
	clearedhome_matches bool
	away_matches        map[int]struct{}
	removedaway_matches map[int]struct{}
	clearedaway_matches bool
	done                bool
	oldValue            func(context.Context) (*Team, error)
	predicates          []predicate.Team
}

var _ ent.Mutation = (*TeamMutation)(nil)

// teamOption allows management of the mutation configuration using functional options.
type teamOption func(*TeamMutation)

// newTeamMutation creates new mutation for the Team entity.
func newTeamMutation(c config, op Op, opts ...teamOption) *TeamMutation {
	m := &TeamMutation{
		config:        c,
		op:            op,
		typ:           TypeTeam,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTeamID sets the ID field of the mutation.
func withTeamID(id int) teamOption {
	return func(m *TeamMutation) {
		var (
			err   error
			once  sync.Once
			value *Team
		)
		m.oldValue = func(ctx context.Context) (*Team, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Team.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTeam sets the old Team of the mutation.
func withTeam(node *Team) teamOption {
	return func(m *TeamMutation) {
		m.oldValue = func(context.Context) (*Team, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TeamMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TeamMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Team entities.
func (m *TeamMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TeamMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *TeamMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TeamMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Team entity.
// If the Team object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TeamMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TeamMutation) ResetName() {
	m.name = nil
}

// AddHomeMatchIDs adds the "home_matches" edge to the Match entity by ids.
func (m *TeamMutation) AddHomeMatchIDs(ids ...int) {
	if m.home_matches == nil {
		m.home_matches = make(map[int]struct{})
	}
	for i := range ids {
		m.home_matches[ids[i]] = struct{}{}
	}
}

// ClearHomeMatches clears the "home_matches" edge to the Match entity.
func (m *TeamMutation) ClearHomeMatches() {
	m.clearedhome_matches = true
}

// HomeMatchesCleared reports if the "home_matches" edge to the Match entity was cleared.
func (m *TeamMutation) HomeMatchesCleared() bool {
	return m.clearedhome_matches
}

// RemoveHomeMatchIDs removes the "home_matches" edge to the Match entity by IDs.
func (m *TeamMutation) RemoveHomeMatchIDs(ids ...int) {
	if m.removedhome_matches == nil {
		m.removedhome_matches = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.home_matches, ids[i])
		m.removedhome_matches[ids[i]] = struct{}{}
	}
}

// RemovedHomeMatches returns the removed IDs of the "home_matches" edge to the Match entity.
func (m *TeamMutation) RemovedHomeMatchesIDs() (ids []int) {
	for id := range m.removedhome_matches {
		ids = append(ids, id)
	}
	return
}

// HomeMatchesIDs returns the "home_matches" edge IDs in the mutation.
func (m *TeamMutation) HomeMatchesIDs() (ids []int) {
	for id := range m.home_matches {
		ids = append(ids, id)
	}
	return
}

// ResetHomeMatches resets all changes to the "home_matches" edge.
func (m *TeamMutation) ResetHomeMatches() {
	m.home_matches = nil
	m.clearedhome_matches = false
	m.removedhome_matches = nil
}

// AddAwayMatchIDs adds the "away_matches" edge to the Match entity by ids.
func (m *TeamMutation) AddAwayMatchIDs(ids ...int) {
	if m.away_matches == nil {
		m.away_matches = make(map[int]struct{})
	}
	for i := range ids {
		m.away_matches[ids[i]] = struct{}{}
	}
}

// ClearAwayMatches clears the "away_matches" edge to the Match entity.
func (m *TeamMutation) ClearAwayMatches() {
	m.clearedaway_matches = true
}

// AwayMatchesCleared reports if the "away_matches" edge to the Match entity was cleared.
func (m *TeamMutation) AwayMatchesCleared() bool {
	return m.clearedaway_matches
}

// RemoveAwayMatchIDs removes the "away_matches" edge to the Match entity by IDs.
func (m *TeamMutation) RemoveAwayMatchIDs(ids ...int) {
	if m.removedaway_matches == nil {
		m.removedaway_matches = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.away_matches, ids[i])
		m.removedaway_matches[ids[i]] = struct{}{}
	}
}

// RemovedAwayMatches returns the removed IDs of the "away_matches" edge to the Match entity.
func (m *TeamMutation) RemovedAwayMatchesIDs() (ids []int) {
	for id := range m.removedaway_matches {
		ids = append(ids, id)
	}
	return
}

// AwayMatchesIDs returns the "away_matches" edge IDs in the mutation.
func (m *TeamMutation) AwayMatchesIDs() (ids []int) {
	for id := range m.away_matches {
		ids = append(ids, id)
	}
	return
}

// ResetAwayMatches resets all changes to the "away_matches" edge.
func (m *TeamMutation) ResetAwayMatches() {
	m.away_matches = nil
	m.clearedaway_matches = false
	m.removedaway_matches = nil
}

// Where appends a list predicates to the TeamMutation builder.
func (m *TeamMutation) Where(ps ...predicate.Team) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TeamMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Team).
func (m *TeamMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TeamMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, team.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TeamMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case team.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TeamMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case team.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Team field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TeamMutation) SetField(name string, value ent.Value) error {
	switch name {
	case team.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Team field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TeamMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TeamMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TeamMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Team numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TeamMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TeamMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TeamMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Team nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TeamMutation) ResetField(name string) error {
	switch name {
	case team.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Team field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TeamMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.home_matches != nil {
		edges = append(edges, team.EdgeHomeMatches)
	}
	if m.away_matches != nil {
		edges = append(edges, team.EdgeAwayMatches)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TeamMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case team.EdgeHomeMatches:
		ids := make([]ent.Value, 0, len(m.home_matches))
		for id := range m.home_matches {
			ids = append(ids, id)
		}
		return ids
	case team.EdgeAwayMatches:
		ids := make([]ent.Value, 0, len(m.away_matches))
		for id := range m.away_matches {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TeamMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedhome_matches != nil {
		edges = append(edges, team.EdgeHomeMatches)
	}
	if m.removedaway_matches != nil {
		edges = append(edges, team.EdgeAwayMatches)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TeamMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case team.EdgeHomeMatches:
		ids := make([]ent.Value, 0, len(m.removedhome_matches))
		for id := range m.removedhome_matches {
			ids = append(ids, id)
		}
		return ids
	case team.EdgeAwayMatches:
		ids := make([]ent.Value, 0, len(m.removedaway_matches))
		for id := range m.removedaway_matches {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TeamMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedhome_matches {
		edges = append(edges, team.EdgeHomeMatches)
	}
	if m.clearedaway_matches {
		edges = append(edges, team.EdgeAwayMatches)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TeamMutation) EdgeCleared(name string) bool {
	switch name {
	case team.EdgeHomeMatches:
		return m.clearedhome_matches
	case team.EdgeAwayMatches:
		return m.clearedaway_matches
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TeamMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Team unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TeamMutation) ResetEdge(name string) error {
	switch name {
	case team.EdgeHomeMatches:
		m.ResetHomeMatches()
		return nil
	case team.EdgeAwayMatches:
		m.ResetAwayMatches()
		return nil
	}
	return fmt.Errorf("unknown Team edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
