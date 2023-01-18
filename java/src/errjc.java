import java.io.IOException;
import java.nio.charset.StandardCharsets;

public class errjc {
	public static void main(String[] args) {
		try {
			say("hello");
			say("");
		} catch (MyError myError) {
			myError.printStackTrace();
		}
	}
	
	public static void say(String s) throws MyError {
		if (s.length() == 0) {
			throw new MyError();
		}
		try {
			System.out.write(s.getBytes(StandardCharsets.UTF_8));
		} catch (IOException e) {
			e.printStackTrace();
		} finally {
			System.out.println("say ok");
		}
	}
}

class MyError extends Throwable {
	public void printStackTrace() {
		System.err.println("MyError");
	}
}
