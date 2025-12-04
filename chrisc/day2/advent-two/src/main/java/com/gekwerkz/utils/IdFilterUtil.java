package com.gekwerkz.utils;

import java.util.ArrayList;
import java.util.List;

public class IdFilterUtil {
	
	List<Long> invalidIds;
	
	public IdFilterUtil(){
		invalidIds = new ArrayList<Long>();
	}
	
	public List<Long> getInvalidIds() {
		return this.invalidIds;
	}
	
	public long getInvalidSum() {
		long sum = 0;
		if(!invalidIds.isEmpty()) {
			for(long id : invalidIds) {
				sum = sum + id;
			}
		}
		return sum;
	}
	
	public void filterIdRange(String range) throws Exception{
		String[] rangeNums = range.split("-");
		String startRangeStr = "";
		String endRangeStr = "";
		long startRange = 0;
		long endRange = 0;
		if(rangeNums.length == 2) {
			startRangeStr = rangeNums[0];
			endRangeStr = rangeNums[1];
			//at least one part of the range has to be even num in length, else there are no repeated strings
			if(startRangeStr.length() % 2 == 0 || endRangeStr.length() % 2 == 0) {
				startRange = Long.parseLong(startRangeStr);
				endRange = Long.parseLong(endRangeStr);
				//catch reversed rangees
				if(startRange > endRange) {
					filterRange(endRange,startRange);
				}else {
					filterRange(startRange,endRange);
				}
				
			}else {
				System.out.println("Range does not contain invalid IDs");
			}
		}else {
			throw new Exception("Invalid range provided");
		}
	}
	
	private void filterRange(long start, long end) {
		long cur = start;
		String curStr = Long.toString(cur);
		int curLen = 0;
		String curFront = "";
		String curBack = "";
		while(cur <= end) {
			curStr = Long.toString(cur);
			//get the split point
			curLen = curStr.length();
			curFront = curStr.substring(0, curLen/2);
			curBack = curStr.substring(curLen/2);
			if(curFront.contains(curBack)) {
				System.out.println("Added to invalid IDS:" + cur);
				invalidIds.add(cur);
			}
			cur++;
		}
	}
	 

}
