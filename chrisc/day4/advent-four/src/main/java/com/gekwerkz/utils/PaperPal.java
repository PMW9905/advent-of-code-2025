package com.gekwerkz.utils;

import java.util.ArrayList;
import java.util.List;

public class PaperPal {
	char[][] paperMap;
	private char paperSymbol = '@';
	private int totalPaper = 0;
	private int maxLenI;
	private int maxLenJ;

	public PaperPal() {
	}

	public int getTotalPaperAccess(char[][] input) {
		this.paperMap = input;
		if (paperMap != null) {
			maxLenI = paperMap.length;
			// loop through values
			for (int i = 0; i < paperMap.length; i++) {
				maxLenJ = paperMap[i].length;
				for (int j = 0; j < paperMap[i].length; j++) {
					char cur = paperMap[i][j];
					if (paperSymbol == cur) {
						if (isPaperValid(getCheckValues(i, j))) {
							totalPaper++;
						}
					}
				}
			}
		}
		return totalPaper;
	}

	private boolean isPaperValid(List<String> checklist) {
		int x = 0;
		int y = 0;
		int z = 0;
		for (String xy : checklist) {
			String[] coordinates = xy.split(",");
			if (coordinates.length == 2) {
				x = Integer.parseInt(coordinates[0]);
				y = Integer.parseInt(coordinates[1]);
				if (paperSymbol == paperMap[x][y]) {
					z++;
				}
				if (z > 3) {
					//if more than 3 paper around the current checked paper, return false
					return false;
				}
			}
		}
		return true;
	}

	private List<String> getCheckValues(int x, int y) {
		List<String> coords = new ArrayList<String>();
		// need following
		// array above
		// array below
		// array adjacent
		if (x > 0) {
			// above to left, only if exists(for side rows)
			if (y > 0) {
				coords.add(buildCoordinateString(x, y, -1, -1));
			}
			// above direct
			coords.add(buildCoordinateString(x, y, -1, 0));
			// above to right
			if (y < maxLenJ - 1) {
				coords.add(buildCoordinateString(x, y, -1, 1));
			}

		}

		if (x < maxLenI - 1) {
			// below to left
			if (y > 0) {
				coords.add(buildCoordinateString(x, y, 1, -1));
			}
			// below direct
			coords.add(buildCoordinateString(x, y, 1, 0));
			// below to right
			if (y < maxLenJ - 1) {
				coords.add(buildCoordinateString(x, y, 1, 1));
			}
		}

		if (y > 0) {
			// to the left
			coords.add(buildCoordinateString(x, y, 0, -1));
		}

		if (y < maxLenJ - 1) {
			// to the right
			coords.add(buildCoordinateString(x, y, 0, 1));
		}

		return coords;
	}

	private String buildCoordinateString(int x, int y, int xMod, int yMod) {
		StringBuilder sb = new StringBuilder();
		sb.append(x + xMod).append(",").append(y + yMod);
		return sb.toString();
	}

}
