package main

import (
	"fmt"
	"regexp"
)

/*
ReplaceAllString returns a copy of src, 
replacing matches of the Regexp with the replacement string repl. 
Inside repl, $ signs are interpreted as in Expand, 
so for instance $1 represents the text of the first submatch.
*/
func main() {
	re := regexp.MustCompile("a(x*)b");
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T")); //-T-T-
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1")); //--xx-
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W")); //---
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W")); //-W-xxW-
}

/*
-T-T-
--xx-
---
-W-xxW-
*/