 **仅限于windows**   

**如果使用了res2go Lazarus IDE插件，则可以直接在IDE中设置图标，Manifest，版本信息，参考[res2go Lazarus IDE插件](https://gitee.com/ying32/govcl/wikis/pages?sort_id=2645001&doc_id=102420)，注：需要插件`1.0.4及以上版本和gnu工具链中的windres(一般情况下lazarus自带有这个)`**  

----

使用mingw64或mingw32中的windres.exe   


**位于`github.com\ying32\govcl\Tools\winRes`目录中有现成的模板。**  

---- 
#### 默认的资源包 

*注：一个go工程中，只能有一个.syso文件，如果采用自定义的，则不能再导入`winappres`包*  
 
```
  // 不带uac的
  import _ "github.com/ying32/govcl/pkgs/winappres"

  // 带uac的
  import _ "github.com/ying32/govcl/pkgs/winappres/uac"
```
-----  

* 准备好rc文件，rc文件格式为资源脚本每一行一个，格式为  资源名 资源格式 文件名。  

如  govcl.rc里面放置以下内容。    
```
1 MANIFEST "your.manifest"
MAINICON ICON "your.ico"

1 VERSIONINFO
FILEVERSION 1,1,1,0
PRODUCTVERSION 1,1,1,0
FILEOS 0x4
FILETYPE 0x1

BEGIN
    BLOCK "StringFileInfo"
    BEGIN
        BLOCK "040904B0"
        BEGIN
			VALUE "CompanyName", "ying32"
			VALUE "FileDescription", "govcl"
			VALUE "FileVersion", "1.1.1.1"
			VALUE "InternalName", "govcl"
			VALUE "LegalCopyright", "Copyright (C) ying32. All Rights Reserved."
			VALUE "OriginalFilename", "govcl.exe"
			VALUE "ProductName", "govcl"
			VALUE "ProductVersion", "1.1.1.1"
        END
    END
    BLOCK "VarFileInfo"
    BEGIN
            VALUE "Translation", 0x0409, 0x04B0
    END
END
```
注：MANIFEST 必须资源名必须为1， manifest文件格式见`方法二`中的。 如果资源中存在一个名为“MAINICON”的图标资源，则会默认加载为application图标。      

* 使用windres.exe编译为syso文件   

* x86
```bat
windres.exe -i app.rc -o defaultRes_windows_386.syso -F pe-i386
```  

* x64
```bat
windres.exe -i app.rc -o defaultRes_windows_amd64.syso -F pe-x86-64
```

注：在rc脚本中也可以放置其它格式的，具体网上找rc文件格式的说明。  


然后将生成的文件复制到你的工程目录下  

> manifest文件内容  
>  
> 如需要请求管理员权限，只需要将  
> level="asInvoker" 改为 level="requireAdministrator"  

```xml  

<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0" xmlns:asmv3="urn:schemas-microsoft-com:asm.v3">
  <asmv3:application>
    <asmv3:windowsSettings xmlns="http://schemas.microsoft.com/SMI/2005/WindowsSettings">
      <!--<dpiAware>True/PM</dpiAware>-->
    </asmv3:windowsSettings>
  </asmv3:application>
  <dependency>
    <dependentAssembly>
      <assemblyIdentity
        type="win32"
        name="Microsoft.Windows.Common-Controls"
        version="6.0.0.0"
        publicKeyToken="6595b64144ccf1df"
        language="*"
        processorArchitecture="*"/>
    </dependentAssembly>
  </dependency>
  <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
      <requestedPrivileges>
        <requestedExecutionLevel
          level="asInvoker"
          uiAccess="false"
        />
        </requestedPrivileges>
    </security>
  </trustInfo>
<compatibility xmlns="urn:schemas-microsoft-com:compatibility.v1"> 
	<application> 
		<!--The ID below indicates app support for Windows Vista -->
		<supportedOS Id="{e2011457-1546-43c5-a5fe-008deee3d3f0}"/> 
		<!--The ID below indicates app support for Windows 7 -->
		<supportedOS Id="{35138b9a-5d96-4fbd-8e2d-a2440225f93a}"/>
		<!--The ID below indicates app support for Windows 8 -->
		<supportedOS Id="{4a2f28e3-53b9-4441-ba9c-d69d4a4a6e38}"/>
		<!--The ID below indicates app support for Windows 8.1 -->
		<supportedOS Id="{1f676c76-80e1-4239-95bb-83d0f6d0da78}"/>
		<!--The ID below indicates app support for Windows 10 -->
		<supportedOS Id="{8e0f7a12-bfb3-4fe8-b9a5-48fd50a15a9a}"/>			
	</application> 
</compatibility>
</assembly>


``` 