**gitee上的govcl仓库用于同步github上的仓库，留作备份，以防止无法访问github时使用**      
**千万不要**直接go get本gitee上的项目，如要引用gitee上的仓库则按以下步骤进行：   

* 1、在“%GOPATH%\src”目录中建立“github.com\ying32”目录。
* 2、然后再在ying32目录中使用`git clone https://gitee.com/ying32/govcl`拉取本仓库。

如命令行：
```bat
rem 假设这里GOPATH中只有一个目录
mkdir %GOPATH%\src\github.com\ying32

rem 切换到新建的目录
cd %GOPATH%\src\github.com\ying32

rem 拉取代码
git clone https://gitee.com/ying32/govcl
```

平时更新的时候只需要进入 `%GOPATH%\src\github.com\ying32\govcl` 
```bat
git pull
```