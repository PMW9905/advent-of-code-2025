package com.gekwerkz;

import java.io.File;
import java.util.Scanner;

import com.gekwerkz.utils.IdFilterUtil;

public class App {
	public static void main(String[] args) {
		File input = new File("INPUTGOESHERE");
		System.out.println("IDMaster 2000 online");
		
		try (Scanner inputReader = new Scanner(input)) {
			IdFilterUtil util = new IdFilterUtil();
			while (inputReader.hasNextLine()) {
				String inputField = inputReader.nextLine();
				String[] idRanges = inputField.split(",");
				if (idRanges.length > 0) {
					for(String range : idRanges) {
						util.filterIdRange(range);
					}
				}
			}
			System.out.println("Total value of invalid IDs: " + util.getInvalidSum());
		} catch (Exception e) {
			System.out.println("Failed.");
			e.printStackTrace();
		}

	}
}