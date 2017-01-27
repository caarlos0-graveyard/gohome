# gohome [![Build Status](https://travis-ci.org/caarlos0/gohome.svg?branch=master)](https://travis-ci.org/caarlos0/gohome) [![Coverage Status](https://coveralls.io/repos/caarlos0/gohome/badge.svg?branch=master&service=github)](https://coveralls.io/github/caarlos0/gohome?branch=master) [![](https://godoc.org/github.com/caarlos0/gohome?status.svg)](http://godoc.org/github.com/caarlos0/gohome) [![](http://goreportcard.com/badge/caarlos0/gohome)](http://goreportcard.com/report/caarlos0/gohome) [![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/caarlos0)

Easily get cache and config folders for your app according to each OS spec

## Usage

```go
import "http://gopkg.in/caarlos0/gohome.v2"

const appName = "my-app"

func main()  {
  config := gohome.Config(appName) // gives you the right config folder for the current OS
  cache  := gohome.Cache(appName)  // gives you the right cache folder for the current OS
  // ...
}
```
