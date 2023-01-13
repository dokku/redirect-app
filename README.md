# redirect-app

An app whose sole purpose is to redirect requests to a given url.

## Requirements

- Go 1.19+

## Usage

```shell
go build

export PORT=5000
export REDIRECT_URL=https://google.com
./redirect-app

# in another shell
curl -L http://localhost:5000
```
