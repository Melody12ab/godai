# GO语言版[百度AI][1]调用


[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

> go实现的对百度AI相关API调用

## Install

```
go get -u github.com/Melody12ab/godai
```

## Usage

**在base.go中修改你的APPID和APPSECRET**


```
import "github.com/Melody12ab/godai"

result, err := baiduai.GeneralBasic("test.jpg")

println(string(result))
```

# Author

**BigfaceMonster**

* <http://github.com/Melody12ab>
* <melody12ab@gmail.com>
* <https://www.jianshu.com/u/ff631089df1f>

[1]: http://ai.baidu.com/docs#/
