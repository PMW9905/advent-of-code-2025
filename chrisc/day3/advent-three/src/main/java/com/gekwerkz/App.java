package com.gekwerkz;

import java.io.File;
import java.util.Scanner;

import com.gekwerkz.utils.BatteryBank;

public class App {
	public static void main(String[] args) {
		File input = new File("INPUTGOESHERE");
		System.out.println("JoltageMaster 3000 online");
		
		try (Scanner inputReader = new Scanner(input)) {
			BatteryBank bank = new BatteryBank();
			while (inputReader.hasNextLine()) {
				String inputLn = inputReader.nextLine();
				bank.findMaxJoltage(inputLn);
			}
			System.out.println("Total value of all joltages: " + bank.getTotalJoltage());
		} catch (Exception e) {
			System.out.println("Failed.");
			e.printStackTrace();
		}

	}
}