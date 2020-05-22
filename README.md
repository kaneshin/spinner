# Spinner - Spin Spin Spin

## Installation

`spinner` provides the command-line tools.

```shell
$ go get github.com/kaneshin/spinner/cmd/...
```

## Usage

### `spinner` command

```shell
$ spinner -i 0.05 -d 3
$ spinner -i 0.05 -c 'sleep 5'
```

### `spinner` package

```go
interval := 50 * time.Millisecond

sp := spinner.New(interval, func(ctx context.Context) {
	doSomething(ctx)
})
sp.Do(context.Background())

spinner.Run(context.Background(), interval, func(sp *spinner.Spinner) func(context.Context) {
	return func(ctx context.Context) {
		doSomething(ctx)
	}
})
```

## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)

## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
