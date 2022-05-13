<p align="center">
  <img alt="gnt-cc logo" src="https://user-images.githubusercontent.com/22923578/123423959-7ba85a80-d5c0-11eb-8852-06ab161d4e3f.png" width="160"/>
</p>
<h2 align="center">gnt-cc - a frontend for Ganeti clusters</h2>

![](https://github.com/sipgate/gnt-cc/workflows/Build/badge.svg)

gnt-cc is a web-based frontend for [Ganeti](https://github.com/ganeti/ganeti/) clusters.

![gnt-cc-screenshot-light](https://user-images.githubusercontent.com/22923578/123422533-a691af00-d5be-11eb-909a-a884b7e0c6bc.png)

# Features

This is an incomplete list of available/upcoming features.

- [x] Builtin and LDAP based authentication
- [x] Multiple Ganeti clusters
- [x] Overview dashboard per cluster
- [x] List nodes, instances, jobs
- [x] Instance details
- [x] Track job status in frontend
- [x] Start/stop/restart/migrate/failover instances
- [x] VNC web console
- [ ] Search resources (nodes, instances) across all clusters
- [ ] Spice web console
- [ ] Create instances
- [ ] Configure instances
- [ ] Cluster details
- [ ] Configure clusters
- [ ] Node details

# Development

For build/test/development information, please check the relevant READMEs for [frontend](web/README.md) and [backend](api/README.md).

## Git Hooks

We provide git hooks to run linters and tests before pushing. To install them use:
```
./bin/setup-git-hooks.sh
```

To remove the hooks use:
```
./bin/remove-git-hooks.sh
```

# Installation

Please check the [Github release section](https://github.com/sipgate/gnt-cc/releases) for the latest version of gnt-cc. Both the API and the frontend are contained in a single Go binary. All you have to do really boils down to these steps:

- Download the binary
- Create a configuration file
- Additionally configure your favorite reverse proxy to add TLS (apache, nginx, haproxy etc.)
- Run the `gnt-cc` binary

## Create a configuration file

By default `gnt-cc` looks for a file named `config.yaml` in the current working directory or in `/etc/gnt-cc`. A sample configuration file ships with every release but is also available [here](api/config.example.yaml). If you're using the .deb file, an example config file will also be placed in `/etc/gnt-cc`. When creating a config file, start it as you would do with any YAML document:

```yaml
---
```

We need to decide where to bind to/listen for requests. We highly recommend running a reverse proxy for TLS in front of `gnt-cc`, so let's bind to localhost only:
```yaml
bind: 127.0.0.1
port: 8000
```

We want to see any errors/warnings, so set the log level accordingly:
```yaml
logLevel: warning
```
Other accepted values are: `debug`, `info`, `error`, `fatal`.

We need to specify one or more Ganeti clusters with their respective RAPI endpoints. The `name` Parameter should only consist of upper- and lowercase letters, dashes or underscores:
```yaml
clusters:
  - name: "test-cluster"
    hostname: "test-cluster.example.com"
    port: 5080
    description: "Ganeti Test Cluster"
    username: "gnt-cc"
    password: "somepassword"
    ssl: True
  - name: "production-cluster"
    hostname: "prod-cluster.example.com"
    port: 5080
    description: "Ganeti Production Cluster"
    username: "gnt-cc"
    password: "somepassword"
    ssl: True
```

We use JSON Web Tokens (JWT) for authentication and need to generate a random string as signing key and set an expiry timeout. The latter uses suffixes like s(econds), m(inutes) or h(ours):
```yaml
jwtSigningKey: "RaNdOmStRiNg123456789"
jwtExpire: "60m"
```

To obtain the aforementioned JWT, we need to authenticate against one of the available authentication providers. Currently `builtin` and `ldap` are supported. If you want a quick and easy start, choose `builtin`.

### Authentication Method 'builtin'

This method authenticates against a list of usernames/passwords in plaintext stored in the configuration file. Add the following to your configuration:
```yaml
authenticationMethod: "builtin"
users:
  - username: "maya"
    password: "mayas-plaintext-password"
  - username: "john"
    password: "johns-plaintext-password"
```

This is **not** recommended for production setups.

### Authentication Method 'ldap'

This method authenticates against a LDAP server (tested with OpenLDAP). Add the following to your configuration:
```yaml
authenticationMethod: "ldap"
ldapConfig:
  host: "my.ldap.server.tld"
  port: 389
  skipCertificateVerify: false
  userFilter: "(&(objectClass=posixAccount)(uid=%s))"
  groupFilter: "(&(objectClass=posixGroup)(cn=someGroupName)(memberUid=%s))"
  baseDn: "dc=domain,dc=org"
```

Please adapt the user and group search filters and base DN to your LDAP schema. You can use e.g. the `ldapsearch` tool to test filters on the commandline. `%s` will be substituted by `gnt-cc` with the username to be authenticated. If your LDAP server uses a self-signed TLS certificate (or the CA is unknown to your local CA trust store) you may set `skipCertificateVerify` to `true`.

## Create a systemd service

You can use systemd to run `gnt-cc`. Please create the file `/etc/systemd/system/gnt-cc.service` with the following content:
```
[Unit]
Description=gnt-cc API server

[Service]
Type=simple
ExecStart=/usr/local/bin/gnt-cc

[Install]
WantedBy=multi-user.target
```

This assumes you have placed the configuration file in the folder `/etc/gnt-cc` and the binary is located at `/usr/local/bin/gnt-cc`. Adapt to your environment as required.

Once the file is in place, tell systemd to re-read its configuration and enable/start the service:
```shell
systemctl daemon-reload
systemctl enable gnt-cc
systemctl start gnt-cc
```

## Example reverse proxy config with apache

This is a minimalistic example configuration to configure apache as a reverse proxy for `gnt-cc`. This also enables proxying of websocket connections which is required if you want to use the included web VNC console.

```
<VirtualHost 1.2.3.4:443>
	ServerAdmin webmaster@example.com

	SSLEngine on
  SSLCertificateFile /etc/ssl/certs/gnt-cc.example.com-cert.pem
  SSLCertificateKeyFile /etc/ssl/private/gnt-cc.example.com-key.pem

	ServerName gnt-cc.example.com
	RequestHeader set X-Forwarded-Proto "https"

	ProxyRequests Off
	ProxyPreserveHost On
	ProxyPass / http://localhost:8000/
	ProxyPassReverse / http://localhost:8000/

	ErrorLog  /var/log/apache2/gnt-cc.example.com_error.log
	CustomLog /var/log/apache2/gnt-cc.example.com_access.log

	RewriteEngine on
	RewriteCond %{HTTP:Upgrade} websocket [NC]
	RewriteCond %{HTTP:Connection} upgrade [NC]
	RewriteRule ^/?(.*) "ws://localhost:8000/$1" [P,L]
</VirtualHost>
```

The above configuration requires the apache modules `ssl`, `proxy`, `proxy_http` and `proxy_wstunnel` to be enabled.

