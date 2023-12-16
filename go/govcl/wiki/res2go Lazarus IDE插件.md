**注：本插件与res2go命令行工具属于互斥产品，二选一既可**   

----

**MacOS下要注意的事项：**   
一、Lazarus现在编译后的不会在原来的目录，而是会编译到`~/.lazarus/bin`目录中，需要复制编译好的lazarus二进制替换原来`/Library/Lazarus/`目录中的。  
二、因为MacOS一些特殊原因(环境变量问题)，目前在IDE中只能设计，编译的话要在res2go选项中设置go的路径，如果用了`GOPATH`可能编译不通过。  

----
**插件基于Lazarus 2.x，如不能编译，请确认是否版本一致(一般为向上兼容，即高版本可以编译)**   

* 1、群文件或者github下载插件源码包     
https://github.com/ying32/res2go-ide-plugin   

* 2、双击"res2goplugin.lpk"文件   
![输入图片说明](https://images.gitee.com/uploads/images/2020/0805/161443_77a7ce23_118989.png "屏幕截图.png")   

* 3、点击“Use”后面的下拉箭头，选择“Install”    
![输入图片说明](https://images.gitee.com/uploads/images/2020/0805/161513_e8c35359_118989.png "屏幕截图.png")  

* 4、此时弹出重编译Lazarus对话框，点击“Yes”，等待编译完成。  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0805/161534_0b093fc5_118989.png "屏幕截图.png")  

* 5、新建一个Lazarus工程，进入“Project -> Project Options... -> Project Options -> res2go”  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0830/182717_3d10ac0a_118989.png "屏幕截图.png")  
 
* 6、上面完成后，以后每次保存Lazarus项目则会生成对应的"go"文件。     
保存分为2种：  
* 1、Save   只会保存当前编辑器的，比如：lfm文件，pas文件。  
* 2、Save All  保存工程中所有已修改的，未修改的不会被保存，比如：lfm文件，pas文件，lpr文件。   

**注：1、如果未转换出对应的窗口，可能是原来的文件未被修改，所以不会有这个动作，只有被修改后才会成功转换，因为已经不再读取lfm文件进行转换了。2、res2go设置界面的语言会根据IDE当前的语言选择，IDE设置的中文则显示中文，反之显示英文。**  

----

#### 利用Lazarus IDE完成Windows可执行文件的图标、Manifest、版本信息修改（需要res2go插件`1.0.4及以上版本`），应用程序标题(`1.0.7以上版本`)。    
* 1、进入“Project -> Project Options... -> Project Options”    
 * 应用程序图标、Manifest修改，选择`Application`  
 ![输入图片说明](https://images.gitee.com/uploads/images/2020/0812/234841_1822c6d6_118989.png "屏幕截图.png")    

 * 版本信息修改，选择`Version Info`    
![输入图片说明](https://images.gitee.com/uploads/images/2020/0812/234958_2e59e91d_118989.png "屏幕截图.png")    
  
**以上修改后，下次`Save All`时会转换生成`386`及`amd64`的资源**    
![输入图片说明](https://images.gitee.com/uploads/images/2020/0812/235141_18cb5e61_118989.png "屏幕截图.png")    

----

#### 利用Lazarus IDE完成窗口顺序和指定某些窗口不在初始时创建   
* 1、进入“Project -> Project Options... -> Project Options -> Forms”    
![输入图片说明](https://images.gitee.com/uploads/images/2020/0819/125021_70bb665b_118989.png "屏幕截图.png")  

----

#### 利用Lazarus IDE完成Go的编译和运行。
* 方法一、点击图中IDE工具栏按钮   
![输入图片说明](https://images.gitee.com/uploads/images/2020/0820/163224_413cc70d_118989.png "屏幕截图.png")    

* 方法二、点击`MenuBar -> Run -> Complie`或者`MenuBar -> Run -> Build`等等（相应的快捷键也是生效的）。  

----

##### 运行时传入命令参数到目标
* 方法一、点击图中IDE工具栏按钮下拉箭头，选择`Run Paramemters...`在`Command line paramseters`中填入参数。  

![输入图片说明](https://images.gitee.com/uploads/images/2020/0820/163508_fe77a8d0_118989.png "屏幕截图.png")    

* 方法二、点击`MenuBar -> Run -> Run Paramemters...` 在`Command line paramseters`中填入参数。  

----

#### 利用Lazaurs IDE完成Windows下移除可执行文件的控制台窗口。  

`MenuBar -> Project -> Project Options... -> Compiler Options -> Config and Target`选中`Win32 gui application`，下次编译时就不会有命令行窗口了。   

![输入图片说明](https://images.gitee.com/uploads/images/2020/0830/160141_29771f83_118989.png "屏幕截图.png")  

----

#### 利用Lazarus IDE完成`GOOS`和`GOARCH`修改。  
`MenuBar -> Project -> Project Options... -> Compiler Options -> Config and Target -> Target platform`，如果都为默认则使用当前Go的默认值，反之使用这里面的。 

![输入图片说明](https://images.gitee.com/uploads/images/2020/0820/164107_6c826bcd_118989.png "屏幕截图.png")    

----

#### 利用Lazarus IDE完成去除调试信息和符号  

`MenuBar -> Project -> Project Options... -> Compiler Options -> Debugging`。   

![输入图片说明](https://images.gitee.com/uploads/images/2020/0830/160023_9b4b86ce_118989.png "屏幕截图.png")  

 