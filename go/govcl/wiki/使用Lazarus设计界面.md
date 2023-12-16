
注：Lazarus默认的IDE是分离式的，如果不喜欢这种风格可以通过重新编译来达到，`适用于Lazarus 2.0.8`或之前的或者`2.2版本以后的`，Lazarus 2.0.10后出现编译错误，原因为因为`sparta_DockedFormEditor`这个引起的，Lazarus官方的说明是合并冲突造成，所以2.0.10中就不要用这个了，或者参考[Lazarus 2.0.10不能安装spata插件解决方法](https://gitee.com/ying32/govcl/wikis/pages?sort_id=2843377&doc_id=102420)中所示的方法解决这个问题，遇到这个错误要么完全卸载后重装，要么卸载`sparta_DockedFormEditor`插件后`Clean all`再重编译。   
方式为：
```
进入Package菜单

选择 Install/Uninstall Packages...

在Available for installation中双击

AnchorDocking

AnchorDockingDsgn

sparta_DockedFormEditor

然后 Save and rebuild IDE.  

```
注：如果是覆盖安装，需要先在Tools -> Configure "Build Lazarus"... 选中“Clean all”, 然后“Save Settings”，再开始上面的步骤。  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0617/182900_e4ce4352_118989.png "屏幕截图.png")

----
* Lazarus IDE区域说明   
![输入图片说明](https://images.gitee.com/uploads/images/2020/0605/204111_5dbe8ace_118989.png "屏幕截图.png")