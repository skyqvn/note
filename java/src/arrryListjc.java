import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

public class arrryListjc {
	public static void main(String[] args) {
		ArrayList<Integer> arr = new ArrayList<>();
		arr.add(1);
		arr.add(4);
		arr.add(3);
		arr.add(0, 2);
		arr.set(1, 4);
		System.out.println(arr.get(1));
		arr.remove(1);
		Collections.sort(arr);
		ArrayList<Integer> arr2 = (ArrayList<Integer>) arr.subList(0, 2);
		System.out.println(arr2);
		System.out.println(arr.isEmpty());
		System.out.println(arr.size());
		Integer[] a = new Integer[5];
		a = (Integer[]) arr.toArray();
		System.out.println(Arrays.toString(a));
	}
}
