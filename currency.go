// 大写数字金额生成器
// 单位:
//       "",  拾,   佰,   仟
//   元
//   万 
//   亿
//   万亿
//   万万亿
//
// 小数点后单位：角 分 
// 
// 数字: 零,壹,贰,叁,肆,伍,陆,柒,捌,玖

package cncap

var (
	c_units = []string{"元", "万", "亿", "万亿", "万万亿"/*, ...*/}  // 更大的数字单位虽然可以往后增加，但float64已经不精准了
	c_fracBases = []string{/*..., */"分", "角"} // 更小的金额单位可往前面增加
)

const (
	c_ending = "整"
	c_zero = "零元"
)

// 把数字金额转换成中文大写金额
func CapitalizeCurrency(amount interface{}) string {
	return capitalize_i(amount, &currencyCapitalizer{})
}

type currencyCapitalizer struct {
	capitalizerAdapter
}

func (c *currencyCapitalizer) capitalize(sAmount string, isNeg bool) (<-chan string) {
	return defCapitalize(c, sAmount, isNeg)
}

func (c *currencyCapitalizer) getEnding() string {
	return c_ending
}
func (c *currencyCapitalizer) getZero() string {
	return c_zero
}
func (c *currencyCapitalizer) getFracLen() int {
	return len(c_fracBases)
}
func (c *currencyCapitalizer) getUnits() []string {
	return c_units
}
func (c *currencyCapitalizer) getFracBases() []string {
	return c_fracBases
}
