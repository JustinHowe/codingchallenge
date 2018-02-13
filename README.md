##  Coding Challenge
Call a "word path" a sequence of words where each successive word differs from the last only by a single letter (added, removed, changed). Given a starting word and a list of words, one per file (as in /usr/dict/words), write a program to find the longest word path in a dictionary of words.

---

* **Sample**: Given list "a", "ab", "ac", "ad","abc", "ade" and word "abcd"

* **Expected result**: "abcd", "abc", "ab", "a", "ac", "ad", "ade"

---

The two files "coding_challenge.go" and "word_helpers.go" complete this assignment to specification. As the expected result asked for numerous values sorted on match completeness from the control word instead of just one value, I created a sorting mechanism that incremented a counter based on every difference that a word had from the control word. If it is different in length, increment by the number of characters missing or added. If index-based same characters from both words were not the same, increment again (and do so for each differing character). This allowed for the standard sorting in GoLang to re-arrange the slice of strings based on this sorting weight. 

**Note**: The specification called for, "...a list of words, one per file", however in the program I only used one file to write words to with commas as separators. For this activity, writing numerous files does not make sense and adds no value to the exercise. For such an effort that would gain value from doing so, one way to accomplish this is to call the saveToFile() function within a for loop that iterates over each word. To call each file, the same process can be done for readFromFile(), where it is called from an iterated for loop by the number of files found that contain a unique identifier (such as, "words_word-name.txt", where words_ is the unique identifier) and then append each containing word to the slice.



**Program output after running:**
![Coding Challenge Output](https://raw.githubusercontent.com/JustinHowe/codingchallenge/master/codingchallengeoutput.PNG)
