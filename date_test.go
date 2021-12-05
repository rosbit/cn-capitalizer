package cncap

import (
	"testing"
	"fmt"
)

func testYear(f interface{}) {
	sYear, _, ok := formatDigit(f, 0)
	if !ok {
		fmt.Printf("%v => %s\n", f, sYear)
	} else {
		fmt.Printf("%s => %s (with base) \n", sYear, CapitalizeYear(f))
		fmt.Printf("%s => %s (without base)\n", sYear, CapitalizeYearWithoutBase(f))
	}
}

func TestYear(t *testing.T) {
	fmt.Printf("---- 年份 ----\n")
	testYear(2021)
	testYear(19800)
	testYear(2980)
	testYear(500200)
	testYear(103)
	testYear(int16(32766))
	testYear(0)
	testYear(-100)
}

func testMonth(f interface{}) {
	sMonth, _, ok := formatDigit(f, 0)
	if !ok {
		fmt.Printf("%v => %s\n", f, sMonth)
	} else {
		fmt.Printf("%s => %s\n", sMonth, CapitalizeMonth(f))
	}
}
func TestMonth(t *testing.T) {
	fmt.Printf("---- 月份 ----\n")
	testMonth(-1)
	testMonth(0)
	testMonth(1)
	testMonth(9)
	testMonth(10)
	testMonth(11)
	testMonth(12)
	testMonth(13)
}

func testDay(f interface{}) {
	sDay, _, ok := formatDigit(f, 0)
	if !ok {
		fmt.Printf("%v => %s\n", f, sDay)
	} else {
		fmt.Printf("%s => %s\n", sDay, CapitalizeDay(f))
	}
}
func TestDay(t *testing.T) {
	fmt.Printf("---- 日份 ----\n")
	testDay(-1)
	testDay(0)
	testDay(1)
	testDay(9)
	testDay(10)
	testDay(11)
	testDay(12)
	testDay(20)
	testDay(21)
	testDay(30)
	testDay(31)
	testDay(32)
}

func testChequeMonth(f interface{}) {
	sMonth, _, ok := formatDigit(f, 0)
	if !ok {
		fmt.Printf("%v => %s\n", f, sMonth)
	} else {
		fmt.Printf("%s => %s\n", sMonth, CapitalizeChequeMonth(f))
	}
}
func TestChequeMonth(t *testing.T) {
	fmt.Printf("---- 用于支票的月份 ----\n")
	testChequeMonth(0)
	testChequeMonth(1)
	testChequeMonth(10)
	testChequeMonth(11)
	testChequeMonth(12)
	testChequeMonth(13)
}

func testChequeDay(f interface{}) {
	sDay, _, ok := formatDigit(f, 0)
	if !ok {
		fmt.Printf("%v => %s\n", f, sDay)
	} else {
		fmt.Printf("%s => %s\n", sDay, CapitalizeChequeDay(f))
	}
}
func TestChequeDay(t *testing.T) {
	fmt.Printf("---- 用于支票的日份 ----\n")
	testChequeDay(0)
	testChequeDay(1)
	testChequeDay(10)
	testChequeDay(19)
	testChequeDay(20)
	testChequeDay(29)
	testChequeDay("30")
	testChequeDay(31)
	testChequeDay(32)
}
