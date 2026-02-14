class Solution {
     public int minSubArrayLen(int target, int[] nums) {
        int left = 0;
        int currentSum = 0;
        int minWindowSize = Integer.MAX_VALUE;
        for(int right = 0 ; right < nums.length ; right++){
            currentSum += nums[right];

            //shrink windows
            while (currentSum >= target){
                int currentWindowSize = (right - left)+1;
                minWindowSize = Integer.min(minWindowSize , currentWindowSize);
                currentSum -= nums[left++];
            }
        }

        if(minWindowSize == Integer.MAX_VALUE){
            return 0;
        }
        return minWindowSize;
    }
}