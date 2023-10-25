package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type File struct {
	Filename string `json:"filename"`
	Content  string `json:"file_content"`
}

type TupleParserInput struct {
	Files []File `json:"files"`
}

func main() {
	files := TupleParserInput{
		Files: []File{
			{Filename: "test1.txt", Content: "a bb ccc a"},
			{Filename: "test2.txt", Content: "a bbb cccc dd"},
			{Filename: "test3.txt", Content: "a bb ccc a"},
		},
	}

	b, err := json.MarshalIndent(files, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// put back in struct
	var tupleParserInput TupleParserInput

	if err := json.Unmarshal([]byte(b), &tupleParserInput); err != nil {
		panic(err)
	}
	for _, val := range tupleParserInput.Files {
		fmt.Println(val.Filename, val.Content)
	}

	s := "sha1 this string"
	h := sha256.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sha1_hash)
}
