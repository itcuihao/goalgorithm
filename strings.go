package main
 
import (
  "fmt"
  "strings"
)
 
func main() {
  fmt.Println(strings.Compare("A", "B"))  // A < B
  fmt.Println(strings.Compare("B", "A"))  // B > A  
  fmt.Println(strings.Compare("Japan", "Australia"))  // J > A
  fmt.Println(strings.Compare("Australia", "Japan"))  // A < J
  fmt.Println(strings.Compare("Germany", "Germany"))  // G == G
  fmt.Println(strings.Compare("Germany", "GERMANY"))  // GERMANY > Germany
  fmt.Println(strings.Compare("", ""))
  fmt.Println(strings.Compare("", " ")) // Space is less

  fmt.Println(strings.Contains("Australia", "Aus"))   // Any part of string
  fmt.Println(strings.Contains("Australia", "Australian")) 
  fmt.Println(strings.Contains("Japan", "JAPAN")) // Case sensitive
  fmt.Println(strings.Contains("Japan", "JAP")) // Case sensitive
  fmt.Println(strings.Contains("Become inspired to travel to Australia.", "Australia"))
  fmt.Println(strings.Contains("", ""))
  fmt.Println(strings.Contains("  ", " ")) // space also consider as string
  fmt.Println(strings.Contains("12554", "1"))   
}