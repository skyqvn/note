import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;

public class iteratorjc {
	public static void main(String[] args) {
		int[] arr = new int[5];
		arr[0] = 10;
		arr[1] = 6;
		arr[2] = 3;
		arr[3] = 5;
		Iterator<Integer> it = Arrays.stream(arr).iterator();
		while (it.hasNext()) {
			System.out.println(it.next());
		}
		
		ArrayList<Integer> al = new ArrayList<>(5);
		al.add(1);
		al.add(2);
		Iterator<Integer> it2 = al.iterator();
		while (it2.hasNext()) {
			System.out.println(it2.next());
		}
	}
}
