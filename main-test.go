package main

import "testing"

func TestAbs(t *testing.T) {
    got := Abs(-1)

    if got != 1 {
        t.Error("Wrong Number")
    }
}

func Abs(input int) int {
    if input < 0{
        return -input
    }
        return input
}
