# 🔒 DevHttps

**Easy https for local web development**

- Nice urls - eg  "dev.example.com" instead of "localhost:3000"
- Catch https issues in development.
- Stop disabling security settings in development.
- No need to install a custom CA.

DevHttps uses [certbot](https://certbot.eff.org/) to generate certificates and [Caddy](https://caddyserver.com/) as a https server.


## Quickstart:

(1) Install:

```shell
brew install divtxt/tap/devhttps certbot caddy
```


(2) Configure DevHttps for your development:

```shell
devhttps add dev.example.com 3000
```

- Use a subdomain on a domain you control. You must be able to create DNS
entries.
- Use the correct development port for your app.
- Create DNS entries as instructed by the command.


(3) Run the https server (wraps caddy run):

```shell
devhttps run
```

(and don't forget to start your application service)

Enjoy local development with https!


----

## Notes

- Config & certificates are stored under `~/.devhttps/`.
- For more commands, use: `devhttps help`
- Not implemented at this time:
  - remove command (edit the config file in ~/.devhttps/)
  - Any detection/warning about multiple apps on the same port
- How it works: automates the steps in
[this document](https://gist.github.com/divtxt/59e8c9ed6a4c7c90af7e73e687534b3b).
