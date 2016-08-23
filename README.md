# Reproduce issue

Install glide, then:
```
glide install
go generate
```

It should fail with:
```
Package 'github.com/ory-am/fosite' not found on GOPATH
main.go:9: running "counterfeiter": exit status 1
```

If you do not have a version of fosite in your `GOPATH`.
If you install one with

```
go get github.com/ory-am/fosite
```

compilation succeeds, but with the potentially wrong interface.


