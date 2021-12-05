package cncap

var (
	common_digits = []string{"零","壹","贰","叁","肆","伍","陆","柒","捌","玖"}
	common_intBases = []string{"", "拾", "佰", "仟"}
)

type capitalizerAdapter struct{}

func (c *capitalizerAdapter) capitalize(sDigit string, isNeg bool) (<-chan string) {
	return nil
}

func (c *capitalizerAdapter) getDigits() []string {
	return common_digits
}
func (c *capitalizerAdapter) getEnding() string {
	return ""
}
func (c *capitalizerAdapter) getZero() string {
	return common_digits[0]
}
func (c *capitalizerAdapter) getSeprator() string {
	return common_digits[0]
}
func (c *capitalizerAdapter) getNegative() string {
	return "负"
}
func (c *capitalizerAdapter) getFracLen() int {
	return 0
}
func (c *capitalizerAdapter) getIntCount() int {
	return len(common_intBases)
}
func (c *capitalizerAdapter) getUnits() []string {
	return nil
}
func (c *capitalizerAdapter) getIntBases() []string {
	return common_intBases
}
func (c *capitalizerAdapter) getFracBases() []string {
	return nil
}
func (c *capitalizerAdapter) outputAnyZero() bool {
	return false
}
func (c *capitalizerAdapter) initAllZero() bool {
	return true
}
func (c *capitalizerAdapter) initPrevZero() bool {
	return false
}
