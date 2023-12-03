# Syso



## 安装

```bash
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
```



### 默认配置

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
	<assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="SomeFunkyNameHere" type="win32"/>
	<dependency>
		<dependentAssembly>
			<assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
		</dependentAssembly>
	</dependency>
	<application xmlns="urn:schemas-microsoft-com:asm.v3">
		<windowsSettings>
			<dpiAwareness xmlns="http://schemas.microsoft.com/SMI/2016/WindowsSettings">PerMonitorV2, PerMonitor</dpiAwareness>
			<dpiAware xmlns="http://schemas.microsoft.com/SMI/2005/WindowsSettings">True</dpiAware>
		</windowsSettings>
	</application>
</assembly>

```



## 管理员权限编译

xxx.exe.manifest

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
	<!-- ………其他配置……… -->
	<trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
		<security>
			<requestedPrivileges>
				<requestedExecutionLevel level="requireAdministrator"/>
			</requestedPrivileges>
		</security>
	</trustInfo>
</assembly>
```


## 应用信息
versioninfo.json

```json
{
    "FixedFileInfo": {
        "FileVersion": {//文件版本
            "Major": 1,//主版本号
            "Minor": 0,//次版本号（一般奇数表示测试版，偶数表示生产版）
            "Patch": 0,//修订号
            "Build": 0//编译次数
        },
        "ProductVersion": {//产品版本
            "Major": 1,//主版本号
            "Minor": 0,//次版本号（一般奇数表示测试版，偶数表示生产版）
            "Patch": 0,//修订号
            "Build": 0//编译次数
        },
        "FileFlagsMask": "3f",
        "FileFlags ": "00",
        "FileOS": "040004",
        "FileType": "01",
        "FileSubType": "00"
    },
    "StringFileInfo": {
        "Comments": "这是程序描述",
        "CompanyName": "这是公司名称",
        "FileDescription": "这是文件说明",
        "FileVersion": "v1.0.0.0",
        "InternalName": "这是程序内部名称",
        "LegalCopyright": "Copyright (c) 2021 XCGUI",
        "LegalTrademarks": "",
        "OriginalFilename": "这是原始文件名",
        "PrivateBuild": "",
        "ProductName": "这是产品名称",
        "ProductVersion": "v1.0.0.0",
        "SpecialBuild": ""
    },
    "VarFileInfo": {
        "Translation": {
            "LangID": "0804",
            "CharsetID": "04B0"
        }
    },
    "IconPath": "icon.ico",//图标
    "ManifestPath": ""//manifest文件位置
}
```



## go程序更改

```go
//go:generate goversioninfo
package main

import (
	"fmt"
)

func main() {
	fmt.Println("go编译带windows版本信息的exe程序")
}
```



## 编译

```bash
go generate
go build
```