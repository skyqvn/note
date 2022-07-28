interface i {
	void print(String s);
}

interface j {
	void print(String s);
}

class c implements i, j {
	public void print(String s) {
		System.out.println(s);
	}
}

class b extends c {

}

public class interfacejc {
	public static void main(String[] args) {
		c cv = new c();
		b bv = new b();
		System.out.println(cv.getClass().getName());
		System.out.println(bv.getClass().getSuperclass().getName());
		System.out.println(bv.getClass().getInterfaces().length);
		System.out.println(cv.getClass().getInterfaces()[0]);
		System.out.println(cv.getClass().getInterfaces()[1]);
		try {
			System.out.println(bv.getClass().getInterfaces()[0]);
		} catch (Exception e) {
			System.out.println("b没有实现接口");
		}
		
		i iv = new c();
		cv = (c) iv;
		iv = bv;
		iv.print("hello world");
		Object ov = cv;
		if (ov instanceof c) {
			cv = (c) ov;
		}
	}
}
