# go-pinentry-minimal

[![PkgGoDev](https://pkg.go.dev/badge/github.com/twpayne/go-pinentry-minimal)](https://pkg.go.dev/github.com/twpayne/go-pinentry-minimal)

Package `pinentry` provides a minimal client to [GnuPG's
pinentry](https://www.gnupg.org/related_software/pinentry/index.html), which
only depends on Go's standard library. It is a fork of
[`github.com/twpayne/go-pinentry`](https://github.com/twpayne/go-pinentry).

## Key Features

* Support for all `pinentry` features.
* Idiomatic Go API.
* Forked from a well-tested library.

## Example

```go
	client, err := pinentry.NewClient(
		pinentry.WithBinaryNameFromGnuPGAgentConf(),
		pinentry.WithDesc("My description"),
		pinentry.WithGPGTTY(),
		pinentry.WithPrompt("My prompt:"),
		pinentry.WithTitle("My title")
	)
	if err != nil {
		return err
	}
	defer client.Close()

	switch pin, fromCache, err := client.GetPIN(); {
	case pinentry.IsCancelled(err):
		fmt.Println("Cancelled")
	case err != nil:
		return err
	case fromCache:
		fmt.Printf("PIN: %s (from cache)\n", pin)
	default:
		fmt.Printf("PIN: %s\n", pin)
	}
```

## License

MIT
