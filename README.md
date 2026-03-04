# 🔒 DevHttps

**Easy https for local web development**

- Develop with `dev.yourapp.com` instead of `http://localhost:3000`
- Catch https issues in development instead of production
- No need to install/trsut a self-signed CA

DevHttps wraps [certbot](https://certbot.eff.org/) to generate certificates
and configures and runs [Caddy](https://caddyserver.com/) in https reverse proxy mode.


## Quickstart:

(1) Install:

```shell
brew install divtxt/tap/devhttps certbot caddy
```


(2) Configure DevHttps for your development:

```shell
devhttps add dev.myapp.com 3000
```

- Use a domain you control - you must be able to create DNS entries!
- Choose a subdomain (use it in place of "dev.yourapp.com")
- Use the correct development port for your app
- Follow instructions on DNS entries etc.


(3) Run the https server (wraps caddy run):

```shell
devhttps run
```

and enjoy development with https!

(don't forget to start your app)


## Notes

- For other commands, try: `devhttps help`
- Based on the [manual steps documented here](https://gist.github.com/divtxt/59e8c9ed6a4c7c90af7e73e687534b3b)
- Not implemented at this time:
  - remove command (edit the config file in ~/.devhttps/)
  - Any detection/warning about multiple apps on the same port
