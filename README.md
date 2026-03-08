# 🔒 DevHttps

**Easy https for local web development**

- Nice urls - eg  "dev.yourapp.com" instead of "localhost:3000"
- No need to install a custom CA certificate
- No need to disable security settings in development.
- Catch https issues in development.

DevHttps uses [certbot](https://certbot.eff.org/) to generate certificates and [Caddy](https://caddyserver.com/) as a https server.


## Quickstart:

(1) Install:

```shell
brew install divtxt/tap/devhttps certbot caddy
```


(2) Configure DevHttps for your development:

```shell
devhttps add dev.myapp.com 3000
```

- Use a subdomain on your application domain or another domain you control. You must be able to create DNS
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


DevHttps is based on this earlier document of manual steps:
[HTTPS in Development (early 2023)](https://gist.github.com/divtxt/59e8c9ed6a4c7c90af7e73e687534b3b). Initial development involved Claude Code and 15 hours of transit in an airport lounge. :P
