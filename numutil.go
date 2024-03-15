package numutil

import (
    "bytes"
    "fmt"
    "strconv"
    "strings"
 )

func IntConverter(input any) (string, int64) {
    var str string
    var int int

    if Typeof(input) == "string" {
        int, _ = strconv.Atoi(input.(string))
        str = IntToString(input.(string), ',')
    } else {
        int, _ = strconv.Atoi(fmt.Sprintf("%d", input))
        str = IntToString(strconv.Itoa(int), ',')
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

func Multiplier(intStr string) int {
    multipliers := map[string]int{"k": 1e3, "m": 1e6, "b": 1e9, "t": 1e12}
    multi       :=  1
    num         :=  intStr
    numabbr     :=  strings.ToLower(num[len(num)-1:])

    if val, ok := multipliers[numabbr]; ok {
        num = strings.ToLower(num[:len(num)-1])
        multi = val
    }

    input, _ := strconv.ParseFloat(num, 1)
    return int(input * float64(multi))
 }

func Typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
 }