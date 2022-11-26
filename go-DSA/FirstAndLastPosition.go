/* https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/*/

/* 
	In this question we have to find position of the no which are same
	Ex:- arr = {2,3,4,5,5,5,5,6,7} target = 5  ans = [4,7] this should be the ans
	if not found then return [-1,-1]

*/

package godsa

/* So to slove this problem this we will use to binary search first one for the first positon and second is for last postion on the target element  */
func searchRange(nums []int, target int) []int {
    var ans = []int{-1,-1} /* this will be the ans if you did't find the target lement  */
     /* For the first position of the target element */
     var start int = search(nums, target,true) /* we are call the func with vale and we will get the target element first position */
     var end int = search(nums,target,false) /* we are calling the same func but with bool value false so that it will return last position of the targetn elemetn  */

    /* now we have assing the element those values to ans varable */
     ans[0] = start
     ans[1] = end

     return ans

}

func search(nums []int, target int, findstartindex bool) int{
     var ans int = -1
     var start int =0
     var end int = len(nums) -1 
     for start <= end{
         var mid int = start + (end - start)/2

         if target > nums[mid]{
             start = mid + 1
         } else if target < nums[mid]{
             end = mid - 1
         }else{
             /* here what we are doing is we have found the target element but there is possbility that this element is not the starting element there may be another same number just to it. This is can happen to the both starting ans ending index  */
             ans = mid
             if(findstartindex){ /* if we are finding for the first no. then we will change the end valuse as we are doing will finding the target element and it will keep doing it untitl we find the first position of the target element  */
                 end = mid -1
             }else{
                 start = mid + 1 /* this is for the last element this will keep excutin until we find the last element for the target element */
             }
         }
     }
     return ans
}