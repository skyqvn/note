import java.util.Calendar;
import java.util.Date;

public class timejc {
	public static void main(String[] args) {
		Date d = new Date();
		System.out.println(d.getTime());
		System.out.println(d.after(new Date(0)));
		System.out.println(d.getTime() - new Date().getTime());
		System.out.println(d.toString());
		d.setTime(10000);
		
		Calendar c = Calendar.getInstance();
		c.setTime(d);
		c.set(2022, 7, 28);
		System.out.println(c.compareTo(Calendar.getInstance()));
		System.out.println(c.get(Calendar.HOUR));
		System.out.println(c.getTime());
		c.setTimeInMillis(d.getTime());
	}
}
