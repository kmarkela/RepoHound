# RepoHound

```
██████╗ ███████╗██████╗  ██████╗ ██╗  ██╗ ██████╗ ██╗   ██╗███╗   ██╗██████╗ 
██╔══██╗██╔════╝██╔══██╗██╔═══██╗██║  ██║██╔═══██╗██║   ██║████╗  ██║██╔══██╗
██████╔╝█████╗  ██████╔╝██║   ██║███████║██║   ██║██║   ██║██╔██╗ ██║██║  ██║
██╔══██╗██╔══╝  ██╔═══╝ ██║   ██║██╔══██║██║   ██║██║   ██║██║╚██╗██║██║  ██║
██║  ██║███████╗██║     ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║ ╚████║██████╔╝
╚═╝  ╚═╝╚══════╝╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═════╝                                                        
```


GitHub Repo discovery Tool. 


## Overview

Main usecase for the tool is discovery of public GitHub repos for list of username during Red Team Operations. 

## Install 

There are three ways to install the RepoHound: building from source, using `go install`, or downloading pre-compiled binaries.

### Building from Source

```sh
git clone https://github.com/kmarkela/RepoHound
cd RepoHound
go build -o RepoHound
```

### Using `go install`

```sh
go install github.com/kmarkela/RepoHound@latest
```

### Binaries 

If you prefer to download a pre-compiled binary for your platform, follow these steps:

1. Navigate to the [Releases page](https://github.com/kmarkela/RepoHound/releases) of the DuffMan repository.
2. Download the appropriate binary for your operating system and architecture.
3. Move the binary to your desired location.

## Usage

```sh
./rh -h
Discovery of public repos for list of users

Usage:
  RepoHound [flags]

Flags:
  -h, --help              help for RepoHound
      --json              Output in JSON Format
  -p, --proxy string      HTTP Proxy
  -u, --userlist string   userlist
      --workers int       Workers (default 5)
```

Example:
```sh
go run main.go -u ulist.txt
Found The Following Repos:
https://github.com/kmarkela/duffman
https://github.com/kmarkela/generative-ai-for-beginners
https://github.com/kmarkela/wiggumizer
https://github.com/kmarkela/SimplePortScanner

```

## License 

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Disclamer 

The RepoHound is intended for security research and testing purposes only. Ethical conduct is required from all users.

The author(s) of this tool take no responsibility for any misuse of the software. It is the end user's responsibility to comply with all applicable local, state, federal, and international laws. By using this tool, you agree that you hold responsibility for any consequences that arise from its use.
