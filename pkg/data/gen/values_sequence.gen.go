// Code generated by tmpl; DO NOT EDIT.
// https://github.com/benbjohnson/tmpl
//
// Source: values_sequence.gen.go.tmpl

package gen

import (
	"github.com/influxdata/influxdb/v2/models"
	"github.com/influxdata/influxdb/v2/tsdb"
)

type FloatValuesSequence interface {
	Reset()
	Write(v []float64)
}

type timeFloatValuesSequence struct {
	vals  floatArray
	ts    TimestampSequence
	vs    FloatValuesSequence
	count int
	n     int
}

func NewTimeFloatValuesSequence(count int, ts TimestampSequence, vs FloatValuesSequence) TimeValuesSequence {
	return &timeFloatValuesSequence{
		vals:  *newFloatArrayLen(tsdb.DefaultMaxPointsPerBlock),
		ts:    ts,
		vs:    vs,
		count: count,
		n:     count,
	}
}

func (s *timeFloatValuesSequence) Reset() {
	s.ts.Reset()
	s.vs.Reset()
	s.n = s.count
}

func (s *timeFloatValuesSequence) Next() bool {
	if s.n > 0 {
		c := min(s.n, tsdb.DefaultMaxPointsPerBlock)
		s.n -= c
		s.vals.Timestamps = s.vals.Timestamps[:c]
		s.vals.Values = s.vals.Values[:c]

		s.ts.Write(s.vals.Timestamps)
		s.vs.Write(s.vals.Values)
		return true
	}

	return false
}

func (s *timeFloatValuesSequence) Values() Values {
	return &s.vals
}

func (s *timeFloatValuesSequence) ValueType() models.FieldType {
	return models.Float
}

type IntegerValuesSequence interface {
	Reset()
	Write(v []int64)
}

type timeIntegerValuesSequence struct {
	vals  integerArray
	ts    TimestampSequence
	vs    IntegerValuesSequence
	count int
	n     int
}

func NewTimeIntegerValuesSequence(count int, ts TimestampSequence, vs IntegerValuesSequence) TimeValuesSequence {
	return &timeIntegerValuesSequence{
		vals:  *newIntegerArrayLen(tsdb.DefaultMaxPointsPerBlock),
		ts:    ts,
		vs:    vs,
		count: count,
		n:     count,
	}
}

func (s *timeIntegerValuesSequence) Reset() {
	s.ts.Reset()
	s.vs.Reset()
	s.n = s.count
}

func (s *timeIntegerValuesSequence) Next() bool {
	if s.n > 0 {
		c := min(s.n, tsdb.DefaultMaxPointsPerBlock)
		s.n -= c
		s.vals.Timestamps = s.vals.Timestamps[:c]
		s.vals.Values = s.vals.Values[:c]

		s.ts.Write(s.vals.Timestamps)
		s.vs.Write(s.vals.Values)
		return true
	}

	return false
}

func (s *timeIntegerValuesSequence) Values() Values {
	return &s.vals
}

func (s *timeIntegerValuesSequence) ValueType() models.FieldType {
	return models.Integer
}

type UnsignedValuesSequence interface {
	Reset()
	Write(v []uint64)
}

type timeUnsignedValuesSequence struct {
	vals  unsignedArray
	ts    TimestampSequence
	vs    UnsignedValuesSequence
	count int
	n     int
}

func NewTimeUnsignedValuesSequence(count int, ts TimestampSequence, vs UnsignedValuesSequence) TimeValuesSequence {
	return &timeUnsignedValuesSequence{
		vals:  *newUnsignedArrayLen(tsdb.DefaultMaxPointsPerBlock),
		ts:    ts,
		vs:    vs,
		count: count,
		n:     count,
	}
}

func (s *timeUnsignedValuesSequence) Reset() {
	s.ts.Reset()
	s.vs.Reset()
	s.n = s.count
}

func (s *timeUnsignedValuesSequence) Next() bool {
	if s.n > 0 {
		c := min(s.n, tsdb.DefaultMaxPointsPerBlock)
		s.n -= c
		s.vals.Timestamps = s.vals.Timestamps[:c]
		s.vals.Values = s.vals.Values[:c]

		s.ts.Write(s.vals.Timestamps)
		s.vs.Write(s.vals.Values)
		return true
	}

	return false
}

func (s *timeUnsignedValuesSequence) Values() Values {
	return &s.vals
}

func (s *timeUnsignedValuesSequence) ValueType() models.FieldType {
	return models.Unsigned
}

type StringValuesSequence interface {
	Reset()
	Write(v []string)
}

type timeStringValuesSequence struct {
	vals  stringArray
	ts    TimestampSequence
	vs    StringValuesSequence
	count int
	n     int
}

func NewTimeStringValuesSequence(count int, ts TimestampSequence, vs StringValuesSequence) TimeValuesSequence {
	return &timeStringValuesSequence{
		vals:  *newStringArrayLen(tsdb.DefaultMaxPointsPerBlock),
		ts:    ts,
		vs:    vs,
		count: count,
		n:     count,
	}
}

func (s *timeStringValuesSequence) Reset() {
	s.ts.Reset()
	s.vs.Reset()
	s.n = s.count
}

func (s *timeStringValuesSequence) Next() bool {
	if s.n > 0 {
		c := min(s.n, tsdb.DefaultMaxPointsPerBlock)
		s.n -= c
		s.vals.Timestamps = s.vals.Timestamps[:c]
		s.vals.Values = s.vals.Values[:c]

		s.ts.Write(s.vals.Timestamps)
		s.vs.Write(s.vals.Values)
		return true
	}

	return false
}

func (s *timeStringValuesSequence) Values() Values {
	return &s.vals
}

func (s *timeStringValuesSequence) ValueType() models.FieldType {
	return models.String
}

type BooleanValuesSequence interface {
	Reset()
	Write(v []bool)
}

type timeBooleanValuesSequence struct {
	vals  booleanArray
	ts    TimestampSequence
	vs    BooleanValuesSequence
	count int
	n     int
}

func NewTimeBooleanValuesSequence(count int, ts TimestampSequence, vs BooleanValuesSequence) TimeValuesSequence {
	return &timeBooleanValuesSequence{
		vals:  *newBooleanArrayLen(tsdb.DefaultMaxPointsPerBlock),
		ts:    ts,
		vs:    vs,
		count: count,
		n:     count,
	}
}

func (s *timeBooleanValuesSequence) Reset() {
	s.ts.Reset()
	s.vs.Reset()
	s.n = s.count
}

func (s *timeBooleanValuesSequence) Next() bool {
	if s.n > 0 {
		c := min(s.n, tsdb.DefaultMaxPointsPerBlock)
		s.n -= c
		s.vals.Timestamps = s.vals.Timestamps[:c]
		s.vals.Values = s.vals.Values[:c]

		s.ts.Write(s.vals.Timestamps)
		s.vs.Write(s.vals.Values)
		return true
	}

	return false
}

func (s *timeBooleanValuesSequence) Values() Values {
	return &s.vals
}

func (s *timeBooleanValuesSequence) ValueType() models.FieldType {
	return models.Boolean
}
