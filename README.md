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
- Chat (AI Chat from the terminal) -**Requires Google AI Studio api key**


## Usage
To build:
```bash
cd flow
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

To run the chat:
```bash
./flow chat
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
