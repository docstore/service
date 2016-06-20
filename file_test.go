package docstore

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestFile_Put(t *testing.T) {

	docID := "something.txt"
	docBody := "theText"
	storer := NewFileStore("")

	testBasicInterface(t, storer, docID, docBody)

	_ = os.Remove(docID)

}

func testBasicInterface(t *testing.T, storer Storer, docID, docBody string) {

	documentBodyReader := bytes.NewReader([]byte(docBody))

	obj := CreateObj{
		Identifier: docID,
		ReadSeeker: documentBodyReader,
	}

	identifier, err := storer.Put(obj)

	if err != nil {
		t.Error(err)
	}

	if identifier != obj.Identifier {
		t.Error("Identifier Is not correct")
	}

	file, openErr := storer.Get(docID)
	defer file.Close()

	if openErr != nil {
		t.Error("file was not stored")
	}

	fileBytes, readErr := ioutil.ReadAll(file)

	if readErr != nil {
		t.Error("could not read file")
	}

	if docBody != string(fileBytes) {
		t.Error("File contents not intact")
	}

}
