package docstore

import (
	"errors"
	"io"
)

var (
	EntryNotFoundError = errors.New("record not found")
	ConnectionError    = errors.New("Error connecting to storage")
	IOError            = errors.New("Error reading reading from storage")
)

//Storer provides an interface for storing and retrieving "documents"(files) by unique identifier
//the document should contain a name and file extension example: /image/cat.jpeg
type Storer interface {
	Put(obj CreateObj) (string, error)        // Reads the body of the object to completion and stores it as the identifier
	Get(fileName string) (RetrieveObj, error) // Retrieves the Object by string identifier
}

//CreateObj contains an necessary information to store a document
type CreateObj struct {
	io.ReadSeeker
	Identifier string
}

//RetrieveObj contains information from retrieving a document
type RetrieveObj struct {
	io.ReadCloser
	Identifier string
}
