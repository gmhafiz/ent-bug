package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("home_id", Match.Type),
		edge.To("away_id", Match.Type),
	}

	// Rotem suggestion
	return []ent.Edge{
		edge.From("home_matches", Match.Type).Ref("home_team"),
		edge.From("away_matches", Match.Type).Ref("away_team"),
	}
}
