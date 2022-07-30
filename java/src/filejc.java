import java.io.*;
import java.nio.charset.StandardCharsets;
import java.util.Arrays;

public class filejc {
	public static void main(String[] args) {
		{
			File in = new File("C:/java");
			if (!in.mkdirs()) {
				System.out.println("无法创建目录");
			}
		}
		
		{
			File in = new File("C:/java/hello");
			try {
				in.createNewFile();
			} catch (IOException e) {
				System.out.println(e.toString());
			}
		}
		
		try {
			InputStream in = new FileInputStream("C:/java/hello");
			byte[] b = new byte[1];
			while (true) {
				try {
					if (in.read(b) == -1) {
						break;
					}
				} catch (IOException e) {
					break;
				}
			}
			System.out.println(new String(b));
		} catch (FileNotFoundException e) {
			System.out.println(e.toString());
		}
		
		try {
			OutputStream in = new FileOutputStream("C:/java/hello");
			try {
				in.write("你好".getBytes(StandardCharsets.UTF_8));
			} catch (IOException e) {
				System.out.println(e.toString());
			}
		} catch (FileNotFoundException e) {
			System.out.println(e.toString());
		}
		
		File f1 = new File("C:/java/hello");
		try {
			InputStream in = new FileInputStream(f1);
		} catch (FileNotFoundException e) {
			System.out.println(e.toString());
		}
		
		try {
			InputStream f2 = new FileInputStream("C:/java/hello");
		} catch (FileNotFoundException e) {
			System.out.println(e.toString());
		}
	}
}
