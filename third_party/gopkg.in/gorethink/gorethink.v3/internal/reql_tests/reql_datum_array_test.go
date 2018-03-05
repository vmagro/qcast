// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/gorethink/gorethink.v3"
	"gopkg.in/gorethink/gorethink.v3/internal/compare"
)

// Tests conversion to and from the RQL array type
func TestDatumArraySuite(t *testing.T) {
	suite.Run(t, new(DatumArraySuite))
}

type DatumArraySuite struct {
	suite.Suite

	session *r.Session
}

func (suite *DatumArraySuite) SetupTest() {
	suite.T().Log("Setting up DatumArraySuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *DatumArraySuite) TearDownSuite() {
	suite.T().Log("Tearing down DatumArraySuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *DatumArraySuite) TestCases() {
	suite.T().Log("Running DatumArraySuite: Tests conversion to and from the RQL array type")

	{
		// datum/array.yaml line #6
		/* [] */
		var expected_ []interface{} = []interface{}{}
		/* r.expr([]) */

		suite.T().Log("About to run line #6: r.Expr([]interface{}{})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// datum/array.yaml line #9
		/* [1] */
		var expected_ []interface{} = []interface{}{1}
		/* r.expr([1]) */

		suite.T().Log("About to run line #9: r.Expr([]interface{}{1})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #9")
	}

	{
		// datum/array.yaml line #14
		/* [1,2,3,4,5] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4, 5}
		/* r.expr([1,2,3,4,5]) */

		suite.T().Log("About to run line #14: r.Expr([]interface{}{1, 2, 3, 4, 5})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4, 5}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #14")
	}

	{
		// datum/array.yaml line #19
		/* 'ARRAY' */
		var expected_ string = "ARRAY"
		/* r.expr([]).type_of() */

		suite.T().Log("About to run line #19: r.Expr([]interface{}{}).TypeOf()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{}).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #19")
	}

	{
		// datum/array.yaml line #24
		/* '[1,2]' */
		var expected_ string = "[1,2]"
		/* r.expr([1, 2]).coerce_to('string') */

		suite.T().Log("About to run line #24: r.Expr([]interface{}{1, 2}).CoerceTo('string')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).CoerceTo("string"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #24")
	}

	{
		// datum/array.yaml line #25
		/* '[1,2]' */
		var expected_ string = "[1,2]"
		/* r.expr([1, 2]).coerce_to('STRING') */

		suite.T().Log("About to run line #25: r.Expr([]interface{}{1, 2}).CoerceTo('STRING')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).CoerceTo("STRING"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #25")
	}

	{
		// datum/array.yaml line #28
		/* [1, 2] */
		var expected_ []interface{} = []interface{}{1, 2}
		/* r.expr([1, 2]).coerce_to('array') */

		suite.T().Log("About to run line #28: r.Expr([]interface{}{1, 2}).CoerceTo('array')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).CoerceTo("array"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// datum/array.yaml line #31
		/* err('ReqlQueryLogicError', 'Cannot coerce ARRAY to NUMBER.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot coerce ARRAY to NUMBER.")
		/* r.expr([1, 2]).coerce_to('number') */

		suite.T().Log("About to run line #31: r.Expr([]interface{}{1, 2}).CoerceTo('number')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).CoerceTo("number"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #31")
	}

	{
		// datum/array.yaml line #34
		/* {'a':1,'b':2} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2}
		/* r.expr([['a', 1], ['b', 2]]).coerce_to('object') */

		suite.T().Log("About to run line #34: r.Expr([]interface{}{[]interface{}{'a', 1}, []interface{}{'b', 2}}).CoerceTo('object')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{[]interface{}{"a", 1}, []interface{}{"b", 2}}).CoerceTo("object"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #34")
	}

	{
		// datum/array.yaml line #37
		/* err('ReqlQueryLogicError', 'Expected array of size 2, but got size 0.') */
		var expected_ Err = err("ReqlQueryLogicError", "Expected array of size 2, but got size 0.")
		/* r.expr([[]]).coerce_to('object') */

		suite.T().Log("About to run line #37: r.Expr([]interface{}{[]interface{}{}}).CoerceTo('object')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{[]interface{}{}}).CoerceTo("object"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// datum/array.yaml line #40
		/* err('ReqlQueryLogicError', 'Expected array of size 2, but got size 3.') */
		var expected_ Err = err("ReqlQueryLogicError", "Expected array of size 2, but got size 3.")
		/* r.expr([['1',2,3]]).coerce_to('object') */

		suite.T().Log("About to run line #40: r.Expr([]interface{}{[]interface{}{'1', 2, 3}}).CoerceTo('object')")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{[]interface{}{"1", 2, 3}}).CoerceTo("object"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #40")
	}

	{
		// datum/array.yaml line #44
		/* [1] */
		var expected_ []interface{} = []interface{}{1}
		/* r.expr([r.expr(1)]) */

		suite.T().Log("About to run line #44: r.Expr([]interface{}{r.Expr(1)})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{r.Expr(1)}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #44")
	}

	{
		// datum/array.yaml line #47
		/* [1,2,3,4] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4}
		/* r.expr([1,3,4]).insert_at(1, 2) */

		suite.T().Log("About to run line #47: r.Expr([]interface{}{1, 3, 4}).InsertAt(1, 2)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 3, 4}).InsertAt(1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #47")
	}

	{
		// datum/array.yaml line #49
		/* [1,2,3] */
		var expected_ []interface{} = []interface{}{1, 2, 3}
		/* r.expr([2,3]).insert_at(0, 1) */

		suite.T().Log("About to run line #49: r.Expr([]interface{}{2, 3}).InsertAt(0, 1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{2, 3}).InsertAt(0, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #49")
	}

	{
		// datum/array.yaml line #51
		/* [1,2,3,4] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4}
		/* r.expr([1,2,3]).insert_at(-1, 4) */

		suite.T().Log("About to run line #51: r.Expr([]interface{}{1, 2, 3}).InsertAt(-1, 4)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).InsertAt(-1, 4), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #51")
	}

	{
		// datum/array.yaml line #53
		/* [1,2,3,4] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4}
		/* r.expr([1,2,3]).insert_at(3, 4) */

		suite.T().Log("About to run line #53: r.Expr([]interface{}{1, 2, 3}).InsertAt(3, 4)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).InsertAt(3, 4), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #53")
	}

	{
		// datum/array.yaml line #55
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* r.expr(3).do(lambda x: r.expr([1,2,3]).insert_at(x, 4)) */

		suite.T().Log("About to run line #55: r.Expr(3).Do(func(x r.Term) interface{} { return r.Expr([]interface{}{1, 2, 3}).InsertAt(x, 4)})")

		runAndAssert(suite.Suite, expected_, r.Expr(3).Do(func(x r.Term) interface{} { return r.Expr([]interface{}{1, 2, 3}).InsertAt(x, 4) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #55")
	}

	{
		// datum/array.yaml line #59
		/* err('ReqlNonExistenceError', 'Index `4` out of bounds for array of size: `3`.', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Index `4` out of bounds for array of size: `3`.")
		/* r.expr([1,2,3]).insert_at(4, 5) */

		suite.T().Log("About to run line #59: r.Expr([]interface{}{1, 2, 3}).InsertAt(4, 5)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).InsertAt(4, 5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #59")
	}

	{
		// datum/array.yaml line #61
		/* err('ReqlNonExistenceError', 'Index out of bounds: -5', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Index out of bounds: -5")
		/* r.expr([1,2,3]).insert_at(-5, -1) */

		suite.T().Log("About to run line #61: r.Expr([]interface{}{1, 2, 3}).InsertAt(-5, -1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).InsertAt(-5, -1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #61")
	}

	{
		// datum/array.yaml line #63
		/* err('ReqlQueryLogicError', 'Number not an integer: 1.5', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Number not an integer: 1.5")
		/* r.expr([1,2,3]).insert_at(1.5, 1) */

		suite.T().Log("About to run line #63: r.Expr([]interface{}{1, 2, 3}).InsertAt(1.5, 1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).InsertAt(1.5, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #63")
	}

	{
		// datum/array.yaml line #65
		/* err('ReqlNonExistenceError', 'Expected type NUMBER but found NULL.', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Expected type NUMBER but found NULL.")
		/* r.expr([1,2,3]).insert_at(null, 1) */

		suite.T().Log("About to run line #65: r.Expr([]interface{}{1, 2, 3}).InsertAt(nil, 1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).InsertAt(nil, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #65")
	}

	{
		// datum/array.yaml line #68
		/* [1,2,3,4] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4}
		/* r.expr([1,4]).splice_at(1, [2,3]) */

		suite.T().Log("About to run line #68: r.Expr([]interface{}{1, 4}).SpliceAt(1, []interface{}{2, 3})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 4}).SpliceAt(1, []interface{}{2, 3}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #68")
	}

	{
		// datum/array.yaml line #70
		/* [1,2,3,4] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4}
		/* r.expr([3,4]).splice_at(0, [1,2]) */

		suite.T().Log("About to run line #70: r.Expr([]interface{}{3, 4}).SpliceAt(0, []interface{}{1, 2})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{3, 4}).SpliceAt(0, []interface{}{1, 2}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #70")
	}

	{
		// datum/array.yaml line #72
		/* [1,2,3,4] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4}
		/* r.expr([1,2]).splice_at(2, [3,4]) */

		suite.T().Log("About to run line #72: r.Expr([]interface{}{1, 2}).SpliceAt(2, []interface{}{3, 4})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).SpliceAt(2, []interface{}{3, 4}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #72")
	}

	{
		// datum/array.yaml line #74
		/* [1,2,3,4] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4}
		/* r.expr([1,2]).splice_at(-1, [3,4]) */

		suite.T().Log("About to run line #74: r.Expr([]interface{}{1, 2}).SpliceAt(-1, []interface{}{3, 4})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).SpliceAt(-1, []interface{}{3, 4}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #74")
	}

	{
		// datum/array.yaml line #76
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* r.expr(2).do(lambda x: r.expr([1,2]).splice_at(x, [3,4])) */

		suite.T().Log("About to run line #76: r.Expr(2).Do(func(x r.Term) interface{} { return r.Expr([]interface{}{1, 2}).SpliceAt(x, []interface{}{3, 4})})")

		runAndAssert(suite.Suite, expected_, r.Expr(2).Do(func(x r.Term) interface{} { return r.Expr([]interface{}{1, 2}).SpliceAt(x, []interface{}{3, 4}) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #76")
	}

	{
		// datum/array.yaml line #80
		/* err('ReqlNonExistenceError', 'Index `3` out of bounds for array of size: `2`.', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Index `3` out of bounds for array of size: `2`.")
		/* r.expr([1,2]).splice_at(3, [3,4]) */

		suite.T().Log("About to run line #80: r.Expr([]interface{}{1, 2}).SpliceAt(3, []interface{}{3, 4})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).SpliceAt(3, []interface{}{3, 4}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #80")
	}

	{
		// datum/array.yaml line #82
		/* err('ReqlNonExistenceError', 'Index out of bounds: -4', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Index out of bounds: -4")
		/* r.expr([1,2]).splice_at(-4, [3,4]) */

		suite.T().Log("About to run line #82: r.Expr([]interface{}{1, 2}).SpliceAt(-4, []interface{}{3, 4})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).SpliceAt(-4, []interface{}{3, 4}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #82")
	}

	{
		// datum/array.yaml line #84
		/* err('ReqlQueryLogicError', 'Number not an integer: 1.5', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Number not an integer: 1.5")
		/* r.expr([1,2,3]).splice_at(1.5, [1]) */

		suite.T().Log("About to run line #84: r.Expr([]interface{}{1, 2, 3}).SpliceAt(1.5, []interface{}{1})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).SpliceAt(1.5, []interface{}{1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #84")
	}

	{
		// datum/array.yaml line #86
		/* err('ReqlNonExistenceError', 'Expected type NUMBER but found NULL.', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Expected type NUMBER but found NULL.")
		/* r.expr([1,2,3]).splice_at(null, [1]) */

		suite.T().Log("About to run line #86: r.Expr([]interface{}{1, 2, 3}).SpliceAt(nil, []interface{}{1})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).SpliceAt(nil, []interface{}{1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #86")
	}

	{
		// datum/array.yaml line #88
		/* err('ReqlQueryLogicError', 'Expected type ARRAY but found NUMBER.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type ARRAY but found NUMBER.")
		/* r.expr([1,4]).splice_at(1, 2) */

		suite.T().Log("About to run line #88: r.Expr([]interface{}{1, 4}).SpliceAt(1, 2)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 4}).SpliceAt(1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #88")
	}

	{
		// datum/array.yaml line #91
		/* [2,3,4] */
		var expected_ []interface{} = []interface{}{2, 3, 4}
		/* r.expr([1,2,3,4]).delete_at(0) */

		suite.T().Log("About to run line #91: r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(0)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #91")
	}

	{
		// datum/array.yaml line #93
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* r.expr(0).do(lambda x: r.expr([1,2,3,4]).delete_at(x)) */

		suite.T().Log("About to run line #93: r.Expr(0).Do(func(x r.Term) interface{} { return r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(x)})")

		runAndAssert(suite.Suite, expected_, r.Expr(0).Do(func(x r.Term) interface{} { return r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(x) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #93")
	}

	{
		// datum/array.yaml line #97
		/* [1,2,3] */
		var expected_ []interface{} = []interface{}{1, 2, 3}
		/* r.expr([1,2,3,4]).delete_at(-1) */

		suite.T().Log("About to run line #97: r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(-1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(-1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #97")
	}

	{
		// datum/array.yaml line #99
		/* [1,4] */
		var expected_ []interface{} = []interface{}{1, 4}
		/* r.expr([1,2,3,4]).delete_at(1,3) */

		suite.T().Log("About to run line #99: r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(1, 3)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(1, 3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #99")
	}

	{
		// datum/array.yaml line #101
		/* [1,2,3,4] */
		var expected_ []interface{} = []interface{}{1, 2, 3, 4}
		/* r.expr([1,2,3,4]).delete_at(4,4) */

		suite.T().Log("About to run line #101: r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(4, 4)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(4, 4), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #101")
	}

	{
		// datum/array.yaml line #103
		/* [] */
		var expected_ []interface{} = []interface{}{}
		/* r.expr([]).delete_at(0,0) */

		suite.T().Log("About to run line #103: r.Expr([]interface{}{}).DeleteAt(0, 0)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{}).DeleteAt(0, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #103")
	}

	{
		// datum/array.yaml line #105
		/* [1,4] */
		var expected_ []interface{} = []interface{}{1, 4}
		/* r.expr([1,2,3,4]).delete_at(1,-1) */

		suite.T().Log("About to run line #105: r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(1, -1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(1, -1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #105")
	}

	{
		// datum/array.yaml line #107
		/* err('ReqlNonExistenceError', 'Index `4` out of bounds for array of size: `4`.', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Index `4` out of bounds for array of size: `4`.")
		/* r.expr([1,2,3,4]).delete_at(4) */

		suite.T().Log("About to run line #107: r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(4)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(4), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #107")
	}

	{
		// datum/array.yaml line #109
		/* err('ReqlNonExistenceError', 'Index out of bounds: -5', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Index out of bounds: -5")
		/* r.expr([1,2,3,4]).delete_at(-5) */

		suite.T().Log("About to run line #109: r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(-5)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4}).DeleteAt(-5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #109")
	}

	{
		// datum/array.yaml line #111
		/* err('ReqlQueryLogicError', 'Number not an integer: 1.5', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Number not an integer: 1.5")
		/* r.expr([1,2,3]).delete_at(1.5) */

		suite.T().Log("About to run line #111: r.Expr([]interface{}{1, 2, 3}).DeleteAt(1.5)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).DeleteAt(1.5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #111")
	}

	{
		// datum/array.yaml line #113
		/* err('ReqlNonExistenceError', 'Expected type NUMBER but found NULL.', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Expected type NUMBER but found NULL.")
		/* r.expr([1,2,3]).delete_at(null) */

		suite.T().Log("About to run line #113: r.Expr([]interface{}{1, 2, 3}).DeleteAt(nil)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).DeleteAt(nil), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #113")
	}

	{
		// datum/array.yaml line #116
		/* [1,2,3] */
		var expected_ []interface{} = []interface{}{1, 2, 3}
		/* r.expr([0,2,3]).change_at(0, 1) */

		suite.T().Log("About to run line #116: r.Expr([]interface{}{0, 2, 3}).ChangeAt(0, 1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{0, 2, 3}).ChangeAt(0, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #116")
	}

	{
		// datum/array.yaml line #118
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* r.expr(1).do(lambda x: r.expr([0,2,3]).change_at(0,x)) */

		suite.T().Log("About to run line #118: r.Expr(1).Do(func(x r.Term) interface{} { return r.Expr([]interface{}{0, 2, 3}).ChangeAt(0, x)})")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Do(func(x r.Term) interface{} { return r.Expr([]interface{}{0, 2, 3}).ChangeAt(0, x) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #118")
	}

	{
		// datum/array.yaml line #122
		/* [1,2,3] */
		var expected_ []interface{} = []interface{}{1, 2, 3}
		/* r.expr([1,0,3]).change_at(1, 2) */

		suite.T().Log("About to run line #122: r.Expr([]interface{}{1, 0, 3}).ChangeAt(1, 2)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 0, 3}).ChangeAt(1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #122")
	}

	{
		// datum/array.yaml line #124
		/* [1,2,3] */
		var expected_ []interface{} = []interface{}{1, 2, 3}
		/* r.expr([1,2,0]).change_at(2, 3) */

		suite.T().Log("About to run line #124: r.Expr([]interface{}{1, 2, 0}).ChangeAt(2, 3)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 0}).ChangeAt(2, 3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #124")
	}

	{
		// datum/array.yaml line #126
		/* err('ReqlNonExistenceError', 'Index `3` out of bounds for array of size: `3`.', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Index `3` out of bounds for array of size: `3`.")
		/* r.expr([1,2,3]).change_at(3, 4) */

		suite.T().Log("About to run line #126: r.Expr([]interface{}{1, 2, 3}).ChangeAt(3, 4)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ChangeAt(3, 4), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #126")
	}

	{
		// datum/array.yaml line #128
		/* err('ReqlNonExistenceError', 'Index out of bounds: -5', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Index out of bounds: -5")
		/* r.expr([1,2,3,4]).change_at(-5, 1) */

		suite.T().Log("About to run line #128: r.Expr([]interface{}{1, 2, 3, 4}).ChangeAt(-5, 1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4}).ChangeAt(-5, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #128")
	}

	{
		// datum/array.yaml line #130
		/* err('ReqlQueryLogicError', 'Number not an integer: 1.5', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Number not an integer: 1.5")
		/* r.expr([1,2,3]).change_at(1.5, 1) */

		suite.T().Log("About to run line #130: r.Expr([]interface{}{1, 2, 3}).ChangeAt(1.5, 1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ChangeAt(1.5, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #130")
	}

	{
		// datum/array.yaml line #132
		/* err('ReqlNonExistenceError', 'Expected type NUMBER but found NULL.', [0]) */
		var expected_ Err = err("ReqlNonExistenceError", "Expected type NUMBER but found NULL.")
		/* r.expr([1,2,3]).change_at(null, 1) */

		suite.T().Log("About to run line #132: r.Expr([]interface{}{1, 2, 3}).ChangeAt(nil, 1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ChangeAt(nil, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #132")
	}
}
