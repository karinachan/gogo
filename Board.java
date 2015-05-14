public class Board {
	
	public int sideLength; 
	public int[][] array;

	public String toString() {
		String s = ""; 
		for(int i = 0; i < sideLength; i++) {
			for(int j = 0; j < sideLength; j++) {
				s = s + array[i][j] + " ";
			}
			s = s + "\n"; 
		}
		return s; 
	}

	public Board(int s) {
		sideLength = s; 
		array = new int[sideLength][sideLength];
	}

	public void placePiece(int row, int col, int num) {
		array[row][col] = num; 
	}


	public boolean isCaptured(int row, int col, int num) {
		boolean captured = true; 

		if(row-1 >= 0) {
			int val = array[row-1][col];
			if(val==0)
				return false;
			if(val==num) {
				placePiece(row, col, 3);
				captured = isCaptured(row-1, col, num);
			}
			if(val==3) 
				placePiece(row-1, col, num);
		}

		if(row+1 <= array.length-1) {
			int val = array[row+1][col];
			if(val==0)
				return false;
			if(val==num) {
				placePiece(row, col, 3);
				captured = isCaptured(row+1, col, num);
			}
			if(val==3) 
				placePiece(row+1, col, num); 
		}

		if(col-1 >= 0) {
			int val = array[row][col-1];
			if(val==0)
				return false;
			if(val==num) {
				placePiece(row, col, 3);
				captured = isCaptured(row, col-1, num);
			}
			if(val==3) 
				placePiece(row, col-1, num); 
		}

		if(col+1 <= array.length-1) {
			int val = array[row][col+1];
			if(val==0)
				return false;
			if(val==num) {
				placePiece(row, col, 3);
				captured = isCaptured(row, col+1, num);
			}
			if(val==3) 
				placePiece(row-1, col, num); 
		}

		return captured;	
	}

	public static void main(String[] args) {
		Board b = new Board(5); 

		b.placePiece(0,0,1);
		b.placePiece(0,1,2);
		b.placePiece(1,0,2);

	/*	b.placePiece(2,2,2);
		b.placePiece(1,2,1);
		b.placePiece(3,2,1);
		b.placePiece(2,1,1);
		b.placePiece(2,3,2);
		b.placePiece(1,3,1);
		b.placePiece(3,3,1);
		b.placePiece(2,4,1);*/


		System.out.println(b.toString());
		System.out.println(b.isCaptured(0,0,1));
	}


}