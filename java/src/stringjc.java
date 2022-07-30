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
		StringBuilder sb = new StringBuilder(s4);
		sb.ensureCapacity(20);
		System.out.println(sb.capacity());
		System.out.println(sb.replace(0, 6, ""));
		System.out.println(sb.insert(0, "hello go,"));
		System.out.println(sb.delete(0, 6));
	}
}
