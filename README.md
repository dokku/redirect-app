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

## Example deployment with Dokku

On the Dokku host:

```shell
# Create the app
dokku apps:create xyz-redirect
# Set the redirect target URL
dokku config:set xyz-redirect REDIRECT_URL=https://example.com/xyz-page.html
# See above for other environment config variables
```

Locally:

```
# Check out this git repository
git clone https://github.com/dokku/redirect-app.git
# Add your dokku host remote address
git remote add live-server dokku@dokku-host.example.com:redirect-app
# Deploy to the dokku host
git push live-server
```

This will produce an application that redirects from `http://xyz-redirect.dokku-host.example.com/` to `https://example.com/xyz-page.html`.

See [Domain Configuration](https://dokku.com/docs/configuration/domains/#domain-configuration) to configure other domains for the application.
