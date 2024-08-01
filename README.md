# gofast

[![build](https://github.com/charmbracelet/bubbletea-app-template/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/charmbracelet/bubbletea-app-template/actions/workflows/build.yml)[![Dependabot Updates](https://github.com/charmbracelet/bubbletea-app-template/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/charmbracelet/bubbletea-app-template/actions/workflows/dependabot/dependabot-updates)[![golangci-lint](https://github.com/charmbracelet/bubbletea-app-template/actions/workflows/lint.yml/badge.svg)](https://github.com/charmbracelet/bubbletea-app-template/actions/workflows/lint.yml)[![release](https://github.com/charmbracelet/bubbletea-app-template/actions/workflows/release.yml/badge.svg)](https://github.com/charmbracelet/bubbletea-app-template/actions/workflows/release.yml)

A TUI app written in Go to scaffold fullstack applications. Generate CRUD operations and their respective API routes, database Model (migrations), and frontend components. Build an app in 10 minutes! - inspired by [kirimase](https://kirimase.dev).

## Installation

```bash
go get github.com/theadambyrne/gofast
```

## Usage

```bash
gofast
```

## Napkin plan

- Init screen: choose stack/drivers (1 option for now)
- Generate screen: choose what to generate (CRUD, API, Model, Frontend)
- CRUD screen: choose entity name, fields, types, and validation
- API screen: choose entity name, fields, types, and validation
- Model screen: choose entity name, fields, types, and validation
- Frontend screen: choose entity name, fields, types, and validation
-  Generate files
-  Run migrations
-  Run frontend
-  Run backend
-  Open browser
