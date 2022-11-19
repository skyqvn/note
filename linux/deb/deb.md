# DEB包打包

[教程](https://blog.csdn.net/fengshengwei3/article/details/124271254)

结构:
```
|----DEBIAN
    |-------control
    |-------postinst(postinstallation)
    |-------postrm(postremove)
    |-------preinst(preinstallation)
    |-------prerm(preremove)
    |-------copyright(版权)
    |-------changlog(修订记录)
    |-------conffiles
|----etc
|----usr
|----opt
|----tmp
|----boot
    |-----initrd-vstools.img
```

## 打包命令:
```bash
dpkg -b dirnamexxx packagenamexxx.deb
```

## DEBIAN:

### control: 
> 必要属性:Package,Version,Architecture,Maintainer,Description  
> **末尾一定要空一行！！！**

```
Package:packagenamexxx #名称
Version:1.0.xxx #版本
Installed-Size: 1073741824xxx #安装后大小
Architecture:amd64 #CPU架构
Maintainer:www.xxx.com #维护者网站
Description:Your description.#描述
Maintainer:yournamexxx<email@yournamexxx.com> #制作此打包文件的作者(名称<邮箱>)
Depends:axxx,bxxx #软件所依赖的其他软件包和库文件。如果是依赖多个软件包和库文件，彼此之间采用逗号隔开
Pre-Depends:axxx #软件安装前必须安装、配置依赖性的软件包和库文件，它常常用于必须的预运行脚本需求
Recommends:axxx #表明推荐的安装的其他软件包和库文件
Suggests:axxx #建议安装的其他软件包和库文件
Priority:optional #软件对于系统的重要程度(如'required','standard','optional','extra')
Section:text #申明软件的类别(常见的有'utils','net','mail','text','x11'等)
Essential:no #申明是否是系统最基本的软件包(yse或者no)
Source:xxx #软件包的源代码名称
Priority:optional #声明这个包的优先级(大部分的时候使用optional(可选的))
Eseential:no #指该软件包是否是必须的(大部分的时候不是)
Conflicts:axxx #表示跟这个程序冲突的软件
Replaces:axxx #表明哪些软件包将被这个程序取代
```

### preinst:
> 处理安装前操作的脚本文件，按需添加

### prerm:
> 处理删除前操作的脚本文件，按需添加

### postinst:
> 处理安装后事件的脚本文件，比如创建快捷方式，删除备份文件等操作，按需添加

### postrm
> 处理删除后操作的脚本文件，按需添加

## 常用文件夹

### ./usr/share/icons
> 图标存放

### ./usr/share/applications
> .desktop文件存放

### ./usr/share/service
> .service文件存放

### ./usr/share/appnamexxx
> 应用所有文件存放

### ./bin/appnamexxx
> 可执行链接文件存放

安装deb包: dpkg -i mydeb.deb

卸载deb包: dpkg -r mysoftware

查看deb包是否安装: dpkg -s mysoftware

查看deb包文件内容: dpkg -c mydeb.deb

查看当前目录某个deb包的信息: dpkg --info mydeb.deb

解压deb包中所要安装的文件: dpkg -X mydeb.deb mydeb

解压deb包中DEBIAN目录下的文件(至少包含control文件): dpkg -e mydeb.deb mydeb/DEBIAN
