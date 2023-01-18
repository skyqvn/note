import java.util.Arrays;

public class arrayjc {
	public static void main() {
		int[] a = {1, 2, 3, 4, 5, 6, 7, 8, 9, 0};
		
		int[] b = Arrays.copyOf(a, a.length);
		//4）type[] copyOfRange(type[] original, int from, int to)
		//这个方法与前面方法相似，但这个方法只复制 original 数组的 from 索引到 to 索引的元素。
		b = Arrays.copyOfRange(a, 0, 5);
		System.out.println(Arrays.equals(a, b));
	}
}
