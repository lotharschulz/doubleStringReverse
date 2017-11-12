package main

import(
  "fmt"
  "myStringUtil"
)

func main() {
  var i string = "日本語_ß_ДЕЖ_व्यंजन_abcDEFghiJKLmnoPQRtsuVWXyzöäüÖÄÜ"
  fmt.Println("myStringUtil.DoubleStringReverse(" + i + "): " + myStringUtil.DoubleStringReverse(i))
}
