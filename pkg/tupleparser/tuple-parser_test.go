package tupleparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleParserForErrors(t *testing.T) {

	inputJson := `{
		"files": [
		  {
			"filename": "a.txt",
			"file_content": "a bb ccc a"
		  },
		  {
			"filename": "bb.txt",
			"file_content": "a bbb cccc dd"
		  },
		  {
			"filename": "ccc.txt",
			"file_content": "a bb ccc a"
		  },
		  {
			"filename": "dddd.txt",
			"file_content": "a bb ccc a"
		  },
		  {
			"filename": "",
			"file_content": "this is an error"
		  },
		  {
			"filename": "error.txt",
			"file_content": ""
		  }
		]
	}`

	tp := NewTupleParser()
	tp.Parse(inputJson)
	tp.PrintOutput()

	assert.Equal(t, 2, tp.TupleParserOutput.Errors)
}
