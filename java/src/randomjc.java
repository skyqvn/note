import java.util.Date;
import java.util.Random;

public class randomjc {
	public static void main(String[] args) {
		Random r = new Random();
		r.setSeed(new Date().getTime());
		System.out.println(r.nextInt(0, 100));
		System.out.println(r.nextBoolean());
	}
}
