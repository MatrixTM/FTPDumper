package Core

import (
	"FTPDumper/Utility"
	"errors"
	"fmt"
	"github.com/integrii/flaggy"
	"os"
	"strings"
	"time"
	"unicode"
)

var (
	Scanner         = "stdin"
	Users           []string
	Passwords       []string
	Ports           []string
	FileFormats     []string
	Limit           = 10
	SaveCredentials = false
	Verbose         = false
	Timeout         = time.Second * 5
	CredFile        *Utility.MutexWriter
	OutputFolder    = "files"
	Type            EscannerType
	Counter         = Utility.NewCounter()
)

var (
	TimeoutErr = errors.New("timeout Error")
	BadCredErr = errors.New("bad Credentials")
)

func init() {
	Counter.Born("Total")
	Counter.Born("BadCred")
	Counter.Born("Success")
	Counter.Born("Stolen")

	flaggy.SetName("FTPDumper")
	flaggy.SetDescription("Scan World FTP Servers and Steal Their Data")
	flaggy.SetVersion("0.0.1")

	flaggy.String(&Scanner, "scan", "scanner", "Ip/CIDR scanner (stdin|filename|cidr|ip)")

	comboFile := ""
	flaggy.String(&comboFile, "c", "combo", "Combo File (user:password)")

	ports := ""
	flaggy.String(&ports, "p", "ports", "Ports Split by , (Default Port: 21)")

	flaggy.String(&OutputFolder, "o", "output", "Output Folder")

	fileFormats := ""
	flaggy.String(&fileFormats, "f", "formats", "File Formats Split by , (Default Format: all)")

	flaggy.Int(&Limit, "l", "limit", "Task limit")

	flaggy.Duration(&Timeout, "t", "timeout", "Timeout in seconds")

	flaggy.Bool(&SaveCredentials, "s", "save", "Save Credentials in hit.txt")
	flaggy.Bool(&Verbose, "v", "verbose", "Verbose Mode")

	flaggy.Parse()

	switch Scanner {
	case "stdin":
		if !Utility.IsInPipeline() {
			fmt.Println("Please pipe input to the program, or use -s file/cidr")
			os.Exit(1)
		}
		Type = ESTDIN

	default:
		if Utility.IsCIDRv4(Scanner) {
			Type = ECIDR
		} else if Utility.IsIPv4(Scanner) {
			Type = EIP
		} else if Utility.FileExists(Scanner) {
			Type = EFILE
		} else {
			fmt.Println("Invalid Input, possible inputs: stdin, filename, cidr, ip")
			os.Exit(1)
		}
	}

	if Utility.FileExists(comboFile) {
		combos, err := Utility.ReadFileLines(comboFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(combos) < 1 {
			fmt.Println("Combo file is empty")
			os.Exit(1)
		}

		for _, combo := range combos {
			userPass := strings.Split(combo, ":")
			if len(userPass) == 2 {
				Users = append(Users, userPass[0])
				Passwords = append(Passwords, userPass[1])
			}
		}
	} else {
		Users = []string{"anonymous"}
		Passwords = []string{"anonymous"}
	}

	if len(Users) < 1 || len(Passwords) < 1 {
		fmt.Println("Combo file Does not contain any credential with user:password format")
		os.Exit(1)
	}

	if ports != "" {
		portsSplited := strings.Split(ports, ",")
		for _, port := range portsSplited {
			for _, r := range port {
				if !unicode.IsDigit(r) {
					fmt.Printf("Invalid Port on -p Argument, Port: \"%s\"", port)
					os.Exit(1)
				}
			}
			Ports = append(Ports, port)
		}
	} else {
		Ports = []string{"21"}
	}

	if fileFormats != "" {
		fileFormatsSplited := strings.Split(fileFormats, ",")
		FileFormats = fileFormatsSplited
	}

	if !Utility.FolderExists(OutputFolder) {
		err := Utility.CreateFolder(OutputFolder)
		if err != nil {
			fmt.Printf("Failed to create output folder: %s\n", err)
			os.Exit(1)
		}
	}

	if SaveCredentials {
		if !Utility.FileExists("hit.txt") {
			err := Utility.CreateFile("hit.txt")
			if err != nil {
				fmt.Printf("Failed to create hit.txt: %s\n", err)
				os.Exit(1)
			}
		}

		file, err := os.OpenFile("hit.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Printf("Failed to open hit.txt: %s\n", err)
			os.Exit(1)
		}

		CredFile = Utility.NewMutexWriter(file)
	}
}
