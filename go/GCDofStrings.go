/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""

*/

import (
    "strings"
)
func gcdOfStrings(str1 string, str2 string) string {
    //Dynamic slice tht you can append to
    var prefix = make([]byte, 0)
    shortLen := 0
    chosenString, verifyString := "", ""
    //Find the shorter string
    if (len(str1) < len(str2)) {
        chosenString = str1
        verifyString = str2
    } else {
        chosenString = str2
        verifyString = str1
    }
    shortLen = len(chosenString)
    for i:=0;i < shortLen;i++{
        //Keep adding chars to prefix
        prefix = append(prefix, chosenString[i])
        //If len of prefix is divides the length of string then only can it be a multiple
        if (shortLen%len(prefix) == 0) {
            var times int = shortLen/len(prefix)
            var temp = strings.Repeat(string(prefix), times)
            if (temp == chosenString) {
                times  = len(verifyString)/len(prefix)
                temp = strings.Repeat(string(prefix), times)
                if (temp == verifyString) {
                    return string(prefix)
                }
            }
        }
           
    }
    return ""
}