import java.util.Arrays;
import java.util.Collections;
import java.util.LinkedList;

public class linkedliatjc {
	public static void main(String[] args) {
		LinkedList<Integer> ll = new LinkedList<>();
		ll.add(1);
		ll.add(4);
		ll.add(3);
		ll.set(1, 4);
		System.out.println(ll.get(1));
		ll.remove(1);
		Collections.sort(ll);
		LinkedList<Integer> ll2 = (LinkedList<Integer>) ll.subList(0, 2);
		System.out.println(ll2);
		System.out.println(ll.isEmpty());
		System.out.println(ll.size());
		Integer[] a = new Integer[5];
		a = (Integer[]) ll.toArray();
		System.out.println(Arrays.toString(a));
	}
}
