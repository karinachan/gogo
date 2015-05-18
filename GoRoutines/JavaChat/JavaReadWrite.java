// JavaReadWrite.java
// Author: Emery Otopalik, September 2014

import java.io.*;

public class JavaReadWrite extends Thread {

	private BufferedReader input;
	private PrintWriter output;

	public JavaReadWrite(InputStream input, OutputStream output) {
		this.input = new BufferedReader(new InputStreamReader(input));
		this.output = new PrintWriter(output, true);
	}

	public void run() {
		try {
			String line;
			while ((line = input.readLine()) != null) {
				output.println(line);
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
}