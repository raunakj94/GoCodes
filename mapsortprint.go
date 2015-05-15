
package main
import "sort"

m:= map[int]string{
4:"four",
5:"five",
6:"six"
2:"two",
}
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}
