// 大写日期生成器 
// 数字: 零,壹,贰,叁,肆,伍,陆,柒,捌,玖
//    拾/廿/卅
//    佰
//    仟
// 支票日期要求:
//    年份: 逐个数字转为大写
//    月/日份：<=10 或 20/30 前面加“零”：零壹、零拾、零贰拾

package cncap

import (
	"strconv"
	"fmt"
)

var (
	y_units = []string{"", "万"/*, ...*/}
	d_tens = []string{"", "拾", "廿", "卅"}
	d_ones = []string{"","壹","贰","叁","肆","伍","陆","柒","捌","玖"}
	cheque_intBases = []string{"", "拾"}
)

// 把数字转为年份（带个拾佰仟）
func CapitalizeYear(dYear interface{}) string {
	return capitalize_i(dYear, &yearCapitalizer{})
}

// 把数字转为年份（不带个拾佰仟,可用于支票）
func CapitalizeYearWithoutBase(dYear interface{}) string {
	return capitalize_i(dYear, &year_no_baseCapitalizer{})
}

// 把数字转为月份
func CapitalizeMonth(dMonth interface{}) string {
	return capitalizeMonthDay(dMonth, 12)
}

// 把数字转为日份
func CapitalizeDay(dDay interface{}) string {
	return capitalizeMonthDay(dDay, 31)
}

// 把数字转为月份（符合支票规范）
func CapitalizeChequeMonth(dMonth interface{}) string {
	return capitalize_i(dMonth, &cheque_monthdayCapitalizer{isMonth:true})
}

// 把数字转为日份（符合支票规范）
func CapitalizeChequeDay(dDay interface{}) string {
	return capitalize_i(dDay, &cheque_monthdayCapitalizer{})
}

// --- 非支票月份、日份 ---
func capitalizeMonthDay(dDigit interface{}, maxVal int) string {
	sDigit, isNeg, ok := formatDigit(dDigit, 0)
	if !ok {
		return sDigit
	}
	if isNeg {
		return tooSmall
	}
	digit, err := strconv.ParseUint(sDigit, 10, 64)
	if err != nil {
		return unsupporting
	}
	if digit == 0 {
		return tooSmall
	}
	md := int(digit)
	if md > maxVal {
		return tooLarge
	}
	return toMonthDate(md)
}

func toMonthDate(md int) string {
	tenIdx := md / 10
	oneIdx := md % 10
	return fmt.Sprintf("%s%s", d_tens[tenIdx], d_ones[oneIdx])
}

// ---- 年（带个拾佰仟）-----
type yearCapitalizer struct {
	capitalizerAdapter
}
func (y *yearCapitalizer) capitalize(sYear string, isNeg bool) (<-chan string) {
	return defCapitalize(y, sYear, isNeg)
}
func (y *yearCapitalizer) getNegative() string {
	return "公元前"
}
func (d *yearCapitalizer) getUnits() []string {
	return y_units
}

// ---- 年（不带个拾佰仟）-----
type year_no_baseCapitalizer struct {
	yearCapitalizer
}
func (y *year_no_baseCapitalizer) capitalize(sYear string, isNeg bool) (<-chan string) {
	return defCapitalize(y, sYear, isNeg)
}
func (y *year_no_baseCapitalizer) getUnits() []string {
	return []string{"", ""}
}
func (y *year_no_baseCapitalizer) getIntBases() []string {
	return []string{"", "", "", ""}
}
func (y *year_no_baseCapitalizer) getIntCount() int {
	return 4
}
func (y *year_no_baseCapitalizer) outputAnyZero() bool {
	return true
}
func (y *year_no_baseCapitalizer) getSeprator() string {
	return ""
}

// --- 支票月份、日份 ---
type cheque_monthdayCapitalizer struct {
	capitalizerAdapter
	sMonthDay string
	isMonth bool
}
func (m *cheque_monthdayCapitalizer) capitalize(sMonthDay string, isNeg bool) (<-chan string) {
	if isNeg {
		return makeTooLarge(isNeg)
	}
	if sMonthDay == "10" {
		res := make(chan string)
		go func() {
			res <- "零拾"
			close(res)
		}()
		return res
	}

	switch len(sMonthDay) {
	case 1:
		if sMonthDay == "0" {
			return makeTooLarge(true)
		}
	case 2:
		if m.isMonth {
			if sMonthDay[0] > '1' || sMonthDay[1] > '2' {
				return makeTooLarge(isNeg)
			}
		} else {
			if sMonthDay[0] > '3' || (sMonthDay[0] == '3' && sMonthDay[1] > '1') {
				return makeTooLarge(isNeg)
			}
		}
	default:
		return makeTooLarge(isNeg)
	}

	m.sMonthDay = sMonthDay
	return defCapitalize(m, sMonthDay, isNeg)
}
func (m *cheque_monthdayCapitalizer) getDigits() []string {
	return d_ones
}
func (m *cheque_monthdayCapitalizer) getEnding() string {
	return ""
}
func (m *cheque_monthdayCapitalizer) getZero() string  {
	return ""
}
func (m *cheque_monthdayCapitalizer) getNegative() string {
	return ""
}
func (m *cheque_monthdayCapitalizer) getIntCount() int {
	return len(cheque_intBases)
}
func (m *cheque_monthdayCapitalizer) getUnits() []string {
	return []string{""}
}
func (m *cheque_monthdayCapitalizer) getIntBases() []string {
	return cheque_intBases
}
func (m *cheque_monthdayCapitalizer) outputAnyZero() bool {
	return true
}
func (m *cheque_monthdayCapitalizer) initAllZero() bool {
	switch len(m.sMonthDay) {
	case 1:
		return false
	default:
		return m.sMonthDay[1] != '0'
	}
}
func (m *cheque_monthdayCapitalizer) initPrevZero() bool {
	switch len(m.sMonthDay) {
	case 1:
		return true
	default:
		return m.sMonthDay[1] == '0'
	}
}
