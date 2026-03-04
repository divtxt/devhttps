# 🔒 DevHttps

**Easy https for local web development**

---

Develop using `https://dev.yourapp.com` instead of `http://localhost:3000` !

DevHttps automatically generates a https certificate,  and configures and runs
[Caddy](https://caddyserver.com/) with https reverse proxy.

**Benefits:**

- Develop & test with https e.g. secure cookies
- Nice URLs: e.g. `https://dev.yourapp.com` instead of `http://localhost:3000`
- No need to disable security checks for non-https development
- Catch https related issues before deploying to production


## Quickstart:

(1) Install DevHttps and Caddy webserver:

```shell
brew install devhttps caddy
```


(2) Choose a subdomain for development e.g. `dev.yourapp.com`:

- Use a domain you control - you must be able to edit DNS entries!


(3) Configure DevHttps for your chosen domain:
  - Use your chosen subdomain
  - Use the development port for your app (e.g. 3000 for node)
  - Create the 2 DNS entries specified as instructed

```shell
devhttps add --start dev.myapp.com 3000

# (create DNS entries as shown by the command)
```


(4) Enjoy https in development:

```shell
open "https://dev.myapp.com/"
```

(Don't forget to start your app e.g. "npm run dev")


## Usage


### `devhttps show`

Show all configured development domains.

```shell
$ devhttps show
dev.myapp.com → port 3000
api.dev.myapp.com → port 8000
```


### `devhttps add`

Add a development domain proxied to a local port.

```shell
$ devhttps add dev.myapp.com 3000
Saved: dev.myapp.com → port 3000
```


### `devhttps check`

Run various checks. Use this to diagnose issues.

```shell
$ devhttps check
certbot: found (/usr/local/bin/certbot)
certbot: version OK (certbot 5.0.0)
caddy: found (/usr/local/bin/caddy)
caddy: version OK (v2.8.4 h1:...)
```

