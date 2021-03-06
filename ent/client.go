// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"entgo.io/bug/ent/migrate"

	"entgo.io/bug/ent/match"
	"entgo.io/bug/ent/team"
	"entgo.io/bug/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Match is the client for interacting with the Match builders.
	Match *MatchClient
	// Team is the client for interacting with the Team builders.
	Team *TeamClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Match = NewMatchClient(c.config)
	c.Team = NewTeamClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Match:  NewMatchClient(cfg),
		Team:   NewTeamClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Match:  NewMatchClient(cfg),
		Team:   NewTeamClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Match.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Match.Use(hooks...)
	c.Team.Use(hooks...)
	c.User.Use(hooks...)
}

// MatchClient is a client for the Match schema.
type MatchClient struct {
	config
}

// NewMatchClient returns a client for the Match from the given config.
func NewMatchClient(c config) *MatchClient {
	return &MatchClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `match.Hooks(f(g(h())))`.
func (c *MatchClient) Use(hooks ...Hook) {
	c.hooks.Match = append(c.hooks.Match, hooks...)
}

// Create returns a create builder for Match.
func (c *MatchClient) Create() *MatchCreate {
	mutation := newMatchMutation(c.config, OpCreate)
	return &MatchCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Match entities.
func (c *MatchClient) CreateBulk(builders ...*MatchCreate) *MatchCreateBulk {
	return &MatchCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Match.
func (c *MatchClient) Update() *MatchUpdate {
	mutation := newMatchMutation(c.config, OpUpdate)
	return &MatchUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MatchClient) UpdateOne(m *Match) *MatchUpdateOne {
	mutation := newMatchMutation(c.config, OpUpdateOne, withMatch(m))
	return &MatchUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MatchClient) UpdateOneID(id int) *MatchUpdateOne {
	mutation := newMatchMutation(c.config, OpUpdateOne, withMatchID(id))
	return &MatchUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Match.
func (c *MatchClient) Delete() *MatchDelete {
	mutation := newMatchMutation(c.config, OpDelete)
	return &MatchDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MatchClient) DeleteOne(m *Match) *MatchDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MatchClient) DeleteOneID(id int) *MatchDeleteOne {
	builder := c.Delete().Where(match.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MatchDeleteOne{builder}
}

// Query returns a query builder for Match.
func (c *MatchClient) Query() *MatchQuery {
	return &MatchQuery{
		config: c.config,
	}
}

// Get returns a Match entity by its id.
func (c *MatchClient) Get(ctx context.Context, id int) (*Match, error) {
	return c.Query().Where(match.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MatchClient) GetX(ctx context.Context, id int) *Match {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHomeTeam queries the home_team edge of a Match.
func (c *MatchClient) QueryHomeTeam(m *Match) *TeamQuery {
	query := &TeamQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(match.Table, match.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, match.HomeTeamTable, match.HomeTeamColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAwayTeam queries the away_team edge of a Match.
func (c *MatchClient) QueryAwayTeam(m *Match) *TeamQuery {
	query := &TeamQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(match.Table, match.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, match.AwayTeamTable, match.AwayTeamColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MatchClient) Hooks() []Hook {
	return c.hooks.Match
}

// TeamClient is a client for the Team schema.
type TeamClient struct {
	config
}

// NewTeamClient returns a client for the Team from the given config.
func NewTeamClient(c config) *TeamClient {
	return &TeamClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `team.Hooks(f(g(h())))`.
func (c *TeamClient) Use(hooks ...Hook) {
	c.hooks.Team = append(c.hooks.Team, hooks...)
}

// Create returns a create builder for Team.
func (c *TeamClient) Create() *TeamCreate {
	mutation := newTeamMutation(c.config, OpCreate)
	return &TeamCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Team entities.
func (c *TeamClient) CreateBulk(builders ...*TeamCreate) *TeamCreateBulk {
	return &TeamCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Team.
func (c *TeamClient) Update() *TeamUpdate {
	mutation := newTeamMutation(c.config, OpUpdate)
	return &TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TeamClient) UpdateOne(t *Team) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeam(t))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TeamClient) UpdateOneID(id int) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeamID(id))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Team.
func (c *TeamClient) Delete() *TeamDelete {
	mutation := newTeamMutation(c.config, OpDelete)
	return &TeamDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TeamClient) DeleteOne(t *Team) *TeamDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TeamClient) DeleteOneID(id int) *TeamDeleteOne {
	builder := c.Delete().Where(team.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TeamDeleteOne{builder}
}

// Query returns a query builder for Team.
func (c *TeamClient) Query() *TeamQuery {
	return &TeamQuery{
		config: c.config,
	}
}

// Get returns a Team entity by its id.
func (c *TeamClient) Get(ctx context.Context, id int) (*Team, error) {
	return c.Query().Where(team.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TeamClient) GetX(ctx context.Context, id int) *Team {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHomeMatches queries the home_matches edge of a Team.
func (c *TeamClient) QueryHomeMatches(t *Team) *MatchQuery {
	query := &MatchQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(match.Table, match.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, team.HomeMatchesTable, team.HomeMatchesColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAwayMatches queries the away_matches edge of a Team.
func (c *TeamClient) QueryAwayMatches(t *Team) *MatchQuery {
	query := &MatchQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(match.Table, match.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, team.AwayMatchesTable, team.AwayMatchesColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TeamClient) Hooks() []Hook {
	return c.hooks.Team
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
