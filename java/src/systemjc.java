public class systemjc {
	public static void main(String[] args) {
		System.gc();
		
		System.out.println(System.getProperty("java.version"));
		
		//属性名          属性说明
		//java.version	Java 运行时环境版本
		//java.home	    Java 安装目录
		//os.name	    操作系统的名称
		//os.version	操作系统的版本
		//user.name     用户的账户名称
		//user.home	    用户的主目录
		//user.dir	    用户的当前工作目录
		
		System.out.println("before exit");
		System.exit(0);
		System.out.println("after exit");
	}
}
