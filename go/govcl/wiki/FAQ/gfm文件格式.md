#### 文件格式
gfm文件格式目前分为2种

* 1、早期的为自定义加密的格式（目前已经废弃），res2go已经删除此选项。
* 2、为原Lazarus的二进制资源格式，res2go默认输出格式。  

#### 文件作用

* 1、如果是使用Lazarus设计器+res2go生成的，且不强制指定从文件加载gfm文件，则没啥用。   
* 2、提供给GoVCL专用设计器的二次编辑文件，即gfm格式。  

#### 将gfm文件还原为lfm文件   

**govcl 2.0.5版本源码中的res2go提供了一个隐藏功能:**
**2.0.5版本中需要修改下res2go的源码，如下图，改为这样：**    
![输入图片说明](https://images.gitee.com/uploads/images/2020/0811/115844_a6d1f105_118989.png "屏幕截图.png")   

```bash
res2go -tolfm # 将当前目录下的gfm文件全部转为lfm文件
```  
##### 如何提取原来lfm中的组件到设计器中   

* 1、用文本编辑器打开lfm文件，将第一个`object`关键字开始到第二个`object`关键字之前的内容全部删除，如下图，删除红框中的：    
![输入图片说明](https://images.gitee.com/uploads/images/2020/0811/115022_750eda44_118989.png "屏幕截图.png")   

* 2、删除lfm文件中最后一个`end`关键字，如下图：  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0811/115130_788273ba_118989.png "屏幕截图.png")  

* 3、检查下首尾有没有`空行`，如果有需要删除掉，保证这个文件中没有`空行`，如下图：  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0811/115244_f2235fd5_118989.png "屏幕截图.png")    

* 4、`Ctrl+A`全选内容，`Ctrl+C`复制内容，然后回到Lazarus设计器中，新建一个空白的窗口, 单击空白窗口，让窗口获得焦点，然后`Ctrl+V`粘贴，此时因为有可能用`GoVCL IDE`生成的gfm文件中有些组件在Lazarus IDE中是没有的，所以选择`Continue loading`,如下图：  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0811/115607_b0b5256d_118989.png "屏幕截图.png")   

* 5、完成后如需要调整窗口的属性，可以参考之前删除掉的那段设置下新窗口的属性。   
![输入图片说明](https://images.gitee.com/uploads/images/2020/0811/115728_413adfca_118989.png "屏幕截图.png")  
 
