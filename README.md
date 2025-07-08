# cloudstart

A production-ready Go CLI tool template using Uber Fx, Uber Config, and Cobra.

## Project Structure

```
.
├── cmd/
│   └── cloudstart/
│       └── main.go         # Entry point for the CLI tool
├── internal/
│   └── info/
│       └── info.go         # Internal implementation of the 'info' command
├── pkg/
│   └── cli/
│       └── cli.go          # CLI setup and wiring
├── configs/
│   └── config.yaml         # Example config file
├── Taskfile.yml            # Task runner tasks
├── go.mod                  # Go module definition
└── README.md
```

## Usage

### Build
```sh
task build
```

### Run
```sh
go run cmd/cloudstart/main.go info
```

### Example Output
```
App Name: cloudstart
App Version: 1.0.0
Cloud Provider: aws
Cloud Region: us-west-2
```

## Conventions
- All entry points go in `cmd/<appname>/main.go`.
- Internal-only code goes in `internal/`.
- Reusable CLI setup and providers go in `pkg/cli/`.
- Config files go in `configs/`.

## Requirements
- Go 1.22+
- [Task](https://taskfile.dev) (for task runner)

---
This template is ready for production use and extensible for new commands and features.
