# Go Workflow


This project is a task/workflow manager written in Go. It is designed to be
simple, easy to use, and easy to extend. It is also designed to be
cross-platform, and should work on any system that Go supports.

## Project Structure

The project is structured as follows:
* `conf/` - Configuration of controllers, services, and other components
* `handle/` - Routing and handling of requests
* `models/` - Data models
* `services/` - Business logic
* `store/` - Data storage

## Roadmap
Current Version: **0.0.3** as of **01-11-2024**

*[ ] Version 1.0.0*
  - [x] Task service creation
  - [x] Planning service creation
  - [x] Goal service creation
  - [x] Plan service creation
  - [ ] Simple reporting service
  - [ ] Basic CLI
  - [ ] Basic Rest API

*[ ] Version 2.0.0*
    - [ ] Web UI
    - [ ] Advanced reporting service


## Planned Usage

### CLI

The CLI will be the primary interface for the application. It will be used to
manage tasks, workflows and reports.  It will also be used to configure the
application (along with config.json).

### REST API

The REST API will be used to allow other applications to interact with the
application. It will be used to manage tasks, workflows and reports.

### Web UI

The web UI will be used to allow users to interact with the application.

## Planned Features
- [ ] Task management
- [ ] Workflow management
- [ ] Planning
- [ ] Reporting
- [ ] CLI
- [ ] Web UI
- [ ] REST API

