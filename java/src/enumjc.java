public class enumjc {
	public static void main(String[] args) {
		enumjcen en = enumjcen.RED;
		System.out.println(en.getColor());
		System.out.println(en.ordinal());
		System.out.println(enumjcen.valueOf("RED"));
		en.printColor();
		for (enumjcen v : enumjcen.values()) {
			System.out.println(v);
		}
	}
}

enum enumjcen {
	RED {
		public String getColor() {
			return "red";
		}
	},
	BLUE {
		public String getColor() {
			return "blue";
		}
	},
	GREEN {
		public String getColor() {
			return "green";
		}
	};
	
	public abstract String getColor();
	
	public void printColor(){
		System.out.println(this.ordinal());
	}
}
