**此方法只适用于Windows和Linux。macOS因为其加载机制不一样，一般不要尝试去这样做。**   

> 例子参见：samples/customLibTest  

自定义加载需要根据go包导入规则来做。 

* 1、首先新建一个包，命名要比`github.com/ying32/govcl/vcl`包小，比如：`alib`
* 2、在`alib`中导入
```go

package alib

import (
	"github.com/ying32/govcl/pkgs/libname"
)

func init() {
	libname.LibName = "目标位置liblcl.dll"
}

```

* 3、在项目中导入`alib`包
```go
package main

// 此包不管在哪都必须要排在`github.com/ying32/govcl/vcl`前面，如果加载失败，则说明你的包名大于此。
import _ "github.com/ying32/govcl/samples/customLibTest/alib"
import "github.com/ying32/govcl/vcl"

type TMainForm struct {
	*vcl.TForm
}

var (
	mainForm *TMainForm
)

func main() {

	vcl.RunApp(&mainForm)
}

```

