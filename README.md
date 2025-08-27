# Golang Wails ReactJS Application

A modern desktop application built with Wails v2, featuring a Go backend with clean architecture and a React TypeScript frontend.

## Features

- **Backend**: Go with clean architecture (domain, application, infrastructure layers)
- **Frontend**: React 18 + TypeScript with Vite
- **UI Framework**: React Bootstrap 5
- **Database**: SQLite with Go database layer
- **Authentication**: User authentication system
- **Validation**: Input validation with go-playground/validator
- **Routing**: React Router DOM for navigation

## Project Structure

```
├── backend/
│   ├── application/     # Application services
│   ├── domain/         # Domain entities and business logic
│   └── pkg/           # Shared packages (database, validation, etc.)
├── frontend/
│   ├── src/           # React TypeScript source
│   └── wailsjs/       # Generated Wails bindings
├── app.go             # Wails application context
├── main.go            # Application entry point
└── wails.json         # Wails configuration
```

## Prerequisites

- Go 1.23+
- Node.js 16+
- Wails v2

## Installation

1. **Install Wails**:
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

2. **Clone and setup**:
   ```bash
   git clone <repository-url>
   cd golang-wails-reactjs
   ```

3. **Install dependencies**:
   ```bash
   # Backend dependencies
   go mod tidy
   
   # Frontend dependencies
   cd frontend
   npm install
   cd ..
   ```

## Development

**Live development mode**:
```bash
wails dev
```

This starts a Vite development server with hot reload. Access the browser dev server at http://localhost:34115 to call Go methods from devtools.

## Building

**Production build**:
```bash
wails build
```

**Build options**:
- `-clean`: Clean build directory
- `-platform`: Target platform (windows, darwin, linux)
- `-ldflags`: Go linker flags

## Configuration

Edit `wails.json` for project settings. See [Wails documentation](https://wails.io/docs/reference/project-config) for details.

## Dependencies

**Backend**:
- Wails v2.10.2
- SQLite driver
- Validator v10
- Crypto utilities

**Frontend**:
- React 18 + TypeScript
- React Bootstrap
- React Router DOM
- Vite build tool
