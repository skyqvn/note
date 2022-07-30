import java.util.ArrayList;

public class fanxingjc {
	public static void main(String[] args) {
		fanxing<Integer> fx = new fanxing<>();
		fx.print1(new Integer[10]);
		fx.print2(new ArrayList<>());
	}
}

class fanxing<T extends Integer> {
	public void print1(T[] tList) {
		for (T t : tList) {
			System.out.println(t);
		}
		System.out.println("\n");
	}
	
	public void print2(ArrayList<?> tList) {
		for (int i = 0; i < tList.size(); i++) {
			System.out.println(tList.get(i));
		}
		System.out.println("\n");
	}
}
