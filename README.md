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
2. Import into your project
   ```bash
   printf '// +build tools\npackage tools\nimport (_ "github.com/DenChenn/blunder")' | gofmt > tools.go

   go mod tidy
   ```
3. Initialize Blunder
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

## Usage
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

Your error will be generated in given path like this:
```
<your_dir_path>/errors/
generated/
  alayer/ 
    errors.go
  blayer/
    errors.go
blunder.yaml
```

Which can be import into your code like this:
```go
if err != nil {
  if errors.Is(err, &alayer.Err1) {
    //...
  }
}
```

Or you can wrap your error like this:
```go
if errors.Is(err, &pgx.ErrNoRows) {
  return &alayer.Err1.Wrap(err) 
}
```

## Type assertion
All generated errors implement `blunder.OrdinaryError` interface, which contains static methods.

```go
ordinaryError, ok := err.(blunder.OrdinaryError)
if ok {
	fmt.Println(ordinaryError.GetId())
}
```
