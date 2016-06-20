package docstore

import (
	"io"
	"os"
)

var _ Storer = File{}

// File defines a document service backed by a local file system
type File struct {
	dir string
}

// NewFileStore instantiates Storer backed by a local directory
func NewFileStore(dir string) File {

	return File{
		dir: dir,
	}
}

func (f File) Put(obj CreateObj) (string, error) {

	filePointer, err := os.Create(f.dir + obj.Identifier)
	if err != nil {
		return "", ConnectionError
	}
	defer filePointer.Close()

	_, readerr := io.Copy(filePointer, obj)
	if readerr != nil {
		return "", IOError
	}

	return obj.Identifier, nil
}

func (f File) Get(fileName string) (RetrieveObj, error) {

	filePointer, err := os.Open(f.dir + fileName)

	if err != nil {
		return RetrieveObj{}, EntryNotFoundError
	}

	doc := RetrieveObj{
		ReadCloser: filePointer,
		Identifier: fileName,
	}

	return doc, nil
}
