package numutil

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
 )

type invalidNumber struct{}

func (m *invalidNumber) Error() string {
	return fmt.Sprintf("%v", m)
 }

func IntConverter(input any) (string, int64) {
	var str string
	var int int
	var err error

	if Typeof(input) == "string" {
		int, err = strconv.Atoi(input.(string))
		str = IntToString(input.(string), ',')
		if err != nil {
			panic(err)
		}
	} else {
		int, err = strconv.Atoi(fmt.Sprintf("%d", input))
		str = IntToString(strconv.Itoa(int), ',')
		if err != nil {
			panic(err)
		}
	}

	return str, int64(int)
 }

func IntToString(s string, sep rune) string {
	n, _ := strconv.Atoi(s)

	startOffset := 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	l := len(s)

	separateAt := 3 - ((l - startOffset) % 3)

	if (separateAt == 3) {
		separateAt = 0
	}

	for i := startOffset; i < l; i++ {
		if (separateAt == 3) {
			buff.WriteRune(sep)
			separateAt = 0
		}
		separateAt++
		buff.WriteByte(s[i])
	}

	return buff.String()
 }

func Multiplier(num string) (int64) {
	multipliers := map[string]int{"k": 1e3, "m": 1e6, "b": 1e9, "t": 1e12}
	numabbr     :=  strings.ToLower(num[len(num)-1:])
	multi       :=  1

	if strings.Count(num, ".") > 1 {
		return 0
	}

	if val, ok := multipliers[numabbr]; ok {
		num = num[:len(num)-1]
		multi = val
	}

	re := regexp.MustCompile("[^0-9.]")
	num = re.ReplaceAllString(num, "")

	input, err := strconv.ParseFloat(num, 64)
	if err != nil {
		panic(err)
	}

	return int64(input * float64(multi))
 }

func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
 }