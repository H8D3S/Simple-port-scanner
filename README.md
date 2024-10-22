# Simple Port Scanner

This is a basic TCP port scanner written in Go that scans a range of ports on a target machine to check if they are open.

## Features
- Scans a range of ports and reports which ones are open.
- Lightweight and simple to use.

## Prerequisites
- Go 1.16 or higher

## How to Use

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/port-scanner.git
   ```

2. Navigate to the project directory:
   ```bash
   cd port-scanner
   ```

3. Run the program:
   ```bash
   go run portscanner.go
   ```

By default, it scans ports 1-1024 on `scanme.nmap.org`. You can modify the target or port range in the `portscanner.go` file.

## Example Output
```
Starting scan on scanme.nmap.org from port 1 to 1024
Port 22 is open
Port 80 is open
```

## Disclaimer
This tool is for educational purposes only. Do not use it to scan hosts without proper authorization. Unauthorized port scanning may violate your network's terms of service.
