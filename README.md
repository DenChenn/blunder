# Blunder
![GitHub](https://img.shields.io/github/license/DenChenn/blunder)
![Go Report Card](https://goreportcard.com/badge/github.com/DenChenn/blunder)

## What is Blunder?
Blunder is a simple, gpt-based and easy-to-use error handling package for golang.

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

### Case: with errors from other packages 
This is a basic example for using blunder : )

<table>
<thead><tr><th>Before</th><th>Blunder</th></tr></thead>
<tbody>
<tr><td>

```go
if err := pgx.QueryRow().Scan(); err != nil {
  if errors.Is(err, pgx.ErrNoRows) {
    return nil, a.Err1
  }
  return nil, err 
}
```

</td><td>

```go
if err := pgx.QueryRow().Scan(); err != nil {
  return blunder.
    Match(err, pgx.ErrNoRows, a.Err1).
    Return() 
}
```

</td></tr>
</tbody></table>

### Case: across grpc 
This case is a little bit complicated, but blunder can handle it easily.  
Handling across layers is a common case in microservice architecture, and blunder is designed for this case.
Moreover, with condition technique, blunder can handle more complicated error mapping cases.

<table>
<thead><tr><th>Before</th><th>Blunder</th></tr></thead>
<tbody>
<tr><td>

```go
// server side
if err := someFunction(); err != nil {
  if errors.Is(err, lowerLayerErr1) || errors.Is(err, lowerLayerErr2) {
    return nil, status.Error(codes.NotFound, a.Err1.Error())
  } errors.Is(err, lowerLayerErr3) {
    return nil, status.Error(codes.NotFound, a.Err2.Error())
  } 
  // ...some more cases 
  return nil, status.Error(codes.Internal, err.Error())
}

// client side
if err := rpc.Call(); err != nil { 
  if s, ok := status.FromError(unknownErr); ok {
    switch s.Message() {
    case a.Err1.Error():
      return b.Err1
    case a.Err2.Error():
      return b.Err2
    default:
      return err // <- undefined
    }
  }
}
```

</td><td>

```go
// server side
if err := someFunction(); err != nil {
  cond := blunder.NewCondition().
    ManyToOne([]error{lowerLayerErr1, lowerLayerErr2}, a.Err1).
    OneToOne(lowerLayerErr3, a.Err2) // some more cases
  return blunder.
        MatchCondition(err, cond).
        ReturnForGrpc()
}

// client side
if err := rpc.Call(); err != nil {
  cond := blunder.NewCondition().
    OneToOne(a.Err1, b.Err1).
    OneToOne(a.Err2, b.Err2) // some more cases
  return blunder.
    MatchCondition(err, cond).
    Return()
}
```

</td></tr>
</tbody></table>

### Case: do something detailed 
Blunder can simply be used as `errors.Is` and `errors.As` to do something detailed.
You can still use `MatchCondition` to handle errors that requiring the same operations.

<table>
<thead><tr><th>Before</th><th>Blunder</th></tr></thead>
<tbody>
<tr><td>

```go
if err := someFunction(); err != nil {
  if errors.Is(err, Err1) || errors.Is(err, Err2) {
    doSomething()
    return nil, a.Err1
  } else if errors.Is(err, Err3) {
    doOtherthing() // no returning
  }
  return nil, err 
}
```

</td><td>

```go

if err := someFunction(); err != nil {
  cond := blunder.NewCondition().ManyToOne([Err1, Err2], a.Err1)
  if blunder.MatchCondition(err, cond).GetIsMatched() {
    doSomething()
    return nil, a.Err1
  } else if blunder.Match(err, Err3).GetIsMatched() {
    doOtherthing()
  }
  return nil, blunder.ErrUndefined
}
```

</td></tr>
</tbody></table>

## Return framework support
- [x] Gin
- [x] Grpc
- [ ] Gqlgen