#### 自定义组件列表： 

* TImageButton(四态图按钮控件)
* TGauge（图形进度条，支持几种样式）
* TLinkLabel（超链接标签）
* TRichEdit（富文本框）
* TXButton（主要是用颜色控制不同状态+背景图类的按钮）
* TMiniWebview（简易的浏览器组件）
----

* 1、安装Lazarus IDE，可以参考[安装Lazarus](https://gitee.com/ying32/govcl/wikis/pages?sort_id=2693253&doc_id=102420)，如果已安装则可以跳过此步。 
* 2、下载[liblcl源代码](https://github.com/ying32/liblcl)。
* 3、**新版本即将可以不用此步骤了** ~~根据[liblcl源代码](https://github.com/ying32/liblcl)中的"[编译引导](https://github.com/ying32/liblcl/blob/master/Compile.README.md)"编译一次liblcl(如果你安装的64的Lazarus就编译64位的liblcl，如果是安装的32位的Lazarus就编译32位的liblcl，**记得一定要编译一次liblcl（因为有个地方不想麻烦，所以才一定需要编译一次）**，否则会无法安装）。~~
* 4、在liblcl源码中找到`mylazpkgs.lpk`这个文件，双击打开(如果无法显示安装文件，则新建一个空项目后再双击这个lpk文件)。  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0829/113441_ba5e4e8a_118989.png "屏幕截图.png")   
* 5、点击 "Use -> Install"  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0829/113534_6fe31b7f_118989.png "屏幕截图.png")  
* 6、此时弹出重编译Lazarus对话框，点击“Yes"。  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0829/113608_4d6242de_118989.png "屏幕截图.png")   
* 7、等待编译完成后，在Lazarus IDE的组件选项卡中就会出现"GoVCL"选项卡。  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0829/113709_71548b66_118989.png "屏幕截图.png")   

**以上正常安装完后，就可以直接将组件拖放设计了。**     
![输入图片说明](https://images.gitee.com/uploads/images/2020/0829/115409_3c808c02_118989.png "屏幕截图.png")  