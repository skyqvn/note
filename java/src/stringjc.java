public class stringjc {
	public static void main(String[] args) {
		String s = "hello";
		String s2 = "JAVA";
		System.out.println(s + " " + s2);
		System.out.println(s.length());
		System.out.println(s.toUpperCase());
		System.out.println(s2.toLowerCase());
		String s3 = "  hello ";
		System.out.println(s3.trim());
		System.out.println(s3.replace(" ", ""));
		System.out.println(s3.replaceAll("[ho]", "q"));
		System.out.println(s3.replaceFirst(" ", ""));
		String s4 = "hello world,hello java";
		System.out.println(s4.substring(6));
		System.out.println(s4.substring(0, 5));
		System.out.println(s.charAt(1));
		System.out.println(s.split(",")[0]);
		System.out.println(s.split("[ ,]")[0]);
		System.out.println(new StringBuilder(s4).replace(0, 6, ""));
		System.out.println(new StringBuilder(s4).insert(0, "hello go,"));
		System.out.println(new StringBuilder(s4).delete(0, 6));
	}
}
