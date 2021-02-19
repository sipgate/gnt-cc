# gnt-cc

![](https://github.com/sipgate/gnt-cc/workflows/Build/badge.svg)

This is gnt-cc, an API + web frontend for one to many Ganeti virtualisation clusters.

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

`gnt-cc` looks for a file named `config.yml` in the current working directory. A sample configuration file ships with every release but is also available [here](api/config.example.yaml). If you are using the systemd unit file provided below, create a folder `/etc/gnt-cc` and touch `config.yml` in that folder. Start your new configuration as you would start any YAML document:

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
WorkingDirectory=/etc/gnt-cc

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
