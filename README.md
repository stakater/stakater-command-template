# cloudstart

A production-ready Go CLI tool template using Uber dig, Uber Config, and Cobra.

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
│       └── cli.go          # CLI setup and dependency injection
├── configs/
│   └── config.yaml         # Example config file
├── Taskfile.yml            # Task runner tasks
├── go.mod                  # Go module definition
└── README.md
```

## Configuration

### Configuration Override

The CLI supports environment-specific configuration overrides using the `--env` flag:

- **Default config**: `configs/config.yaml`
- **Environment override**: `configs/config.{env}.yaml`

When you specify `--env <environment>`, the CLI will:
1. Load the base configuration from `configs/config.yaml`
2. Override with environment-specific settings from `configs/config.{env}.yaml`

#### Example Configuration Files

**Base config (`configs/config.yaml`):**
```yaml
app:
  name: cloudstart
  version: 1.0.0
cloud:
  provider: aws
  region: us-west-2
```

**Local override (`configs/config.local.yaml`):**
```yaml
cloud:
  region: us-east-1
```

**E2E override (`configs/config.e2e.yaml`):**
```yaml
cloud:
  provider: gcp
  region: europe-west1
```

## Usage

### Build
```sh
task build
```

### Run
```sh
task run -- info
task run -- info --env local
task run -- info --env e2e
```

#### Usage Examples

```sh
# Use default config
task run -- info

# Use local environment override
task run -- info --env local

# Use e2e environment override  
task run -- info --env e2e
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
- Reusable CLI setup and dependency injection go in `pkg/cli/`.
- Config files go in `configs/`.

## Requirements
- Go 1.22+
- [Task](https://taskfile.dev) (for task runner)

## Versioning and Automated Releases

This project uses automated semantic versioning and release notes based on commit messages, following the Angular/Conventional Commits convention. The version bump is determined by your commit message prefix:

- **fix:** → PATCH release (e.g., v1.2.0 → v1.2.1)
- **feat:** → MINOR release (e.g., v1.2.0 → v1.3.0)
- **BREAKING CHANGE:** in the body or footer → MAJOR release (e.g., v1.2.0 → v2.0.0)

### Examples

- **Patch release:**
  ```
  fix: correct typo in CLI output
  
  This fixes a small typo in the help message.
  ```
- **Minor release:**
  ```
  feat: add new cloud provider support
  ```
- **Major release:**
  ```
  feat: overhaul config system
  
  BREAKING CHANGE: config file format is now YAML only
  ```

When you push to the main branch with one of these commit messages, the CI will automatically bump the version and create a new release accordingly.

---
This template is ready for production use and extensible for new commands and features.
