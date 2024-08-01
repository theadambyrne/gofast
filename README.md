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

- Models: Creates model with necessary configurations, relations, and validations.
- Views: Scaffolds views using Shadcn-UI, enabling immediate CRUD operations.
- Controllers: Generates API routes.

### Commands 

- `gofast new` - Create a new project (set up and config files ...)
- `gofast model` - Create a new model (fields, relations, validations ...)
- `gofast view` - Create a new view (CRUD operations)
- `gofast controller` - Create a new controller (API routes)
- `gofast scaffold` - Create a new scaffold (model, view, controller)

**Stack:** Go, [shadcn-ui/ui](https://github.com/shadcn-ui/ui) and [htmx](https://github.com/bigskysoftware/htmx).

**Drivers:** sqllite3
