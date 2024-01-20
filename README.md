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
#### Modes
- CLI
- Server

## Roadmap
Current Version: **0.0.4** as of **01-19-2024**
- [ ] *Version 1.0.0*
  - [x] Task service creation
  - [x] Planning service creation
  - [x] Goal service creation
  - [x] Plan service~~ creation
  - [X] Basic CLI - 0.0.5
    - [X] MVP
    - [ ] Testing
  - [X] Basic Rest API - 0.0.5
    - [X] MVP
    - [ ] Testing
  - [ ] Simple reporting service - 0.0.6
  - [ ] Bulk Upload and management - 0.0.7
  - [ ] Flow Modules (pre-defined "flows") - 0.0.8
    - [ ] Project Module
    - [ ] Learning Module

-[ ] *Version 2.0.0*
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

