import java.util.HashMap;
import java.util.Map;

public class hashmapjc {
	public static void main(String[] args) {
		HashMap<String, Integer> hm = new HashMap<>();
		hm.put("hello", 13);
		hm.put("world", 19);
		
		HashMap<String, Integer> hm2 = new HashMap<>();
		hm2.put("java", 20);
		hm2.put("go", 11);
		hm2.put("c++", 24);
		hm.putAll(hm2);
		
		hm.remove("hello");
		
		hm.replace("world", 13, 17);
		hm.replace("world", 14);
		
		System.out.println(hm.get("hello"));
		
		System.out.println(hm.size());
		System.out.println(hm.isEmpty());
		
		for (String s : hm.keySet()) {
			System.out.printf("k:%s,v:%d\n", s, hm.get(s));
		}
		
		for (int i : hm.values()) {
			System.out.println(i);
		}
	}
}
