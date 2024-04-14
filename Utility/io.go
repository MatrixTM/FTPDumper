package Utility

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
)

func ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines, scanner.Err()
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func FolderExists(folder string) bool {
	info, err := os.Stat(folder)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func CreateFolder(folder string) error {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return err
	}
	return nil
}

type MutexWriter struct {
	*sync.RWMutex
	writer io.Writer
}

func NewMutexWriter(writer io.Writer) *MutexWriter {
	return &MutexWriter{
		RWMutex: new(sync.RWMutex),
		writer:  writer,
	}
}

func (w *MutexWriter) Write(p []byte) (n int, err error) {
	w.RWMutex.Lock()
	defer w.RWMutex.Unlock()
	return w.writer.Write(p)
}

func (w *MutexWriter) Close() {
	w.RWMutex.Lock()
	defer w.RWMutex.Unlock()
	w.writer = nil
}

func (w *MutexWriter) GetWriter() io.Writer {
	return w.writer
}
