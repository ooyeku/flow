# Go Flow

This project is a task/workflow manager written in Go. It is designed to be
simple, easy to use, and easy to extend. It is also designed to be
cross-platform, and should work on any system that Go supports.

## Project Structure

The project is structured as follows:
#### Layers:
- Model
- Store
- Service
- Handle (control)
- API
#### Modes:
- CLI
- Server

## Installation
To install the CLI, run the following command:
```bash
go get github.com/ooyeku/workflow/
```


## Usage
To build:
```bash
cd workflow
go build
```

To run the CLI:
```bash
./flow cli
```

To run the server:
```bash
./flow server
```

