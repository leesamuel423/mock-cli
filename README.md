# Interview CLI

A command-line tool for practicing interview questions. Add your own questions, categorize them, and conduct practice sessions.

## Features

- Add interview questions with categories (behavioral/technical) and tags
- List all stored questions grouped by category
- Conduct practice sessions with customizable filters and number of questions
- Local storage of questions in project directory

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/interview-cli.git
cd interview-cli

# Build the binary
make build
```

## Usage

### Add a new question

```bash
make add
# or
./interview-cli -add
```

### List all questions

```bash
make list
# or
./interview-cli -list
```

### Start a practice session

```bash
# Default practice session (5 random questions)
make practice
# or
./interview-cli -practice

# Custom number of questions
make practice N=10
# or
./interview-cli -practice -n 10

# Filter by category
make practice CATEGORY=behavioral
# or
./interview-cli -practice -category behavioral

# Filter by tags
make practice TAGS="leadership,problem-solving"
# or
./interview-cli -practice -tags "leadership,problem-solving"

# Combine options
make practice N=3 CATEGORY=technical TAGS="algorithms,data-structures"
# or
./interview-cli -practice -n 3 -category technical -tags "algorithms,data-structures"
```

## Available Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the binary |
| `make run` | Build and run the binary (runs default practice session) |
| `make clean` | Remove the binary |
| `make lint` | Run linter and vet |
| `make vet` | Run go vet (checks for potential issues in code) |
| `make fmt` | Format code |
| `make practice` | Start practice session |
| `make list` | List all questions |
| `make add` | Add a new question |
| `make test` | Run all tests |
| `make test-short` | Run only unit tests (skip integration tests) |
| `make test-verbose` | Run tests with verbose output |
| `make test-coverage` | Run tests with coverage report |

## Testing

The project includes a comprehensive test suite located in the `tests/` directory. The tests cover the core functionality of the application:

- Model validations and utility functions
- Repository operations (loading, saving, filtering questions)
- Session functionality
- Basic CLI commands

### Running Tests

```bash
# Run all tests
make test

# Run only unit tests (skip integration tests)
make test-short

# Run tests with verbose output
make test-verbose

# Generate test coverage report
make test-coverage
```

The test coverage report will be generated in `coverage.html` and can be viewed in a web browser.

## Data Storage

All questions are stored in a JSON file (`questions.json`) in the project directory. The configuration file (`config.json`) is also stored in the project directory.

## License

See the [LICENSE](LICENSE) file for details.