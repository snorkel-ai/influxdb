// Code generated by tmpl; DO NOT EDIT.
// https://github.com/benbjohnson/tmpl
//
// Source: array_cursor_iterator.gen.go.tmpl

package tsm1

import (
	"context"

	"github.com/influxdata/influxdb/v2/influxql/query"
	"github.com/influxdata/influxdb/v2/models"
	"github.com/influxdata/influxdb/v2/tsdb"
)

// buildFloatArrayCursor creates an array cursor for a float field.
func (q *arrayCursorIterator) buildFloatArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) (tsdb.FloatArrayCursor, error) {
	var err error
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Float == nil {
			q.asc.Float = newFloatArrayAscendingCursor()
		}
		err = q.asc.Float.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.asc.Float, nil
	} else {
		if q.desc.Float == nil {
			q.desc.Float = newFloatArrayDescendingCursor()
		}
		err = q.desc.Float.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.desc.Float, nil
	}
}

// buildIntegerArrayCursor creates an array cursor for a integer field.
func (q *arrayCursorIterator) buildIntegerArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) (tsdb.IntegerArrayCursor, error) {
	var err error
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Integer == nil {
			q.asc.Integer = newIntegerArrayAscendingCursor()
		}
		err = q.asc.Integer.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.asc.Integer, nil
	} else {
		if q.desc.Integer == nil {
			q.desc.Integer = newIntegerArrayDescendingCursor()
		}
		err = q.desc.Integer.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.desc.Integer, nil
	}
}

// buildUnsignedArrayCursor creates an array cursor for a unsigned field.
func (q *arrayCursorIterator) buildUnsignedArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) (tsdb.UnsignedArrayCursor, error) {
	var err error
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Unsigned == nil {
			q.asc.Unsigned = newUnsignedArrayAscendingCursor()
		}
		err = q.asc.Unsigned.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.asc.Unsigned, nil
	} else {
		if q.desc.Unsigned == nil {
			q.desc.Unsigned = newUnsignedArrayDescendingCursor()
		}
		err = q.desc.Unsigned.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.desc.Unsigned, nil
	}
}

// buildStringArrayCursor creates an array cursor for a string field.
func (q *arrayCursorIterator) buildStringArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) (tsdb.StringArrayCursor, error) {
	var err error
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.String == nil {
			q.asc.String = newStringArrayAscendingCursor()
		}
		err = q.asc.String.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.asc.String, nil
	} else {
		if q.desc.String == nil {
			q.desc.String = newStringArrayDescendingCursor()
		}
		err = q.desc.String.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.desc.String, nil
	}
}

// buildBooleanArrayCursor creates an array cursor for a boolean field.
func (q *arrayCursorIterator) buildBooleanArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) (tsdb.BooleanArrayCursor, error) {
	var err error
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Boolean == nil {
			q.asc.Boolean = newBooleanArrayAscendingCursor()
		}
		err = q.asc.Boolean.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.asc.Boolean, nil
	} else {
		if q.desc.Boolean == nil {
			q.desc.Boolean = newBooleanArrayDescendingCursor()
		}
		err = q.desc.Boolean.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		if err != nil {
			return nil, err
		}
		return q.desc.Boolean, nil
	}
}
