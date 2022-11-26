/* https://leetcode.com/problems/find-smallest-letter-greater-than-target/ */
/* 
	So in this quation what we have do is the we use the binary search algo to slove the problem 
	everything is just same as the binary search algo we are just returning the letters[start % len(letters)]
	this will 

*/
package godsa

func nextGreatestLetter(letters []byte, target byte) byte {
    var start int = 0
    var end int = len(letters) - 1

    for start <= end{
        var mid int = start + (end - start) /  2

        if target < letters[mid]{
            end = mid-1
        }else{
            start = mid+1
        }
    }
    return letters[start % len(letters)]
}