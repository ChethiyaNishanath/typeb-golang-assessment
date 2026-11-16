## How to run the application

With Make:

```sh
make run          # run from source
make build-run    # build bin/server(.exe), then run it [windows build]
```

Without Make:

```sh
go run ./cmd
```

## How to run the tests

```sh
make test         # run all unit tests (no caching)
make build-test   # build the binary, then run all unit tests
```

Without Make:

```sh
go test ./... -v -count=1
```

Run a single test:

```sh
go test ./internal/greeting -run TestGreet_EmptyName -v
```

## Assumptions

- **The name comes from the `name` query parameter.** Leading and trailing spaces are trimmed before validation.
- **An empty or missing name is invalid.**
- **All validation failures return `400 Bad Request` with the same `"Invalid Input"` message.** The specific reason (empty name vs. second-half letter) is not shown to the client.