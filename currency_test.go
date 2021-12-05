package cncap

import (
	"testing"
	"fmt"
)

func testCurrency(f interface{}) {
	sAmount, _, ok := formatDigit(f, len(c_fracBases))
	if !ok {
		fmt.Printf("%v => %s\n", f, sAmount)
	} else {
		fmt.Printf("%s => %s\n", sAmount, CapitalizeCurrency(f))
	}
}

func TestCurrency(t *testing.T) {
	fmt.Printf("----- 金额 ----\n")
	testCurrency(435235324532.0) // 元整
	testCurrency(100.02)         // 佰元零x分
	testCurrency(100.2)          // 佰元零x角
	testCurrency(10.23)          // 拾元零x角x分
	testCurrency(101.23)         // 元x角x分
	testCurrency(101.03)         // 零x分
	testCurrency(float32(10100.02))
	testCurrency(340210100.02)
	testCurrency(-340210100.02)  // 负
	testCurrency(3400000000.02)
	testCurrency(3400000.02)
	testCurrency(4352352343400000000.02)    // 万万亿
	testCurrency(uint64(14352352343400000000))
	testCurrency(214352352343400000000.02)  // 超大数额
	testCurrency(-214352352343400000000.02) // 超小数额
	testCurrency(.00)    // 零元
	testCurrency(.12)    // x角x分
	testCurrency(9999)   // 元整
	testCurrency(19800)  // 佰元整
	testCurrency(2980)   // 拾元整
	testCurrency(500200)
	testCurrency(103)    // 佰零x元整
	testCurrency(int16(32766))
	testCurrency(-100.23)
	testCurrency(0)
}

