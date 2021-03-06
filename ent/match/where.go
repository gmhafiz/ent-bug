// Code generated by entc, DO NOT EDIT.

package match

import (
	"time"

	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartDate), v))
	})
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Match {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Match(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartDate), v...))
	})
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Match {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Match(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartDate), v...))
	})
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartDate), v))
	})
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartDate), v))
	})
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartDate), v))
	})
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartDate), v))
	})
}

// HasHomeTeam applies the HasEdge predicate on the "home_team" edge.
func HasHomeTeam() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HomeTeamTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HomeTeamTable, HomeTeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHomeTeamWith applies the HasEdge predicate on the "home_team" edge with a given conditions (other predicates).
func HasHomeTeamWith(preds ...predicate.Team) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HomeTeamInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HomeTeamTable, HomeTeamColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAwayTeam applies the HasEdge predicate on the "away_team" edge.
func HasAwayTeam() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AwayTeamTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AwayTeamTable, AwayTeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAwayTeamWith applies the HasEdge predicate on the "away_team" edge with a given conditions (other predicates).
func HasAwayTeamWith(preds ...predicate.Team) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AwayTeamInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AwayTeamTable, AwayTeamColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Match) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Match) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Match) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		p(s.Not())
	})
}
