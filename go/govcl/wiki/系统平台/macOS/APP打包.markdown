MacOS上的应用与linux、Windows的打包方式不一样。
需要生成MacOS下app专有格式目录及文件，而且不在app包中运行的会造成输入控件，如`TEdit, TMemo`等被转发到终端上。  

macapp包是一个快捷打包工具，只要导入此包即可生成MacOS下的app。

> 在main包中导入，必须第一个被导入的包，根据go包导入规则，同一包按文件名从小到大执行顺序，创建一个最小名起始的文件，如：`0.go`文件，在里面写入以下：
* 第一种方式，此包会每次将新编译的可执行文件复制到指定位置下再运行 
```go
package main

// 注意，编写此包的初衷只是为了开发时测试使用，正式的请不要使用此包，请手动将生成的可执行文件复制到对应文件夹。
import _ "github.com/ying32/govcl/pkgs/macapp"
```
 
* （未编写代码，后续做）第二种方式，此包只会在第一次运行时创建一个软链到指定目录下，链接指向当前可执行文件，然后运行
```go
package main

import _ "github.com/ying32/govcl/pkgs/macapp/debug"
```


**先将liblcl.dylib复制到`$GOPATH/bin`目录中，运行时打包工具会自动复制此目录的`liblcl.dylib`到目标的`xxx.app/Contents/MacOS/`目录。**    


MacOS App目录结构举例：  

```
govcl.app  
   |----Contents  
         |----MacOS  
              |----govcl         // 二进制文件  
              |----liblcl.dylib  // 核心UI库
         |----PkgInfo            //   
         |----info.plist         // 非常重要的文件  
         |----Resources          // 资源目录
              |----govcl.icns    // 图标   
```

info.plist文件格式：  

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>NSAppTransportSecurity</key>
	<dict>
		<key>NSAllowsArbitraryLoads</key>
		<true/>
	</dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>zh_CN<!--应用开发区域的语言--></string>
	<key>CFBundleExecutable</key>
	<string>govcl<!--应用程序名与可执行文件同名--></string>
	<key>CFBundleName</key>
	<string>govcl<!--应用程序名与可执行文件同名--></string>
	<key>CFBundleIdentifier</key>
	<string>ying32.govcl<!--这里要填, 像java包一样的命名--></string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleSignature</key>
	<string>proj</string>
	<key>CFBundleShortVersionString</key>
	<string>0.1</string>
	<key>CFBundleVersion</key>
	<string>1</string>
	<key>CSResourcesFileMapped</key>
	<true/>
	<key>CFBundleIconFile</key>
	<string>govcl.icns<!--这里是应用程序图标--></string>
	<key>CFBundleDocumentTypes</key>
	<array>
		<dict>
			<key>CFBundleTypeRole</key>
			<string>Viewer</string>
			<key>CFBundleTypeExtensions</key>
			<array>
				<string>*</string>
			</array>
			<key>CFBundleTypeOSTypes</key>
			<array>
				<string>fold</string>
				<string>disk</string>
				<string>****</string>
			</array>
		</dict>
	</array>
	<key>NSHighResolutionCapable</key>
	<true/>
        <key>NSHumanReadableCopyright</key>
	<string>copyright 2017-2018 ying32.com<!--这里是版权信息--></string>
</dict>
</plist>

```