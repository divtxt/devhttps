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

(1) Install:

```shell
brew install certbot caddy

# Clone the repo
cd ~/go/
git clone git@github.com:divtxt/devhttps.git
cd devhttps/

# Build
go build main.go
```


(2) Choose a subdomain for development e.g. `dev.yourapp.com`:

Use a domain you control (in place of "yourapp.com") - you must be able to create DNS entries!


(3) Configure DevHttps for your chosen domain:
  - Use your chosen subdomain (in place of "dev.yourapp.com")
  - Use the correct development port for your app (in place of 3000)
  - Follow instructions

```shell
./devhttps add dev.myapp.com 3000

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
https://dev.myapp.com/ → http://localhost:3000/
https://api.dev.myapp.com/ → http://localhost:8000/
```


### `devhttps add`

Add a development domain proxied to a local port.

```shell
$ devhttps add dev.myapp.com 3000
Saved: dev.myapp.com → port 3000

Your service is now available at:

https://dev.myapp.com

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

