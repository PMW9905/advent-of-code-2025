package com.gekwerkz;

import java.io.File;
import java.util.Scanner;

public class App {
    public static void main(String[] args) {
        {
            File input = new File("E:\\Advent of Code\\advent-of-code-2025\\chrisc\\day1\\advent-one\\data\\input.txt");
            Dial dial = new Dial();
            System.out.println("Dialmaster 2000 online");
            System.out.println("Current Dial setting:" + dial.getValue());

            try (Scanner inputReader = new Scanner(input)) {
                while (inputReader.hasNextLine()) {
                    String inputField = inputReader.nextLine();
                    if (!inputField.isEmpty()) {
                        char direction = inputField.charAt(0);
                        String inputNumString = inputField.substring(1);
                        int inputNum = Integer.parseInt(inputNumString);
                        if (direction == 'R') {
                            dial.turnRight(inputNum);
                        }
                        if (direction == 'L') {
                            dial.turnLeft(inputNum);
                        }
                    }
                }
            } catch (Exception e) {
                System.out.println("Failed.");
                e.printStackTrace();
            }
            System.out.println("Total number of zeroes: " + dial.getZeroCount());
        }

    }
}