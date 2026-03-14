# 🔒 DevHttps

**Easy https for local web development**

- Nice urls - eg  "dev.example.com" instead of "localhost:3000"
- Catch https issues in development.
- Stop disabling security settings in development.
- No need to install a custom CA.

DevHttps uses [certbot](https://certbot.eff.org/) to generate SSL certificates and runs [Caddy](https://caddyserver.com/) as a https server.


## Quickstart:

(1) Install:

```shell
brew install divtxt/tap/devhttps certbot caddy
```


(2) Generate certificate for your development domain:

```shell
devhttps add dev.example.com 3000
```

- Use a subdomain on a domain you control. You must be able to create DNS
entries.
- Use the correct development port for your app.


(3) Run the https server (wraps caddy run):

```shell
devhttps run dev.example.com 3000
```

- Use the correct development port for your app.
- Don't forget to start your application service!
- Use Ctrl-C to stop the http server.
- Enjoy development with https!


## Add to your project's git repository:

It will be useful to capture the caddy configuration in your project source code.
This also lets you customize the caddy configuration -
for example, by having caddy serve static assets directly.
(see [Caddyfile](https://caddyserver.com/docs/caddyfile) documentation)


Use the following steps:

(1) Generate caddy configuration and add it to your git repo:

```shell
devhttps caddyfile dev.example.com 3000 > ./Caddyfile

git add Caddyfile

git commit -m "Development https using caddy"
```

(2) Run caddy directly:

```shell
caddy run
```

(3) Optional: share generated certificates with your team

It is bad security practice to add the certificate to your git repo,
so you will have to share the relevant certificate files by a secure mechanism.
(check the Caddyfile for the certificate files to be shared)


## Notes

- For more commands, use: `devhttps help`
- Genearate files are stored under `~/.devhttps/`.
- There is no detection/warning about multiple apps on the same port
- How it works: automates the steps in
[this document](https://gist.github.com/divtxt/59e8c9ed6a4c7c90af7e73e687534b3b).
