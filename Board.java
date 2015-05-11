public class Board {

	public int sideLength;
	public int[][] array;

	public String toString() {
		String s = "";
		for(int i = 0; i < sideLength; i++) {
			for(int j = 0; j < sideLength; j++) {
				s = s + array[i][j];
			}
			s = s + "\n";
		}
		return s;
	}

	public Board() {
		sideLength = 10;
		array = new int[sideLength][sideLength];
	}

	public static void main(String[] args) {
		Board b = new Board();
		System.out.println(b.toString());
	}


}
