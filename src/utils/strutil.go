
/*

   strutils.go implements several helpful string functions that
   will assist in string exercises/manipulations.

   Authored on 1/24/2015 by rctillotson25


*/

package utils



/*
   CharCountMap (could probably be renamed) takes in a string
   and adds each of the individual characters into a map and returns the map.

*/
func CharCountMap(s string) map[rune]int {
	hash := make(map[rune]int)

	for _, value := range s {
		hash[value]++	
	}

	return hash
}
