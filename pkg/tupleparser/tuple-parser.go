package tupleparser

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
)

type File struct {
	Filename string `json:"filename"`
	Content  string `json:"file_content"`
}

type TupleParserInput struct {
	Files []File
}

type TupleParserOutputItem struct {
	Sha256                string `json:"sha256"`
	LongestFilenameLength int    `json:"longest_filename_length"`
	LongestWordLength     int    `json:"longest_word_length"`
	NumFiles              int    `json:"num_files"`
	NumUniqueWords        int    `json:"num_unique_words"`
	NumWords              int    `json:"num_words"`
}

type TupleParserOutput struct {
	Errors                 int
	TupleParserOutputItems []TupleParserOutputItem
}

type TupleParser struct {
	InputJsonStr      string
	OutputJsonStr     string
	outputMap         map[string]TupleParserOutputItem
	TupleParserInput  *TupleParserInput
	TupleParserOutput *TupleParserOutput
}

func NewTupleParser() *TupleParser {
	// initialize the map
	outputMap := map[string]TupleParserOutputItem{}
	tupleParserInput := TupleParserInput{}
	tupleParserOutput := TupleParserOutput{}

	return &TupleParser{
		outputMap: outputMap,
		TupleParserInput: &tupleParserInput,
		TupleParserOutput: &tupleParserOutput,
	}
}

func (tp *TupleParser) parseInputString(in string) error {
	var tupleParserInput TupleParserInput

	if err := json.Unmarshal([]byte(in), &tupleParserInput); err != nil {
		fmt.Println("Error parsing input json: ", err.Error())
		return err
	}

	for _, val := range tupleParserInput.Files {
		fmt.Println(val.Filename, val.Content)
	}

	tp.TupleParserInput = &tupleParserInput

	return nil
}

func (tp *TupleParser) processFile(file *File) error {

	if len(file.Content) == 0 || len(file.Filename) == 0 {
		tp.TupleParserOutput.Errors += 1
		return errors.New("filename are content was bad")
	}

	// check map, if not in calculate sha256 and add it to the map
	h := sha256.New()
	h.Write([]byte(file.Content))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	// lookup the output struct
	outputItem, ok := tp.outputMap[sha1_hash]

	// if we have not seen the content hash
	if !ok {
		outputItem := &TupleParserOutputItem{
			Sha256:                sha1_hash,
			LongestFilenameLength: len(file.Filename),
			LongestWordLength:     3,
			NumFiles:              1,
			NumUniqueWords:        3,
			NumWords:              4,
		}
		tp.outputMap[sha1_hash] = *outputItem
		fmt.Println("adding sha256: ", outputItem.Sha256)
	} else {
		outputItem.NumFiles += 1
		tp.outputMap[sha1_hash] = outputItem
		fmt.Println("updating sha256:", outputItem.Sha256)
	}

	return nil
}

func (tp *TupleParser) process() error {

	fmt.Println("in process...")
	// for each file process for errors and accumulate statistic
	for _, file := range tp.TupleParserInput.Files {
		err := tp.processFile(&file)
		if err != nil {
			fmt.Println("Error processing file: ", err.Error())
		}
	}

	for key , val := range tp.outputMap {
		fmt.Println(key, val.NumFiles)
		tp.TupleParserOutput.TupleParserOutputItems = append(tp.TupleParserOutput.TupleParserOutputItems, val)
	}

	return nil
}

func (tp *TupleParser) Parse(in string) error {
	fmt.Println("parsing :", in)

	tp.InputJsonStr = in
	err := tp.parseInputString(in)
	if err != nil {
		fmt.Println("Error parsing input json: ", err.Error())
		return err
	}

	// process the files
	if err := tp.process(); err != nil {
		fmt.Println("Error parsing input json: ", err.Error())
	}

	tp.InputJsonStr = in

	return nil
}

func (tp *TupleParser) PrintOutput() error {
	
	b, err := json.MarshalIndent(tp.TupleParserOutput, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))


	return nil
}
