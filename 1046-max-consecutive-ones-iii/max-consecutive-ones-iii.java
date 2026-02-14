class Solution {
    public int longestOnes(int[] nums, int k) {
        int left = 0;
        int zeroCounter = 0;
        int longestOneResp = 0;

        //Expand window
        for(int right = 0 ; right < nums.length ; right++){
            //Caluclate shrink flag
            if(nums[right] == 0){
                zeroCounter++;
            }

            //Shrink logic
            while(zeroCounter > k){
                if(nums[left] == 0 ){
                    zeroCounter--;
                }
                left++;
            }
            // calculate resp base on current window size
            longestOneResp = Integer.max(longestOneResp , right - left + 1);
        }

        return longestOneResp;
    }
}