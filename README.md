# Blunder
![GitHub](https://img.shields.io/github/license/DenChenn/blunder)
![Go Report Card](https://goreportcard.com/badge/github.com/DenChenn/blunder)

## What is Blunder?
Blunder is a simple, gpt-based and easy-to-use error handling package for golang. 
It generate typed errors and manage them in centralized way, which reduce the dependency between packages and make error handling more convenient.

## Getting Started
1. Install package
   ```bash
    go get github.com/DenChenn/blunder
   ```
2. Initialize Blunder
   ```bash
    go run github.com/DenChenn/blunder init <your_dir_path>
   ```
3. Define your error in `blunder.yaml` according to example ❤️
4. Generate all errors
   ```bash
    go run github.com/DenChenn/blunder gen
   ```
   or generate with gpt-based auto-completion
   ```bash
    export OPENAI_API_TOKEN=<your_openai_api_token>
    go run github.com/DenChenn/blunder gen --complete=true 
   ```

## Usage [WIP]
Suppose your errors in `blunder.yaml` are defined like this:
```yaml
details:
- package: alayer
  errors:
    - code: Err1
      #...
    - code: Err2
      #...
    - code: Err3
      #...
- package: blayer
  errors:
    - code: Err1
      #...
    - code: Err2
      #...
    - code: Err3
      #...
```