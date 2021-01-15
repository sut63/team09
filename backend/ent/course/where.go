// Code generated by entc, DO NOT EDIT.

package course

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team09/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Namecourse applies equality check predicate on the "namecourse" field. It's identical to NamecourseEQ.
func Namecourse(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNamecourse), v))
	})
}

// NamecourseEQ applies the EQ predicate on the "namecourse" field.
func NamecourseEQ(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNamecourse), v))
	})
}

// NamecourseNEQ applies the NEQ predicate on the "namecourse" field.
func NamecourseNEQ(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNamecourse), v))
	})
}

// NamecourseIn applies the In predicate on the "namecourse" field.
func NamecourseIn(vs ...string) predicate.Course {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Course(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNamecourse), v...))
	})
}

// NamecourseNotIn applies the NotIn predicate on the "namecourse" field.
func NamecourseNotIn(vs ...string) predicate.Course {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Course(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNamecourse), v...))
	})
}

// NamecourseGT applies the GT predicate on the "namecourse" field.
func NamecourseGT(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNamecourse), v))
	})
}

// NamecourseGTE applies the GTE predicate on the "namecourse" field.
func NamecourseGTE(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNamecourse), v))
	})
}

// NamecourseLT applies the LT predicate on the "namecourse" field.
func NamecourseLT(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNamecourse), v))
	})
}

// NamecourseLTE applies the LTE predicate on the "namecourse" field.
func NamecourseLTE(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNamecourse), v))
	})
}

// NamecourseContains applies the Contains predicate on the "namecourse" field.
func NamecourseContains(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNamecourse), v))
	})
}

// NamecourseHasPrefix applies the HasPrefix predicate on the "namecourse" field.
func NamecourseHasPrefix(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNamecourse), v))
	})
}

// NamecourseHasSuffix applies the HasSuffix predicate on the "namecourse" field.
func NamecourseHasSuffix(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNamecourse), v))
	})
}

// NamecourseEqualFold applies the EqualFold predicate on the "namecourse" field.
func NamecourseEqualFold(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNamecourse), v))
	})
}

// NamecourseContainsFold applies the ContainsFold predicate on the "namecourse" field.
func NamecourseContainsFold(v string) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNamecourse), v))
	})
}

// HasTrainings applies the HasEdge predicate on the "trainings" edge.
func HasTrainings() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TrainingsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TrainingsTable, TrainingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTrainingsWith applies the HasEdge predicate on the "trainings" edge with a given conditions (other predicates).
func HasTrainingsWith(preds ...predicate.Training) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TrainingsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TrainingsTable, TrainingsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDetails applies the HasEdge predicate on the "details" edge.
func HasDetails() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DetailsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DetailsTable, DetailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDetailsWith applies the HasEdge predicate on the "details" edge with a given conditions (other predicates).
func HasDetailsWith(preds ...predicate.Detail) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DetailsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DetailsTable, DetailsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
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
func Not(p predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		p(s.Not())
	})
}
