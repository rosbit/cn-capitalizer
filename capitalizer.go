// 大写数字生成器

package cncap

import (
	"reflect"
	"strings"
	"fmt"
	"strconv"
)

type Capitalizer interface {
	getDigits() []string
	getEnding() string
	getZero() string
	getSeprator() string
	getNegative() string
	getFracLen() int
	capitalize(s string, isNeg bool) (<-chan string)
	getIntCount() int
	getUnits() []string
	getIntBases() []string
	getFracBases() []string
	outputAnyZero() bool
	initAllZero() bool
	initPrevZero() bool
}

type dataGroup struct {
	start, end int
	bases []string
	unit    string
	lastInt bool
}

const (
	tooLarge = "超大数字"
	tooSmall = "超小数字"
	unsupporting = "数据类型不支持"
)

func capitalize_i(d interface{}, c Capitalizer) string {
	sDigit, isNeg, ok := formatDigit(d, c.getFracLen())
	if !ok {
		return sDigit
	}

	res := c.capitalize(sDigit, isNeg)
	r := &strings.Builder{}
	for s := range res {
		r.WriteString(s)
	}
	return r.String()
}

func defCapitalize(c Capitalizer, sDigit string, isNeg bool) (<-chan string) {
	if isNeg {
		sDigit = sDigit[1:]
	}
	length := len(sDigit)
	intLength := strings.IndexByte(sDigit, '.')
	if intLength < 0 {
		intLength = length
	}

	intCount := c.getIntCount()
	nGroup := intLength / intCount
	gEnd := intLength % intCount
	if gEnd > 0 {
		nGroup += 1
	} else if nGroup > 0 {
		gEnd = intCount
	}
	if nGroup > len(c.getUnits()) {
		return makeTooLarge(isNeg)
	}

	groups := make(chan *dataGroup)
	res := capitalize(c, groups, sDigit, isNeg)

	go func() {
		defer close(groups)

		intBases := c.getIntBases()
		units := c.getUnits()
		nGroup -= 1
		for gStart := 0; nGroup >= 0; nGroup-- {
			groups <- &dataGroup {
				start: gStart,
				end: gEnd,
				bases: intBases,
				unit: units[nGroup],
				lastInt: gEnd>=intLength,
			}
			gStart = gEnd
			gEnd += intCount
		}
		groups <- &dataGroup{
			start: intLength + 1,
			end: length,
			bases: c.getFracBases(),
		}
	}()

	return res
}

// 数字型值转为字符串
func formatDigit(d interface{}, fracLen int) (sDigit string, isNeg, ok bool) {
	switch d.(type) {
	case int, int8, int16, int32, int64:
		a := reflect.ValueOf(d).Int()
		isNeg = a < 0
		sDigit = fmt.Sprintf("%d", a)
	case uint, uint8, uint16, uint32, uint64:
		sDigit = fmt.Sprintf("%v", d)
	case float64, float32:
		a := reflect.ValueOf(d).Float()
		return parseFloat(a, fracLen)
	case string:
		a, err := strconv.ParseFloat(d.(string), 64)
		if err != nil {
			sDigit = err.Error()
			return
		}
		return parseFloat(a, fracLen)
	case []byte:
		a, err := strconv.ParseFloat(string(d.([]byte)), 64)
		if err != nil {
			sDigit = err.Error()
			return
		}
		return parseFloat(a, fracLen)
	default:
		sDigit = unsupporting
		return
	}
	ok = true
	return
}

func parseFloat(a float64, fracLen int) (sDigit string, isNeg, ok bool) {
	isNeg = a < 0.0
	floatFormat := fmt.Sprintf("%%.%df", fracLen)
	sDigit = fmt.Sprintf(floatFormat, a)
	ok = true
	return
}

func capitalize(c Capitalizer, g <-chan *dataGroup, strToCap string, isNeg bool) (<-chan string) {
	neg := c.getNegative()
	sep := c.getSeprator()
	zero := c.getZero()
	digits := c.getDigits()
	ending := c.getEnding()
	outputAnyZero := c.outputAnyZero()

	res := make(chan string)

	go func() {
		defer close(res)
		if isNeg {
			res <- neg
		}

		prevZero := c.initPrevZero()
		allZero := c.initAllZero()
		prevGroupAllZero := true

		for dg := range g {
			prevGroupAllZero = true
			gStart, gEnd, bases, unit := dg.start, dg.end, dg.bases, dg.unit
			for i, idx := gStart, gEnd - gStart - 1; i<gEnd; i, idx = i+1, idx-1 {
				d := strToCap[i]
				if d == '0' {
					if !outputAnyZero {
						prevZero = true
					} else if !allZero {
						res <- digits[d - '0']
						res <- bases[idx]
					}
					continue
				}

				if !allZero && prevZero {
					res <- sep
				}

				res <- digits[d - '0']
				res <- bases[idx]
				allZero, prevZero, prevGroupAllZero = false, false, false
			}
			if !prevGroupAllZero || (dg.lastInt && !allZero) {
				res <- unit
			}
		}

		if allZero {
			res <- zero
		}
		if prevGroupAllZero {
			res <- ending
		}
	}()

	return res
}

func makeTooLarge(isNeg bool) (<-chan string) {
	res := make(chan string)
	go func() {
		defer close(res)
		if isNeg {
			res <- tooSmall
		} else {
			res <- tooLarge
		}
	}()
	return res
}
