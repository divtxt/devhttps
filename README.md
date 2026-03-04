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
brew install devhttps certbot caddy
```


(2) Configure DevHttps for your application:
  - Use a domain you control - you must be able to create DNS entries!
  - Choose a subdomain (use it in place of "dev.yourapp.com")
  - Use the correct development port for your app (in place of 3000)
  - Run this command and follow instructions

```shell
./devhttps add dev.myapp.com 3000
```


(3) Run the https server (wraps caddy run):

```shell
./devhttps run
```

(and don't forget to start your app)


## Usage


### `devhttps add`

Add a development domain proxied to a local port.

```shell
$ devhttps add dev.myapp.com 3000
...

✓ Domain added. Run https server using:  devhttps run
```


### `devhttps run`

Run Caddy https reverse proxy for the configured domains.

```shell
$ devhttps run

Configured domains:
  ✓ https://dev.myapp.com → :3000  (cert: VALID, 89 days left)


Starting Caddy...

...
```


### `devhttps show`

Show all configured development domains.

```shell
$ devhttps show
https://dev.myapp.com/ → http://localhost:3000/
```



### `devhttps check`

Run various checks. Use this to diagnose issues.

```shell
$ devhttps check
Tools:
  ✓ certbot: certbot 5.3.1 (/opt/homebrew/bin/certbot)
  ✓ caddy: v2.10.2 h1:g/gTYjGMD0dec+UgMw8SnfmJ3I9+M2TdvoRL/Ovu6U8= (/opt/homebrew/bin/caddy)

Config:
  ✓ /Users/div/.devhttps/config.json

Configured domains:
  ✓ https://dev.myapp.com → :3000  (cert: VALID, 89 days left)


(to edit port or renew certificates, use add command)

```

