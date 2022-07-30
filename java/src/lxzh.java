import java.nio.charset.StandardCharsets;

public class lxzh {
	public static void main(String[] args) {
		int i = Integer.parseInt("200");
		String s = Integer.toString(20);
		float f = Float.parseFloat("200.9");
		byte[] b = s.getBytes(StandardCharsets.UTF_8);
		s = new String(b);
		System.out.println(b);
		System.out.println(s);
		System.out.println(i);
		System.out.println(s);
		System.out.println(f);
	}
}
