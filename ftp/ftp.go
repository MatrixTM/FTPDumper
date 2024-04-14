package ftp

import (
	"FTPDumper/Core"
	"github.com/jlaffaye/ftp"
	"io"
	"os"
)

type Client interface {
	Connect(address string) error
	Login() error
	Disconnect() error
	GetFiles() ([]File, error)
	DownloadFile(output, filePath string) error
	UploadFile(fileName string, reader io.Reader) error
}

type FTP struct {
	conn     *ftp.ServerConn
	Username string
	Password string
}

type File struct {
	Name string
	Size int64
}

func NewFTPClient(username, password string) Client {
	f := &FTP{
		Username: username,
		Password: password,
	}

	return f
}

// Connect establishes a connection to the FTP server at the specified address.
//
// It takes the server address as a parameter and returns an error.
func (f *FTP) Connect(address string) error {
	conn, err := ftp.Dial(address, ftp.DialWithTimeout(Core.Timeout)) // make timeout in args
	if err != nil {
		return err
	}

	f.conn = conn

	return nil
}

// Login logs the FTP client into the server.
// It takes no parameters and returns an error.
func (f *FTP) Login() error {
	err := f.conn.Login(f.Username, f.Password)
	if err != nil {
		return err
	}

	return nil
}

// Disconnect closes the FTP connection.
//
// No parameters.
// Returns an error.
func (f *FTP) Disconnect() error {
	return f.conn.Quit()
}

// GetFiles retrieves a list of files from the FTP server.
//
// No parameters. It returns a slice of File struct and an error.
func (f *FTP) GetFiles() ([]File, error) {
	// List the files in the retrieved directory
	list, err := f.conn.List(".")
	if err != nil {
		return nil, err
	}

	// Create a slice of File structs to store the file information
	files := make([]File, len(list))

	// Iterate through the list of files and retrieve file size
	for i, file := range list {
		fileSize, err := f.conn.FileSize(file.Name)
		if err != nil || fileSize < 1 {
			continue
		}
		// Populate the File struct with file name and size
		files[i] = File{
			Name: file.Name,
			Size: fileSize,
		}
	}

	// Return the list of files and no error
	return files, nil
}

// DownloadFile downloads a file from an FTP server.
//
// Parameters:
// - output: the path where the downloaded file will be saved.
// - filePath: the path of the file to be downloaded from the FTP server.
//
// Returns an error if any occurs during the download process.
func (f *FTP) DownloadFile(output, filePath string) error {
	// Retrieve the file from the FTP server
	reader, err := f.conn.Retr(filePath)
	if err != nil {
		return err
	}

	// Ensure the reader is closed after the function returns
	defer reader.Close()

	// Create the output file
	file, err := os.Create(output)
	if err != nil {
		return err
	}

	// Ensure the output file is closed after the function returns
	defer file.Close()

	// Copy data from the FTP server reader to the output file
	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}

	return nil
}

// UploadFile uploads a file to the FTP server.
func (f *FTP) UploadFile(fileName string, reader io.Reader) error {
	return f.conn.Stor(fileName, reader)
}
