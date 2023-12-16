### liblcl  

* `liblcl`基于Lazarus 2.0.x版本制作，可以用作跨平台，有些功能或者性能略低于libvcl。   

### libvcl（已弃用）  

* `libvcl`基于Delphi 10.2.x版本制作，只能应用于Windows下，稳定性较好。   

### 何时使用libvcl？何时使用liblcl?    

**这个其实主要看自己的需求了。**  

* 不跨平台：如果没有跨平台考虑的话可以使用libvcl或者对UI要求不太高的话可以选择使用liblcl。  
* 需跨平台：如果是一开始就有跨平台的考虑，还是直接用liblcl较好，使用Lazarus IDE设计界面 + liblcl组合，这样所以平台统一性较好。  