package com.gekwerkz;

public class Dial {

    private int value;
    private int zeroCount;

    Dial() {
        value = 50;
        zeroCount = 0;
    }

    public int getValue() {
        return this.value;
    }

    public int getZeroCount() {
        return this.zeroCount;
    }

    public void turnLeft(int input) {
        input = input * -1;
        turnDial(input);
    }

    public void turnRight(int input) {
        turnDial(input);
    }

    private void turnDial(int turns) {
        int totalZeroOccurences = 0;
        if(this.value > 0 && this.value + turns <= 0 && turns < 0){
            this.value = this.value + -100;
        }
        this.value = this.value + turns;
        totalZeroOccurences = turns < 0 ? (this.value / 100) * -1: this.value / 100;
        this.value = this.value % 100;
        if(this.value < 0){
            this.value = this.value + 100;
        }
        zeroCount = zeroCount + totalZeroOccurences;
    }
}
