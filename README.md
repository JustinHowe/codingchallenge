# codingchallenge
Call a "word path" a sequence of words where each successive word differs from the last only by a single letter (added, removed, changed). Given a starting word and a list of words, one per file (as in /usr/dict/words), write a program to find the longest word path in a dictionary of words.

Sample: Given list "a", "ab", "ac", "ad","abc", "ade" and word "abcd"
Expected result: "abcd", "abc", "ab", "a", "ac", "ad", "ade"

The two files "coding_challenge.go" and "word_helpers.go" complete this assignment to specification. As the expected result asked for numerous values sorted on match completeness from the control word instead of just one value, I created a sorting mechanism that incremented a counter based on every difference that a word had from the control word. If it is different in length, increment by the number of characters missing or added. If index-based same characters from both words were not the same, increment again (and do so for each differing character). This allowed for the standard sorting in GoLang to re-arrange the slice of strings based on this sorting weight. 

Program output after running:
![Coding Challenge Output](https://raw.githubusercontent.com/JustinHowe/codingchallenge/master/codingchallengeoutput.png)
