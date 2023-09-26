package main

import (
	"fmt"

	"github.com/mccuskero/go-parse-app/pkg/tupleparser"
)

func main() {
	fmt.Println("Tuple parser starting... ")

	inputJson := `{
		"files": [
		  {
			"filename": "",
			"file_content": "a bb ccc dddd"
		  },
		  {
			"filename": "a.txt",
			"file_content": ""
		  },
		  {
			"filename": "bb.txt",
			"file_content": "a bb ccc dddd"
		  },
		  {
			"filename": "ccc.txt",
			"file_content": "e ff e ff e ff"
		  },
		  {
			"filename": "dddd.txt",
			"file_content": ""
		  },
		  {
			"filename": "eeeee.txt",
			"file_content": "g g g g g"
		  },
		  {
			"filename": "ffffff.txt",
			"file_content": "a bb ccc dddd"
		  },
		  {
			"filename": "ggggggg.txt",
			"file_content": "a bb ccc dddd"
		  },
		  {
			"filename": "hhhhhhhh.txt",
			"file_content": "a bb ccc dddd"
		  },
		  {
			"filename": "iiiiiiiii.txt",
			"file_content": ""
		  },
		  {
			"filename": "jjjjjjjjjj.txt",
			"file_content": ""
		  },
		  {
			"filename": "kkkkkkkkkkk.txt",
			"file_content": "g g g g g"
		  },
		  {
			"filename": "ccc.txt",
			"file_content": "a bb ccc dddd"
		  }
		]
	}`

	tp := tupleparser.NewTupleParser()
	tp.Parse(inputJson)
	tp.PrintOutput()
}
