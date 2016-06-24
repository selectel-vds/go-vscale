# go-vscale

go-vscale is a Go client library for accessing the [Vscale API](https://developers.vscale.io/documentation/api/v1/).

## Usage

```go
import "github.com/vscale/go-vscale"
```

First step is constructing Vscale client which allows to use API services.
You can generate token in [Vscale Panel](https://vscale.io/panel/settings/tokens/).

```go
client := NewClient("token should be here")
account, _, err := client.Account.Get()
```

Some operations with scalets can be started both sync and async.

```go
// Second argument is "wait" which expects boolean value
// true - if you want to wait until the end of operation
// false - if you want this operation to be handled in background
client := NewClient("token should be here")
scalet, _, err := client.Scalet.Rebuild(11111, true)
```

## Tests

You can run tests which make requests straightly to Vscale API.
For now they can't be run together. Run specific test if you want to test some method.

```bash
$ go test -v github.com/vscale/go-vscale -run TestAccountService_Get
```

For convenience you can use "VSCALE_API_TOKEN" env for not passing token to every test.
