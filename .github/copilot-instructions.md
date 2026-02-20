# CLIProxyAPI Plus - GitHub Copilot Custom Instructions

Welcome to the CLIProxyAPI Plus project! This repository is a feature-rich proxy API written primarily in Go, with a TypeScript-based MCP server (`mcp-cliproxyapi`) and a Vue frontend (`management-panel`).

When generating code, answering questions, or refactoring in this repository, please adhere strictly to the following guidelines and conventions.

## 1. Project Architecture & Languages

The repository is divided into three main components:
- **Core Backend (`/`, `cmd/`, `internal/`)**: Go (Golang) 1.26+. Uses `gin-gonic/gin` for HTTP routing.
- **MCP Server (`mcp-cliproxyapi/`)**: TypeScript / Node.js. Uses `@modelcontextprotocol/sdk`.
- **Management Panel (`management-panel/`)**: Vue 3 / Vite (frontend UI).

### Go Backend Guidelines
- **Frameworks**: Use `gin-gonic/gin` for new HTTP routes, middlewares, and handlers (`internal/api/handlers/`, `internal/api/middleware/`).
- **Structure**: 
  - `cmd/server/main.go`: Application entrypoint.
  - `internal/auth/`: Integration logic for different AI providers (e.g., `copilot/`, `kiro/`, `openai/`, `claude/`).
  - `internal/config/`: Configuration parsing (`config.yaml`).
  - `internal/translator/`: Request/Response translation between unified formats and provider-specific formats.
- **Error Handling**: Do not swallow errors. Always return or log errors using the established logger in `internal/logging/`. Use structured logging (`logrus` or the wrapper provided in the repo).
- **Concurrency**: Use Go routines carefully. When spawning background tasks (like background token refreshes), ensure proper synchronization (mutexes, waitgroups) and context cancellation (`context.Context`).
- **Formatting**: Adhere to standard `gofmt` and `goimports` rules. Keep functions short and focused.

### TypeScript / Node.js Guidelines (`mcp-cliproxyapi`)
- **Language**: Use strict TypeScript (`tsconfig.json`). No `any` types unless absolutely necessary.
- **Validation**: Use `zod` for parsing and validating configuration and API payloads.
- **Asynchronous Code**: Use `async`/`await` over raw Promises. Handle rejections using `try...catch`.
- **Structure**: Follow the existing split between `index.ts` (entrypoint), `cliproxy.ts` (API client), `config.ts` (environment loading), and `tools.ts` (MCP tool definitions).

### Vue Frontend Guidelines (`management-panel`)
- **Composition API**: Use Vue 3 Composition API (`<script setup>`).
- **Styling**: Follow existing styling conventions (likely TailwindCSS or scoped CSS depending on existing `.vue` files).
- **Reactivity**: Use `ref` and `computed` appropriately.

## 2. Build and Run Commands

When suggesting how to run or test the code, use these specific commands:

### Go Backend
- Run locally: `go run cmd/server/main.go -config config.yaml` (or via `start.bat`)
- Build binary: `go build -o CLIProxyAPI.exe cmd/server/main.go`
- Run tests: `go test ./...`
- Docker build (Linux/macOS): `./docker-build.sh`
- Docker build (Windows): `.\docker-build.ps1`

### MCP Server (`mcp-cliproxyapi`)
- Install dependencies: `npm install`
- Run development server (stdio): `npm run dev`
- Run development server (http): `npm run dev:http`
- Build TypeScript: `npm run build`

## 3. General Coding Conventions

- **Third-Party Providers**: This "Plus" fork contains specific third-party integrations (e.g., GitHub Copilot OAuth, Kiro/AWS CodeWhisperer). Keep logic for different providers strictly isolated in their respective folders under `internal/auth/` and `internal/translator/`.
- **Documentation**: Write GoDoc-style comments for public functions, structs, and interfaces. For TypeScript, use JSDoc format for exported functions and types.
- **Config**: Do not hardcode values like ports, hosts, or URLs. Always load them from `config.yaml` (in Go) or `process.env` (in Node.js).
- **Imports**: Avoid circular dependencies. Group standard library imports separately from third-party imports.
- **File Naming**: 
  - Go: `snake_case.go` or `snake_case_test.go`.
  - TypeScript: `kebab-case.ts` or `snake_case.ts` following existing files (`http_server.ts`).

## 4. Specific Context Hints

- **OAuth & Tokens**: When working on OAuth features (like the Kiro web login), ensure secure handling of tokens. Do not print access tokens, refresh tokens, or client secrets in logs.
- **Device Fingerprints**: Understand that the Plus version includes enhanced security like device fingerprints and rate limiting. Ensure any new provider integration respects the rate limiter middleware.
- **MCP Context**: When adding new tools to the `mcp-cliproxyapi`, remember to define the input schema using `zod` and implement the execution logic in `tools.ts` or equivalent, then register it in `index.ts` with the `@modelcontextprotocol/sdk`.

When generating code, please provide the complete block of code that needs to be inserted or replaced, ensuring it fits seamlessly into the existing architecture.