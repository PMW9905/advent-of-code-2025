package com.gekwerkz.utils;

import java.util.ArrayList;
import java.util.List;

public class BatteryBank {

	private List<Integer> joltages;

	public BatteryBank() {
		joltages = new ArrayList<Integer>();
	}

	public void findMaxJoltage(String input) {
		// final number values;
		int finalFirstValue = 0;
		int finalSecondValue = 0;
		// placeholders for searching
		int numFirstInstanceIndex;
		int numLastInstanceIndex;
		String joltBank = input;
		for (int i = 9; i > 0; i--) {
			// make sure number exists in string
			char curNum = (char) (i+'0');
			if (joltBank.indexOf(curNum) != -1) {
				//then set index values
				numFirstInstanceIndex = joltBank.indexOf(curNum);
				numLastInstanceIndex = joltBank.lastIndexOf(curNum);
				//if only one instance of number exists, check length, then build from substring
				if(numFirstInstanceIndex==numLastInstanceIndex) {
					//if there is enough length left after the instance of biggest number, use that substring to build
					if(numFirstInstanceIndex <= joltBank.length() - 2) {
						finalFirstValue = i;
						finalSecondValue = getSecondJoltageValue(joltBank.substring(numFirstInstanceIndex + 1));
					}else {
						/*
						 * When the number is too close to the end(eg 81119)
						 * the loop needs to break and move to next highest num
						 * to calc joltage
						 * */
						continue;
					}
				}else if(numFirstInstanceIndex!=numLastInstanceIndex) {
					finalFirstValue = i;
					finalSecondValue = getSecondJoltageValue(joltBank.substring(numFirstInstanceIndex + 1));
				}
				
				
				
			}
			if(finalFirstValue != 0 && finalSecondValue != 0) {
				break;
			}
		}
		if(finalFirstValue + finalSecondValue > 0 && finalFirstValue != 0) {
			// :)
			addJoltage(Integer.parseInt(Integer.toString(finalFirstValue).concat(Integer.toString(finalSecondValue))));
		}
	}
	
	private int getSecondJoltageValue(String input) {
		int topNum = 0;
		int curNum;
		String curStr;
		for(int j = 0; j < input.length(); j++) {
			curStr = String.valueOf(input.charAt(j));
			curNum = Integer.parseInt(curStr);
			if(curNum > topNum) {
				topNum = curNum;
			}
		}
		return topNum;
	}
	
	private void addJoltage(int joltage) {
		joltages.add(joltage);
	}

	public void printAllJolts() {
		
	}
	
	public long getTotalJoltage() {
		long total = 0;
		for (int joltage : joltages) {
			total = total + joltage;
		}
		return total;
	}
}
