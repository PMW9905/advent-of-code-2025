package com.gekwerkz;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

/**
 * Unit test for simple App.
 */
public class AppTest 
{
    /**
     * Rigorous Test :-)
     */
    @Test
    public void testDialNegative()
    {
        int expected = 0;
        Dial dial = new Dial();
        
        dial.turnLeft(50);
        

        int actual = dial.getValue();
        assertEquals(Integer.toString(expected), Integer.toString(actual));
        

    }

    /**
     * Rigorous Test :-)
     */
    @Test
    public void testDialPositive()
    {
        int expected = 0;
        Dial dial = new Dial();
        dial.turnRight(50);

        int actual = dial.getValue();
        assertEquals(Integer.toString(expected), Integer.toString(actual));
        

    }

    /**
     * Rigorous Test :-)
     */
    @Test
    public void testDialLargeNegative()
    {
        int expected = 0;
        Dial dial = new Dial();
        dial.turnLeft(250);

        int actual = dial.getValue();
        assertEquals(Integer.toString(expected), Integer.toString(actual));
        

    }

    @Test
    public void testDialLargePositive()
    {
        int expected = 0;
        Dial dial = new Dial();
        dial.turnRight(250);

        int actual = dial.getValue();
        assertEquals(Integer.toString(expected), Integer.toString(actual));
        

    }

    @Test
    public void testDialListNegative()
    {
        int expected = 94;
        Dial dial = new Dial();
        dial.turnLeft(50);
        dial.turnLeft(250);
        dial.turnLeft(2050);
        dial.turnLeft(1);
        dial.turnLeft(2);
        dial.turnLeft(3);

        int actual = dial.getValue();
        assertEquals(Integer.toString(expected), Integer.toString(actual));
        

    }

    @Test
    public void testDialListPositive()
    {
        int expected = 6;
        Dial dial = new Dial();
        dial.turnRight(1); 
        dial.turnRight(2); 
        dial.turnRight(3); 
        dial.turnRight(50);
        dial.turnRight(250);
        dial.turnRight(250);

        int actual = dial.getValue();
        assertEquals(Integer.toString(expected), Integer.toString(actual));
        

    }
}
