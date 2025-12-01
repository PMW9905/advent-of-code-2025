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
        this.value = this.value + turns;
        this.value = this.value % 100;
        if(this.value < 0){
            this.value = this.value + 100;
        }
        if(this.value == 0){
            this.zeroCount++;
        }
    }
}
