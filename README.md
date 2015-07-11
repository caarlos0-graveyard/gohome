# gohome [![Build Status](https://drone.io/github.com/caarlos0/gohome/status.png)](https://drone.io/github.com/caarlos0/gohome/latest)
Easily get cache and config folders for your app according to each OS spec

## Usage

```go
import "github.com/caarlos0/gohome"

const appName = "my-app"

func main()  {
  config := gohome.Config(appName)
  cache := gohome.Cache(appName)
  // ...
}
```
