// Code generated by entc, DO NOT EDIT.

package training

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team09/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Branch applies equality check predicate on the "branch" field. It's identical to BranchEQ.
func Branch(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBranch), v))
	})
}

// Dateone applies equality check predicate on the "dateone" field. It's identical to DateoneEQ.
func Dateone(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateone), v))
	})
}

// Datetwo applies equality check predicate on the "datetwo" field. It's identical to DatetwoEQ.
func Datetwo(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatetwo), v))
	})
}

// BranchEQ applies the EQ predicate on the "branch" field.
func BranchEQ(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBranch), v))
	})
}

// BranchNEQ applies the NEQ predicate on the "branch" field.
func BranchNEQ(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBranch), v))
	})
}

// BranchIn applies the In predicate on the "branch" field.
func BranchIn(vs ...string) predicate.Training {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Training(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBranch), v...))
	})
}

// BranchNotIn applies the NotIn predicate on the "branch" field.
func BranchNotIn(vs ...string) predicate.Training {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Training(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBranch), v...))
	})
}

// BranchGT applies the GT predicate on the "branch" field.
func BranchGT(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBranch), v))
	})
}

// BranchGTE applies the GTE predicate on the "branch" field.
func BranchGTE(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBranch), v))
	})
}

// BranchLT applies the LT predicate on the "branch" field.
func BranchLT(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBranch), v))
	})
}

// BranchLTE applies the LTE predicate on the "branch" field.
func BranchLTE(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBranch), v))
	})
}

// BranchContains applies the Contains predicate on the "branch" field.
func BranchContains(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBranch), v))
	})
}

// BranchHasPrefix applies the HasPrefix predicate on the "branch" field.
func BranchHasPrefix(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBranch), v))
	})
}

// BranchHasSuffix applies the HasSuffix predicate on the "branch" field.
func BranchHasSuffix(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBranch), v))
	})
}

// BranchEqualFold applies the EqualFold predicate on the "branch" field.
func BranchEqualFold(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBranch), v))
	})
}

// BranchContainsFold applies the ContainsFold predicate on the "branch" field.
func BranchContainsFold(v string) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBranch), v))
	})
}

// DateoneEQ applies the EQ predicate on the "dateone" field.
func DateoneEQ(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateone), v))
	})
}

// DateoneNEQ applies the NEQ predicate on the "dateone" field.
func DateoneNEQ(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateone), v))
	})
}

// DateoneIn applies the In predicate on the "dateone" field.
func DateoneIn(vs ...time.Time) predicate.Training {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Training(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateone), v...))
	})
}

// DateoneNotIn applies the NotIn predicate on the "dateone" field.
func DateoneNotIn(vs ...time.Time) predicate.Training {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Training(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateone), v...))
	})
}

// DateoneGT applies the GT predicate on the "dateone" field.
func DateoneGT(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateone), v))
	})
}

// DateoneGTE applies the GTE predicate on the "dateone" field.
func DateoneGTE(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateone), v))
	})
}

// DateoneLT applies the LT predicate on the "dateone" field.
func DateoneLT(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateone), v))
	})
}

// DateoneLTE applies the LTE predicate on the "dateone" field.
func DateoneLTE(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateone), v))
	})
}

// DatetwoEQ applies the EQ predicate on the "datetwo" field.
func DatetwoEQ(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatetwo), v))
	})
}

// DatetwoNEQ applies the NEQ predicate on the "datetwo" field.
func DatetwoNEQ(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDatetwo), v))
	})
}

// DatetwoIn applies the In predicate on the "datetwo" field.
func DatetwoIn(vs ...time.Time) predicate.Training {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Training(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDatetwo), v...))
	})
}

// DatetwoNotIn applies the NotIn predicate on the "datetwo" field.
func DatetwoNotIn(vs ...time.Time) predicate.Training {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Training(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDatetwo), v...))
	})
}

// DatetwoGT applies the GT predicate on the "datetwo" field.
func DatetwoGT(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDatetwo), v))
	})
}

// DatetwoGTE applies the GTE predicate on the "datetwo" field.
func DatetwoGTE(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDatetwo), v))
	})
}

// DatetwoLT applies the LT predicate on the "datetwo" field.
func DatetwoLT(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDatetwo), v))
	})
}

// DatetwoLTE applies the LTE predicate on the "datetwo" field.
func DatetwoLTE(v time.Time) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDatetwo), v))
	})
}

// HasCourse applies the HasEdge predicate on the "course" edge.
func HasCourse() predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourseWith applies the HasEdge predicate on the "course" edge with a given conditions (other predicates).
func HasCourseWith(preds ...predicate.Course) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDoctor applies the HasEdge predicate on the "doctor" edge.
func HasDoctor() predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorTable, DoctorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorWith applies the HasEdge predicate on the "doctor" edge with a given conditions (other predicates).
func HasDoctorWith(preds ...predicate.Doctor) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorTable, DoctorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Training) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Training) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
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
func Not(p predicate.Training) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		p(s.Not())
	})
}
