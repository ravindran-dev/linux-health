<p align="center">
  <img src="https://img.shields.io/badge/Linux%20Health-System%20Checkup-blueviolet?style=for-the-badge" />
</p>

# Linux-health
A lightweight, fast, and dependency-free Linux system health monitoring tool written in Go.
The tool provides a clean, boxed CLI view of essential system metrics for quick diagnostics.



## Overview

`linux-health` is designed to give a concise yet meaningful snapshot of a Linux system's current health.
It avoids heavy dependencies, background daemons, or graphical interfaces, making it suitable for servers, virtual machines, and SSH environments.



## Features

### System Metrics

* CPU usage
* Memory usage
* Disk usage
* Load average
* System uptime

### Process Analysis

* Detects the top memory-consuming process
* Displays process name, PID, and memory usage

### Disk Analysis

* Identifies the largest directories (disk hotspots) inside the user home directory

### Service Health

* Detects failed or inactive systemd services
* Displays a healthy message when no issues are found

### Network Status

* Determines whether any non-loopback network interface is active

### Output

* Clean, boxed terminal output
* Designed for readability and consistency across terminals



## Project Structure

```
linux-health/
├── cmd/
│   └── linux-health/
│       └── main.go
│
├── internal/
│   ├── system/       
│   ├── process/       
│   ├── disk/        
│   ├── service/      
│   └── network/      
│
├── pkg/
│   └── output/        
│       └── block.go
│
├── go.mod
└── README.md
```



## Requirements

* Linux operating system
* Go 1.20 or newer
* systemd (for service health checks)



## Installation

Clone the repository:

```bash
git clone https://github.com/ravindran-dev/linux-health.git
cd linux-health
```

Build the binary:

```bash
go build -o linux-health ./cmd/linux-health
```



## Usage

Run directly using Go:

```bash
go run ./cmd/linux-health
```

Or run the compiled binary:

```bash
./linux-health
```



## Sample Output

```
┌──────────────────────────────────────┐
│          LINUX SYSTEM HEALTH         │
├──────────────────────────────────────┤
│ SCORE   : 100 / 100  [ GOOD ]        │
│ UPTIME  : 0h 7m                      │
├──────────────────────────────────────┤
│ CPU     : 3.8 %     [ OK ]           │
│ MEMORY  : 13.8 %    [ OK ]           │
│ DISK    : 77.9 %    [ OK ]           │
│ LOAD    : 0.74      [ OK ]           │
├──────────────────────────────────────┤
│ TOP MEM : gnome-shell                │
│           PID 2381   269 MB          │
├──────────────────────────────────────┤
│ HOTSPOTS:                            │
│   /home/user/Downloads 12.4G         │
├──────────────────────────────────────┤
│ SERVICES:                            │
│   No failed services                 │
├──────────────────────────────────────┤
│ NETWORK : active                     │
└──────────────────────────────────────┘
```



## Design Principles

* Minimal dependencies
* Predictable and stable output
* Suitable for automation and scripting
* Clear separation between data collection and presentation



## Future Enhancements

* Optional colorized output
* Watch mode for periodic refresh
* Export output as JSON
* Threshold-based alerts
* 
## License

MIT License

##  Author - **Ravindran S** 


Developer • ML Enthusiast • Neovim Customizer • Linux Power User  

Hi! I'm **Ravindran S**, an engineering student passionate about:

-  Linux & System Engineering  
-  AIML (Artificial Intelligence & Machine Learning)  
-  Full-stack Web Development  
-  Hackathon-grade project development  




## 🔗 Connect With Me

You can reach me here:

###  **Socials**
<a href="www.linkedin.com/in/ravindran-s-982702327" target="_blank">
  <img src="https://img.shields.io/badge/LinkedIn-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white">
</a>


<a href="https://github.com/ravindran-dev" target="_blank">
  <img src="https://img.shields.io/badge/GitHub-111111?style=for-the-badge&logo=github&logoColor=white">
</a>


###  **Contact**
<a href="mailto:ravindrans.dev@gmail.com" target="_blank">
  <img src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white">
</a>



