// JavaClient.java
// Author: Emery Otopalik, September 2014

import java.net.*;

public class JavaClient {

	public static void main(String [] args) {
		try {
			int port = Integer.parseInt(args[0]);
			Socket socket = new Socket(args[1], port);
			System.err.println("Connected to " + args[1] + " on port " + port);
			new JavaReadWrite(System.in, socket.getOutputStream()).start();
			new JavaReadWrite(socket.getInputStream(), System.out).start();
		} catch (Exception e) {
			e.printStackTrace();
			System.err.println("\nUsage: java Client <port> localhost");
		}
	}

}