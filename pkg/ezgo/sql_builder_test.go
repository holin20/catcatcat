package ezgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqlBuilder(t *testing.T) {
	// 	sb := NewSqlBuilder()
	// 	sql, err := sb.Select("foo", "bar", "baz").
	// 		From("table").
	// 		Where("foo > 0").
	// 		GroupBy("foo", "bar").
	// 		OrderBy("bar", "baz").
	// 		Build()
	// 	assert.NoError(t, err)
	// 	assert.Equal(t,
	// 		`SELECT foo, bar, baz
	// FROM table
	// WHERE foo > 0
	// GROUP BY foo, bar
	// ORDER BY bar, baz`, sql)

	// 	//
	// 	sb = NewSqlBuilder()
	// 	sql, err = sb.Select("foo", "bar", "baz").
	// 		From("table").
	// 		GroupBy("foo", "bar").
	// 		Build()
	// 	assert.NoError(t, err)
	// 	assert.Equal(t,
	// 		`SELECT foo, bar, baz
	// FROM table
	// GROUP BY foo, bar`, sql)

	sql := BuildInsertSql("table", map[string]*SqlCol{
		"foo": SqlColString("123"),
		"bar": SqlColInt(123),
	})
	assert.Equal(t, "INSERT INTO table (foo, bar) VALUES ('123', 123)", sql)
}
