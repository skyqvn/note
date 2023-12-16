## 不推荐用此方法，将来可能会删除这些代码。推荐使用[将dll打包到可执行文件中](https://gitee.com/ying32/govcl/wikis/pages?sort_id=2364531&doc_id=102420)

## 用此方法前，请确认`liblcl`与`govcl`代码是否一致，最好两者都更新为最新版本。  

----


**注：此种方法目前只支持Windows 32位，而且有被杀毒软件杀的风险，而且被编译到exe内部的liblcl不能使用upx进行压缩、不能使用upx进行压缩、不能使用upx进行压缩。**    

#### 从内存中加载dll的几个步骤：

* 一、资源脚本准备(*.rc文件)  

```
1 MANIFEST "exe.manifest"
MAINICON ICON "icon.ico"
GOVCLLIB RCDATA "liblcl.dll"
```

以上定义了3个资源，manifest文件、程序图标文件、UI库。  manifest文件内容可以[从这里获取](https://gitee.com/ying32/govcl/wikis/pages?sort_id=410058&doc_id=102420)。   

* 二、编译资源文件

编译资源需要用到mingw中的winres.exe  

```bash
windres.exe -i app.rc -o defaultRes_windows_386.syso -F pe-i386
```

* 三、在go的编译命令行中加入tags参数`memorydll`，如：  

```
go build -i -tags memorydll
```

以上做完后发布的时候就不需要再带上liblcl了。  