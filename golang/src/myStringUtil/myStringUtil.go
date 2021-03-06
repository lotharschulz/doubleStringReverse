package myStringUtil

import(
  "unicode"
  //"fmt"
  //"strconv"
)


func DoubleStringReverse(s string) string {
  r := []rune(s)
  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
    //fmt.Println(string(r[i]) + " at [" + strconv.Itoa(i) + "] " + strconv.FormatBool(unicode.IsUpper(r[i])) + " swapcase(" + strconv.Itoa(i) + "): " + string(swapCase(r[i])) )
    //fmt.Println(string(r[j]) + " at [" + strconv.Itoa(j) + "] " + strconv.FormatBool(unicode.IsUpper(r[j])) + " swapcase(" + strconv.Itoa(j) + "): " + string(swapCase(r[j])) )
    r[i], r[j] = swapCase(r[j]), swapCase(r[i])
  }

  if len(r) > 0 && len(r) % 2 == 1 {
    var index = (len(r)-1)/2
    //fmt.Println(string(r[index]) + " at [" + strconv.Itoa(index) + "] " + strconv.FormatBool(unicode.IsUpper(r[index])) + " swapcase(" + strconv.Itoa(index) + "): " + string(swapCase(r[index])) )
    r[index] = swapCase(index)
  }
  return string(r)
}

func swapCase(r rune) rune {
    var res rune
    res = unicode.ToLower(r)
    if unicode.IsLower(r) {
	     res = unicode.ToUpper(r)
    }
    return res
}
