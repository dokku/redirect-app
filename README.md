# redirect-app

An app whose sole purpose is to redirect requests to a given url.

## Requirements

- Go 1.19+

## Usage

```shell
go build

# set some env vars
export REDIRECT_URL=https://google.com

# optionally define an alternative PORT (default=5000)
export PORT=3000

# optionally define an alternative redirect status code (default=302)
export REDIRECT_STATUS_CODE=301

# start the server
./redirect-app

# in another shell
curl -L http://localhost:5000
```
