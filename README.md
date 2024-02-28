# Flow ♾️

Easily create and manage your own workflows from the terminal.  
Flow is a CLI tool that allows you to create and manage goals, tasks, projects, and more from the terminal,
while also providing a chat mode to interact with an AI chatbot. The project began as a keep track of projects and workflows
from the terminal, but has since evolved to include a chat mode that uses the PerplexityAI API (more coming soon).

### Status
This project is still in development and is not yet ready for production use.  The CLI, Server, and Chat modes are all functional, but 
features are still being added and bugs are still being fixed.  The project is being developed in my free time, so updates may take a while.  if you would like to contribute, please feel free to submit a pull request.


#### Modes:
- CLI
- Server
- Chat **Requires PerplexityAI API Key** [Chat README](cmd/chat/README.md)


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
