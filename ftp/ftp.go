package ftp

import (
	"github.com/jlaffaye/ftp"
	"os"
	"time"
)

type Client interface {
	Connect() error
	Login() error
	Disconnect() error
	GetFiles() error
	UploadFile(filePath string) ([]File, error)
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

func NewFTPClient(username, password string) *FTP {
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
	conn, err := ftp.Dial(address, ftp.DialWithTimeout(time.Second*5))
	if err != nil {
		return err
	}

	f.conn = conn

	return nil
}

// Login logs the FTP client into the server.
// It takes no parameters and returns an error.
func (f *FTP) Login() error {
	return f.conn.Login(f.Username, f.Password)
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
	//// Retrieve the current directory from the FTP server
	//dir, err := f.conn.CurrentDir()
	//if err != nil {
	//	return nil, err
	//}

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
		if err != nil {
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
// filePath specifies the path of the file to be downloaded.
// Returns an error if any occurs during the download process.
func (f *FTP) DownloadFile(filePath string) error {
	// Store the file from the FTP server
	_, err := f.conn.Retr(filePath)
	return err
}

// UploadFile uploads a file to an FTP server.
//
// filePath specifies the path of the file to be uploaded.
// Returns an error if any occurs during the upload process.
func (f *FTP) UploadFile(filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	// Ensure the file is closed after the function returns
	defer file.Close()

	// Store the file on the FTP server
	return f.conn.Stor(filePath, file)
}
