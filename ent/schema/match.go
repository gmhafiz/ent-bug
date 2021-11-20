package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Match holds the schema definition for the Match entity.
type Match struct {
	ent.Schema
}

// Fields of the Match.
func (Match) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Time("start_date"),
	}
}

// Edges of the Match.
func (Match) Edges() []ent.Edge {
	return nil

	// Rotem suggestion
	return []ent.Edge{
		edge.To("home_team", Team.Type),
		edge.To("away_team", Team.Type),
	}
}
