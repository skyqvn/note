因为Lazarus 2.0.10中的fpc自带有sparta中的泛型，造成与之冲突，所以官方没有合并且没有修改。  

* 进入Lazarus安装目录，比如：`F:\lazarus\components\sparta\mdi`   
  
* 双击`sparta_mdi.lpk`文件。    
![输入图片说明](https://images.gitee.com/uploads/images/2020/0910/211112_544535e5_118989.png "屏幕截图.png")    

* 右键删除`sparta_Generices`包。  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0910/211245_9c8ab354_118989.png "屏幕截图.png")    
 
* 点击`Compile`测试是否正常编译。    
* 编译正常则回到原来的 [使用Lazarus设计界面](https://gitee.com/ying32/govcl/wikis/pages?sort_id=2294674&doc_id=102420)。     

**注：如原来几个安装包都安装了还是没有出现，则全部卸载三个插件编译再添加再编译。**