// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/bug/ent/match"
	"entgo.io/ent/dialect/sql"
)

// Match is the model entity for the Match schema.
type Match struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate    time.Time `json:"start_date,omitempty"`
	team_home_id *int
	team_away_id *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Match) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case match.FieldID:
			values[i] = new(sql.NullInt64)
		case match.FieldStartDate:
			values[i] = new(sql.NullTime)
		case match.ForeignKeys[0]: // team_home_id
			values[i] = new(sql.NullInt64)
		case match.ForeignKeys[1]: // team_away_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Match", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Match fields.
func (m *Match) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case match.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case match.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				m.StartDate = value.Time
			}
		case match.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_home_id", value)
			} else if value.Valid {
				m.team_home_id = new(int)
				*m.team_home_id = int(value.Int64)
			}
		case match.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_away_id", value)
			} else if value.Valid {
				m.team_away_id = new(int)
				*m.team_away_id = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Match.
// Note that you need to call Match.Unwrap() before calling this method if this Match
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Match) Update() *MatchUpdateOne {
	return (&MatchClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Match entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Match) Unwrap() *Match {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Match is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Match) String() string {
	var builder strings.Builder
	builder.WriteString("Match(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", start_date=")
	builder.WriteString(m.StartDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Matches is a parsable slice of Match.
type Matches []*Match

func (m Matches) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}