# Go Auth0 Web Application

Auth0 Integration Samples for Go Web Applications.

## Configuration

The first thing you need to do is create an [Auth0](http://www.auth0.com/) application, and then modify the relevant configuration.

| Env | Description | Default Value |
|---|---|---|
| `APP_MODE` | prod or dev | `prod` |
| `HTTP_ADDR` | HTTP address | `:8080` |
| `AUTH0_CLIENT_ID` | Auth0 client ID | - |
| `AUTH0_DOMAIN` | Auth0 domain | - |
| `AUTH0_CLIENT_SECRET` | Auth0 client secret | - |
| `AUTH0_CALLBACK_URL` | Auth0 callback URL | `http://localhost:8080/callback` |

> Don't forget to modify the corresponding callback URL for Auth0 application.

Checkout [.env.example](.env.example) for details.

## Usage

```shell
$ go install github.com/razonyang/go-auth0-web-app
$ touch .env
$ go-auth0-web-app serve
```

## Structure

```text
/cmd                commands
/internal
    /core           core components
    /handlers       HTTP handler
    /middleware     middleware
/public             static resources
/views              templates
```
