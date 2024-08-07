# PortPatrol

PortPatrol is a network management and security tool designed to monitor network traffic, generate logs of open ports, and block unused ports by activating the firewall. This tool helps administrators maintain server security by automating the process of identifying and closing unnecessary ports.

## Features

- **Network Traffic Monitoring**: Monitors network traffic on the server.
- **Log Generation**: Creates log files of open ports after a specified duration.
- **Firewall Activation**: Automatically blocks unused ports based on the generated log.
- **Scheduled Execution**: Allows users to schedule regular intervals for monitoring and securing the server.

## Project Structure


- `device/`: Contains code for device-related functionalities.
- `firewall/`: Contains code for firewall activation and port management.
- `loger-port/`: Handles logging of open ports.
- `packets/`: Manages packet-related functionalities.
- `dockercompose.yml`: Docker Compose configuration.
- `Dockerfile`: Docker configuration for building the application.
- `go.mod`: Module definition for Go dependencies.
- `go.sum`: Checksums for Go dependencies.
- `LICENSE`: License information.
- `main.go`: Main entry point of the application.

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/SUmidcyber/PortPatrol.git
   cd PortPatrol
## Build the Docker image:

    docker build -t portpatrol .
## Run the Docker container:

    docker-compose up
## Usage

## Start the application:
    go run main.go
Choose an option:

    To monitor and analyze network traffic, press 1.
    To close open ports, press 2.
    Close unused ports:
    
    After monitoring, the application will generate a log of open ports.
    Choose whether to close unused ports based on the log file.


