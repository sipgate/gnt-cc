# gnt-cc
An API wrapper with local/ldap authentication around one or more Ganeti RAPI backends

## Howto build & run

- Init your swag
```
go get -u github.com/swaggo/swag/cmd/swag
swag init
```

- Run
```
go run main.go
```

## Frontend

### Development

```shell script
yarn && yarn serve
```

### Build (TODO)

```shell script
yarn build
```

## API Documentation

You can access the API documentation through `/swagger/index.html` on your gnt-cc instance, e.g. `https://gnt-cc.example.com/swagger/index.html` or `http://localhost:8080/swagger/index.html`.

## Features

### Implemented

- Authentication with local or LDAP backend
- Integrated API documentation (using swagger)
- Enumerate clusters, nodes, instances, and jobs
- Read cluster, instance and job details
- Provide a WebSocket proxy to enable authenticated access to the VNC or Spice port of an instance (needs a HTML5 VNC or Spice client)

### Planned

- implement the full RAPI functionallity

## Configuration

Please use the provided `config.yaml.example` as a template. The service expects to find a `config.yaml` file in its working directory. If you use systemd to start `gnt-cc`, please make sure to set the `WorkingDirectory` parameter to the folder in which `config.yaml` resides.

### Authentication Backends

Currently gnt-cc supports `builtin` or `ldap` as authentication backends. `builtin` uses a static and plaintext user/passwort list defined in `config.yaml` and should only be used for development. For production environments, use of the `ldap` backend is recommended. The `config.yaml.example` file contains templates for local users as well as LDAP configurations.
