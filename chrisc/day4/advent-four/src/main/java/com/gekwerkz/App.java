package com.gekwerkz;

import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

import com.gekwerkz.utils.PaperPal;

public class App {
	public static void main(String[] args) {
		File input = new File("INPUTGOESHERE");
		System.out.println("PrinterMaster 4000 online");
		
		try (Scanner inputReader = new Scanner(input)) {
			PaperPal pal = new PaperPal();
			List<char[]> inputArr = new ArrayList<char[]>();
			while (inputReader.hasNextLine()) {
				String inputLn = inputReader.nextLine();
				inputArr.add(inputLn.toCharArray());
				
			}
			int paperSize = inputArr.size();
			int paperRowLen = inputArr.get(0).length;
			char[][] paperMap = new char[paperSize][paperRowLen];
			for(int i = 0; i < paperSize; i++) {
				paperMap[i] = inputArr.get(i);
			}
			System.out.println("Total paper to be accessed: " + pal.getTotalPaperAccess(paperMap));
		} catch (Exception e) {
			System.out.println("Failed.");
			e.printStackTrace();
		}

	}
}