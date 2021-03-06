// Copyright 2018 Couchbase, Inc. All rights reserved.

package gojsonsm

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Makes sure that the parsing of subcontext works
func TestSimpleParserSubContext1(t *testing.T) {
	assert := assert.New(t)

	testString := "true || name.first == 'Neil'"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	assert.Equal(5, len(ctx.parserDataNodes))
	assert.Equal(5, ctx.parserTree.NumNodes())
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(1, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(3, ctx.parserTree.data[1].Right)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(1, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(4, ctx.parserTree.data[3].Right)
	assert.Equal(2, ctx.parserTree.data[3].Left)
}

func TestSimpleParserSubContext2(t *testing.T) {
	assert := assert.New(t)

	testString := "true && name.first == 'Neil' || age < 50"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(5, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(3, ctx.parserTree.data[1].Right)
	assert.Equal(1, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(2, ctx.parserTree.data[3].Left)
	assert.Equal(4, ctx.parserTree.data[3].Right)
	assert.Equal(5, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[5].Left)
	assert.Equal(7, ctx.parserTree.data[5].Right)
	assert.Equal(5, ctx.parserTree.data[7].ParentIdx)
	assert.Equal(6, ctx.parserTree.data[7].Left)
	assert.Equal(8, ctx.parserTree.data[7].Right)
}

func TestSimpleParserSubContext2a(t *testing.T) {
	assert := assert.New(t)

	testString := "(true && name.first == 'Neil') || age < 50"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(5, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(3, ctx.parserTree.data[1].Right)
	assert.Equal(1, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(2, ctx.parserTree.data[3].Left)
	assert.Equal(4, ctx.parserTree.data[3].Right)
	assert.Equal(5, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[5].Left)
	assert.Equal(7, ctx.parserTree.data[5].Right)
	assert.Equal(5, ctx.parserTree.data[7].ParentIdx)
	assert.Equal(6, ctx.parserTree.data[7].Left)
	assert.Equal(8, ctx.parserTree.data[7].Right)
}
func TestSimpleParserSubContext3(t *testing.T) {
	assert := assert.New(t)

	testString := "name.first == 'Neil' && age < 50"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(3, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
	assert.Equal(3, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[3].Left)
	assert.Equal(5, ctx.parserTree.data[3].Right)
	assert.Equal(3, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(4, ctx.parserTree.data[5].Left)
	assert.Equal(6, ctx.parserTree.data[5].Right)
}

func TestSimpleParserSubContext4(t *testing.T) {
	assert := assert.New(t)

	testString := "name.first == 'Neil' && age < 50 || isActive == true"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(3, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
	assert.Equal(7, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[3].Left)
	assert.Equal(5, ctx.parserTree.data[3].Right)
	assert.Equal(3, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(4, ctx.parserTree.data[5].Left)
	assert.Equal(6, ctx.parserTree.data[5].Right)
	assert.Equal(7, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[7].ParentIdx)
	assert.Equal(3, ctx.parserTree.data[7].Left)
	assert.Equal(9, ctx.parserTree.data[7].Right)
	assert.Equal(7, ctx.parserTree.data[9].ParentIdx)
	assert.Equal(8, ctx.parserTree.data[9].Left)
	assert.Equal(10, ctx.parserTree.data[9].Right)
}

func TestSimpleParserSubContext4a(t *testing.T) {
	assert := assert.New(t)

	// This should have short circuiting -> name.first should be checked first
	testString := "name.first == 'Neil' && age < 50 && isActive == true"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	ctx.enableShortCircuitEvalIfPossible() // NOTE this call - usually wrapped in main func
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(3, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
	assert.Equal(3, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[3].Left)
	assert.Equal(7, ctx.parserTree.data[3].Right)
	assert.Equal(7, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(4, ctx.parserTree.data[5].Left)
	assert.Equal(6, ctx.parserTree.data[5].Right)
	assert.Equal(3, ctx.parserTree.data[7].ParentIdx)
	assert.Equal(5, ctx.parserTree.data[7].Left)
	assert.Equal(9, ctx.parserTree.data[7].Right)
	assert.Equal(7, ctx.parserTree.data[9].ParentIdx)
	assert.Equal(8, ctx.parserTree.data[9].Left)
	assert.Equal(10, ctx.parserTree.data[9].Right)
}

func TestSimpleParserSubContext4b(t *testing.T) {
	assert := assert.New(t)

	// Same as 4a but no short circuit eval
	testString := "name.first == 'Neil' && age < 50 && isActive == true"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(3, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
	assert.Equal(7, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[3].Left)
	assert.Equal(5, ctx.parserTree.data[3].Right)
	assert.Equal(3, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(4, ctx.parserTree.data[5].Left)
	assert.Equal(6, ctx.parserTree.data[5].Right)
	assert.Equal(7, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[7].ParentIdx)
	assert.Equal(3, ctx.parserTree.data[7].Left)
	assert.Equal(9, ctx.parserTree.data[7].Right)
	assert.Equal(7, ctx.parserTree.data[9].ParentIdx)
	assert.Equal(8, ctx.parserTree.data[9].Left)
	assert.Equal(10, ctx.parserTree.data[9].Right)
}

func TestSimpleParserSubContext5(t *testing.T) {
	assert := assert.New(t)

	testString := "((name.first == 'Neil'))"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(1, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
}

func TestSimpleParserSubContext5a(t *testing.T) {
	assert := assert.New(t)

	testString := "( name.first == 'Neil')"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(1, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
}

func TestSimpleParserSubContext6(t *testing.T) {
	assert := assert.New(t)

	testString := "name.first == 'Neil' && (age < 50 || isActive == true)"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(3, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
	assert.Equal(3, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[3].Left)
	assert.Equal(7, ctx.parserTree.data[3].Right)
	assert.Equal(7, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(4, ctx.parserTree.data[5].Left)
	assert.Equal(6, ctx.parserTree.data[5].Right)
	assert.Equal(3, ctx.parserTree.data[7].ParentIdx)
	assert.Equal(5, ctx.parserTree.data[7].Left)
	assert.Equal(9, ctx.parserTree.data[7].Right)
	assert.Equal(7, ctx.parserTree.data[9].ParentIdx)
	assert.Equal(8, ctx.parserTree.data[9].Left)
	assert.Equal(10, ctx.parserTree.data[9].Right)
}

func TestSimpleParserSubContext7(t *testing.T) {
	assert := assert.New(t)

	testString := "(name.first == 'Neil') && (age < 50 || isActive == true)"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(3, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
	assert.Equal(3, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[3].Left)
	assert.Equal(7, ctx.parserTree.data[3].Right)
	assert.Equal(7, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(4, ctx.parserTree.data[5].Left)
	assert.Equal(6, ctx.parserTree.data[5].Right)
	assert.Equal(3, ctx.parserTree.data[7].ParentIdx)
	assert.Equal(5, ctx.parserTree.data[7].Left)
	assert.Equal(9, ctx.parserTree.data[7].Right)
	assert.Equal(7, ctx.parserTree.data[9].ParentIdx)
	assert.Equal(8, ctx.parserTree.data[9].Left)
	assert.Equal(10, ctx.parserTree.data[9].Right)
}

func TestSimpleParserSubContext7a(t *testing.T) {
	assert := assert.New(t)

	testString := "(name.first == 'Neil' )&& (age < 50 || isActive == true)"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
	node := ctx.parserDataNodes[ctx.subCtx.lastParserDataNode]
	assert.NotNil(node)

	assert.Equal(3, ctx.parserTree.data[1].ParentIdx)
	assert.Equal(0, ctx.parserTree.data[1].Left)
	assert.Equal(2, ctx.parserTree.data[1].Right)
	assert.Equal(3, ctx.treeHeadIndex)
	assert.Equal(-1, ctx.parserTree.data[3].ParentIdx)
	assert.Equal(1, ctx.parserTree.data[3].Left)
	assert.Equal(7, ctx.parserTree.data[3].Right)
	assert.Equal(7, ctx.parserTree.data[5].ParentIdx)
	assert.Equal(4, ctx.parserTree.data[5].Left)
	assert.Equal(6, ctx.parserTree.data[5].Right)
	assert.Equal(3, ctx.parserTree.data[7].ParentIdx)
	assert.Equal(5, ctx.parserTree.data[7].Left)
	assert.Equal(9, ctx.parserTree.data[7].Right)
	assert.Equal(7, ctx.parserTree.data[9].ParentIdx)
	assert.Equal(8, ctx.parserTree.data[9].Left)
	assert.Equal(10, ctx.parserTree.data[9].Right)
}

func TestContextShortCircuit1(t *testing.T) {
	assert := assert.New(t)
	testString := "name.first == 'Neil' || (age < 50) || (true)"
	ctx, _ := NewExpressionParserCtx(testString)

	ctx.enableShortCircuitEvalIfPossible()
	assert.True(ctx.shortCircuitEnabled)
}

func TestContextShortCircuit2(t *testing.T) {
	assert := assert.New(t)
	testString := "name.first == 'Neil' || (age < 50) && (true)"
	ctx, _ := NewExpressionParserCtx(testString)

	ctx.enableShortCircuitEvalIfPossible()
	assert.False(ctx.shortCircuitEnabled)
}

func TestContextParserToken(t *testing.T) {
	assert := assert.New(t)
	testString := "name.first == 'Neil' || (age < 50) || (true)"
	ctx, err := NewExpressionParserCtx(testString)

	// name.first
	_, tokenType, err := ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeField))
	assert.Nil(err)
	ctx.advanceToken()

	// ==
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeOperator))
	assert.Nil(err)
	ctx.advanceToken()

	// 'Neil'
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeValue))
	assert.Nil(err)
	ctx.advanceToken()

	// ||
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeOperator))
	assert.Nil(err)
	ctx.advanceToken()

	// (age -- will trim and will auto advance
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeParen))
	assert.Nil(err)

	// age
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeField))
	assert.Nil(err)
	//	fmt.Printf("age token: %v\n", token)
	ctx.advanceToken()

	// <
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeOperator))
	assert.Nil(err)
	ctx.advanceToken()

	// 50)
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeValue))
	assert.Nil(err)
	ctx.advanceToken()

	// )
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeEndParen))
	assert.Nil(err)
	ctx.advanceToken()

	// ||
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeOperator))
	assert.Nil(err)
	ctx.advanceToken()

	// (true -- will trim and auto advance
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeParen))
	assert.Nil(err)

	// true
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeTrue))
	assert.Nil(err)
	ctx.advanceToken()

	// )
	_, tokenType, err = ctx.getCurrentToken()
	assert.Equal(tokenType, (ParseTokenType)(TokenTypeEndParen))
	assert.Nil(err)
	ctx.advanceToken()
}

func TestSimpleParserCompare(t *testing.T) {
	assert := assert.New(t)

	testString := "something >= somethingElse"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Nil(err)
}

func TestParserExpressionOutput(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
		["equals",
			["field", "isActive"],
			["value", true]
		],
		["lessthan",
			["field", "age"],
			["value", 50]
		]
	]`)
	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	strExpr := "isActive == true || age < 50"
	ctx, err := NewExpressionParserCtx(strExpr)
	assert.Nil(err)

	err = ctx.parse()
	assert.Nil(err)

	simpleExpr, err := ctx.outputExpression()
	assert.Nil(err)

	assert.Equal(jsonExpr.String(), simpleExpr.String())
}

func TestParserExpressionOutput2a(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
	  ["equals",
	    ["field", "name", "first"],
	    ["value", "Neil"]
	  ],
	  ["and",
	    ["lessthan",
	      ["field", "age"],
	      ["value", 50]
	    ],
	    ["equals",
	      ["field", "isActive"],
	      ["value", true]
	    ]
	  ]
    ]`)

	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	strExpr := "name.first == 'Neil' || (age < 50 && isActive == true)"

	ctx, err := NewExpressionParserCtx(strExpr)
	assert.Nil(err)

	err = ctx.parse()
	assert.Nil(err)

	simpleExpr, err := ctx.outputExpression()
	assert.Nil(err)

	assert.Equal(jsonExpr.String(), simpleExpr.String())

	var trans Transformer
	matchDef := trans.Transform([]Expression{simpleExpr})
	m := NewMatcher(matchDef)

	userData := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Neil",
		},
		"isActive": true,
		"age":      32,
	}
	udMarsh, err := json.Marshal(userData)
	if err != nil {
	}
	match, err := m.Match(udMarsh)
	assert.Nil(err)
	assert.True(match)
}

func TestParserExpressionOutputNot(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
	  ["equals",
	    ["field", "name", "first"],
	    ["value", "Neal"]
	  ],
	  ["and",
	    ["not",
	      ["lessthan",
	        ["field", "age"],
	        ["value", 50]
	      ]
	    ],
	    ["equals",
	      ["field", "isActive"],
	      ["value", true]
	    ]
	  ]
    ]`)

	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	var trans Transformer
	matchDef := trans.Transform([]Expression{jsonExpr})
	assert.NotNil(matchDef)

	m := NewMatcher(matchDef)

	userData := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Neil",
		},
		"isActive": false,
		"age":      32,
	}
	udMarsh, err := json.Marshal(userData)
	if err != nil {
	}
	match, err := m.Match(udMarsh)
	assert.Nil(err)
	assert.False(match)

}

func TestParserExpressionOutputNot2(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
	  ["equals",
	    ["field", "name", "first"],
	    ["value", "Neal"]
	  ],
	  ["and",
	    ["not",
	      ["lessthan",
	        ["field", "age"],
	        ["value", 50]
	      ]
	    ],
	    ["equals",
	      ["field", "isActive"],
	      ["value", true]
	    ]
	  ]
    ]`)

	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	var trans Transformer
	matchDef := trans.Transform([]Expression{jsonExpr})
	assert.NotNil(matchDef)

	m := NewMatcher(matchDef)

	userData := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Neil",
		},
		"isActive": true,
		"age":      50,
	}
	udMarsh, err := json.Marshal(userData)
	if err != nil {
	}
	match, err := m.Match(udMarsh)
	assert.Nil(err)
	assert.True(match)
}

func TestParserExpressionOutputNot3(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
	  ["equals",
	    ["field", "name", "first"],
	    ["value", "David"]
	  ],
	  ["and",
	    ["lessthan",
	      ["field", "age"],
	      ["value", 50]
	    ],
		["notequals",
		  ["field", "isActive"],
		  ["value", true]
		]
	  ]
    ]`)

	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	var trans Transformer
	matchDef := trans.Transform([]Expression{jsonExpr})
	assert.NotNil(matchDef)

	m := NewMatcher(matchDef)

	userData := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Goliath",
		},
		"isActive": false,
		"age":      49,
	}
	udMarsh, err := json.Marshal(userData)
	if err != nil {
	}
	match, err := m.Match(udMarsh)
	assert.Nil(err)
	assert.True(match)

	strExpr := "name.first == 'David' || (age < 50 && isActive != true)"

	ctx, err := NewExpressionParserCtx(strExpr)
	assert.Nil(err)

	err = ctx.parse()
	assert.Nil(err)

	simpleExpr, err := ctx.outputExpression()
	assert.Nil(err)

	assert.Equal(jsonExpr.String(), simpleExpr.String())
}

func TestParserExpressionOutputGreaterThan(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
	  ["equals",
	    ["field", "name", "first"],
	    ["value", "David"]
	  ],
	  ["and",
	    ["greaterthan",
	      ["field", "age"],
	      ["value", 50]
	    ],
	    ["equals",
	      ["field", "isActive"],
	      ["value", true]
	    ]
	  ]
    ]`)

	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	var trans Transformer
	matchDef := trans.Transform([]Expression{jsonExpr})
	assert.NotNil(matchDef)

	m := NewMatcher(matchDef)

	userData := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Goliath",
		},
		"isActive": true,
		"age":      51,
	}
	udMarsh, err := json.Marshal(userData)
	if err != nil {
	}
	match, err := m.Match(udMarsh)
	assert.Nil(err)
	assert.True(match)

	strExpr := "name.first == 'David' || (age > 50 && isActive == true)"

	ctx, err := NewExpressionParserCtx(strExpr)
	assert.Nil(err)

	err = ctx.parse()
	assert.Nil(err)

	simpleExpr, err := ctx.outputExpression()
	assert.Nil(err)

	assert.Equal(jsonExpr.String(), simpleExpr.String())
}

func TestParserExpressionOutputGreaterThanEquals(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
	  ["equals",
	    ["field", "name", "first"],
	    ["value", "David"]
	  ],
	  ["and",
	    ["greaterequals",
	      ["field", "age"],
	      ["value", 50]
	    ],
	    ["equals",
	      ["field", "isActive"],
	      ["value", true]
	    ]
	  ]
    ]`)

	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	var trans Transformer
	matchDef := trans.Transform([]Expression{jsonExpr})
	assert.NotNil(matchDef)

	m := NewMatcher(matchDef)

	userData := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Goliath",
		},
		"isActive": true,
		"age":      50,
	}
	udMarsh, err := json.Marshal(userData)
	if err != nil {
	}
	match, err := m.Match(udMarsh)
	assert.Nil(err)
	assert.True(match)

	strExpr := "name.first == 'David' || (age >= 50 && isActive == true)"

	ctx, err := NewExpressionParserCtx(strExpr)
	assert.Nil(err)

	err = ctx.parse()
	assert.Nil(err)

	simpleExpr, err := ctx.outputExpression()
	assert.Nil(err)

	assert.Equal(jsonExpr.String(), simpleExpr.String())
}

func TestParserExpressionOutputLessThan(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
	  ["equals",
	    ["field", "name", "first"],
	    ["value", "David"]
	  ],
	  ["and",
	    ["lessthan",
	      ["field", "age"],
	      ["value", 50]
	    ],
	    ["equals",
	      ["field", "isActive"],
	      ["value", true]
	    ]
	  ]
    ]`)

	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	var trans Transformer
	matchDef := trans.Transform([]Expression{jsonExpr})
	assert.NotNil(matchDef)

	m := NewMatcher(matchDef)

	userData := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Goliath",
		},
		"isActive": true,
		"age":      49,
	}
	udMarsh, err := json.Marshal(userData)
	if err != nil {
	}
	match, err := m.Match(udMarsh)
	assert.Nil(err)
	assert.True(match)

	strExpr := "name.first == 'David' || (age < 50 && isActive == true)"

	ctx, err := NewExpressionParserCtx(strExpr)
	assert.Nil(err)

	err = ctx.parse()
	assert.Nil(err)

	simpleExpr, err := ctx.outputExpression()
	assert.Nil(err)

	assert.Equal(jsonExpr.String(), simpleExpr.String())
}

func TestParserExpressionOutputLessThanEq(t *testing.T) {
	assert := assert.New(t)

	matchJson := []byte(`
	["or",
	  ["equals",
	    ["field", "name", "first"],
	    ["value", "David"]
	  ],
	  ["and",
	    ["lessequals",
	      ["field", "age"],
	      ["value", 50]
	    ],
	    ["equals",
	      ["field", "isActive"],
	      ["value", true]
	    ]
	  ]
    ]`)

	jsonExpr, err := ParseJsonExpression(matchJson)
	assert.Nil(err)

	var trans Transformer
	matchDef := trans.Transform([]Expression{jsonExpr})
	assert.NotNil(matchDef)

	m := NewMatcher(matchDef)

	userData := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Goliath",
		},
		"isActive": true,
		"age":      50,
	}
	udMarsh, err := json.Marshal(userData)
	if err != nil {
	}
	match, err := m.Match(udMarsh)
	assert.Nil(err)
	assert.True(match)

	strExpr := "name.first == 'David' || (age <= 50 && isActive == true)"

	ctx, err := NewExpressionParserCtx(strExpr)
	assert.Nil(err)

	err = ctx.parse()
	assert.Nil(err)

	simpleExpr, err := ctx.outputExpression()
	assert.Nil(err)

	assert.Equal(jsonExpr.String(), simpleExpr.String())
}

// NEGATIVE test cases
func TestSimpleParserParenMismatch(t *testing.T) {
	assert := assert.New(t)

	testString := "(name.first == 'Neil'))"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Equal(ErrorParenMismatch, err)
}

func TestSimpleParserParenMismatch2(t *testing.T) {
	assert := assert.New(t)

	testString := "((name.first == 'Neil')"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Equal(ErrorParenMismatch, err)
}

func TestSimpleParserParenMismatch3(t *testing.T) {
	assert := assert.New(t)

	testString := ")>= 3"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Equal(ErrorParenMismatch, err)
}

func TestSimpleParserParenSyntaxErr(t *testing.T) {
	assert := assert.New(t)

	testString := "(aField)> 3"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserParenSyntaxErr2(t *testing.T) {
	assert := assert.New(t)

	testString := "(someField == true)&& true"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

// Test for when the first token is NOT a field value
func TestSimpleParserNeg(t *testing.T) {
	assert := assert.New(t)

	testString := "|| true"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserNeg2(t *testing.T) {
	assert := assert.New(t)

	testString := "age < Neil == true"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserNeg3(t *testing.T) {
	assert := assert.New(t)

	testString := "something >= true"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserNeg4(t *testing.T) {
	assert := assert.New(t)

	testString := ">= 2"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserNeg5(t *testing.T) {
	assert := assert.New(t)

	testString := "( true)&&( false)"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserNeg6(t *testing.T) {
	assert := assert.New(t)

	testString := "'Neil' == name.first && 50 > age"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserNeg7(t *testing.T) {
	assert := assert.New(t)

	testString := "abc(def"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.Equal(ErrorParenWSpace, err)
}

func TestSimpleParserNeg8(t *testing.T) {
	assert := assert.New(t)

	testString := "someField == true &&(def) == false"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserNeg9(t *testing.T) {
	assert := assert.New(t)

	testString := ".somefield == 1"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestSimpleParserNeg10(t *testing.T) {
	assert := assert.New(t)

	testString := "somefield. == 1"
	ctx, err := NewExpressionParserCtx(testString)
	assert.Equal(fieldMode, ctx.subCtx.currentMode)
	err = ctx.parse()
	assert.NotNil(err)
}

func TestParserExpressionOutputNeg(t *testing.T) {
	assert := assert.New(t)

	emptyString := ""
	ctx, err := NewExpressionParserCtx(emptyString)
	assert.Nil(err)

	_, err = ctx.outputExpression()
	assert.NotNil(err)
}

func TestParserExpressionWithGreaterThan(t *testing.T) {
	assert := assert.New(t)

	strExpr := "age > 50"

	ctx, err := NewExpressionParserCtx(strExpr)
	assert.Nil(err)

	err = ctx.parse()
	assert.Nil(err)

	_, err = ctx.outputExpression()
	assert.Nil(err)
}
