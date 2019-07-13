package types

import (
	"io"
	"os"
)

// WebDav available methods
type WebDav interface {
	// ReadDir reads the contents of a remote directory
	ReadDir(path string) ([]os.FileInfo, error)
	// Stat returns the file stats for a specified path
	Stat(path string) (os.FileInfo, error)
	// Remove removes a remote file
	Remove(path string) error
	// RemoveAll removes remote files
	RemoveAll(path string) error
	// Mkdir makes a directory
	Mkdir(path string, _ os.FileMode) error
	// MkdirAll like mkdir -p, but for webdav
	MkdirAll(path string, _ os.FileMode) error
	// Rename moves a file from A to B
	Rename(oldpath, newpath string, overwrite bool) error
	// Copy copies a file from A to B
	Copy(oldpath, newpath string, overwrite bool) error
	// Read reads the contents of a remote file
	Read(path string) ([]byte, error)
	// ReadStream reads the stream for a given path
	ReadStream(path string) (io.ReadCloser, error)
	// Write writes data to a given path
	Write(path string, data []byte, _ os.FileMode) error
	// WriteStream writes a stream
	WriteStream(path string, stream io.Reader, _ os.FileMode) error
}
