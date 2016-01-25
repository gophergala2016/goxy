# Goxy

A forward proxy for Golang. Useful to intercept HTTP traffic on your phone or computer.

### Usage

Start the proxy server:

```
git clone git@github.com:montanaflynn/roxy.git
cd goxy
go run goxy.go -port 4444
```

Make requests through it:

```
curl -i localhost:4444 -H "host:gifs.com" -H "goxy-scheme-override:https"
```
