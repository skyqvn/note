## 用此方法前，请确认`liblclbinres`与`govcl`代码是否一致，最好两者都更新为最新版本。  

----

## 此方法只支持windows及Linux，macOS因为没有必要，所以不需要这种方式

----

#### 此方法与内存中加载dll不一样，是首次运行时会在临时文件目录创建一个名为`liblcl/{crc32}/liblcl.{ext}`的文件   

----

### 在Windows上相对来说，也有被杀毒软件干掉的风险   

----

### 使用方法 ，如果使用了go.mod则可以忽略第一步，因为线上我已经上传了新版本。

* **第一步**

**两种方式来获取`liblclbinres`包**  

* 方法一  
**使用`Tools/genbinres`将liblcl预编译二进制压缩包转为资源**  
```bash
genbinres liblcl-2.0.3.zip
```
运行后会在`$GOPATH/src/github.com/ying32/liblclbinres`生成3个文件，分别是    
```bash
# linux 64bit可链接
liblcl_linux_amd64.go

# windows 32bit可链接
liblcl_windows_386.go

# windows 64bit可链接
liblcl_windows_amd64.go

```

* 方法二  
使用`go get`命令，如： `go get -u github.com/ying32/liblclbinres` 

* **第二步**  
编译时加上`-tags tempdll`

```bash
go build -i -tags tempdll
```

上面做完后，生成的可执行文件将`不再需要`带上`liblcl.dll`或者`liblcl.so`。  

**注：打包的文件都是经过zlib压缩过，然后生成.go文件的，解压时会校验`crc32`，如果不相等则删除，然后重新释放，如果释放错误，也会删除已经存在的错误文件。**     