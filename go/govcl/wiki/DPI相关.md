**目前没有2k及以上屏幕，所以这方面暂时处于搁置状态，所以在windows上的manifest中默认是关掉dpi的，可以自己开启，**  

原Lazarus中相关的api有`Application.SetScale` 在`vcl.Application.Initialize()`之前调用下，设置为true，windows下再配合manifest启用dpi的。 

对于Windows上可以通过manifest文件暂时控制下(`如果Windows下关闭了dpi则UI在高分辨率上会模糊`)。  

* manifest文件
```xml
<dpiAware>True/PM</dpiAware>
```
或者,适用于高版本的windows
```xml
<dpiAware>PerMonitorV2</dpiAware>
```
取值参考：
![输入图片说明](https://images.gitee.com/uploads/images/2020/0523/094406_699946ab_118989.png "屏幕截图.png")  
![输入图片说明](https://images.gitee.com/uploads/images/2020/0523/094532_0696a1da_118989.png "屏幕截图.png")