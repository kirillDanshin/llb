# llb [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go)

### What the f--k it is?
It's a very simple but quick backend for proxy servers. You can setup redirect to your main domain or just show `HTTP/1.1 404 Not Found` with zero memory allocation and very fast response. Also useful when you need to serve many redirects or not found errors but don't want any overheads.

### Quick start

```sh
git clone git@github.com:kirillDanshin/llb.git
cd llb
go build
# run in background with default ports 
# but custom redirect destination
((./llb -redirDest="https://google.com/" &)&)&
```

Than setup your proxy to send any invalid requests like invalid host etc to this backend.
