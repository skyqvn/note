import java.util.BitSet;

public class bitsetjc {
	public static void main(String[] args) {
		BitSet bs = new BitSet();
		
		BitSet bs2 = new BitSet();
		
		bs.and(bs2);
		bs.or(bs2);
		bs.andNot(bs2);
		bs.flip(0, bs.length());
		
		bs.set(0, true);
		bs.set(0, 3, true);
		
		bs2 = (BitSet) bs.clone();
		
		System.out.println(bs.size());
		System.out.println(bs.length());
		
		System.out.println(bs.get(0));
		System.out.println(bs.get(0, 8));
		
		System.out.println(bs.isEmpty());
		
		System.out.println(bs.hashCode());
	}
}
