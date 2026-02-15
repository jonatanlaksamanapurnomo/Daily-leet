class Solution {
    public double findMaxAverage(int[] nums, int k) {
        double currentWindowSum = 0;
        int left = 0;

        for(int i = 0 ; i < k ;i++){
            currentWindowSum += nums[i];
        }

        double maxSum = currentWindowSum;
        for(int right = k ; right < nums.length ; right++){
            currentWindowSum += nums[right] - nums[left];
            left++;
            maxSum = Double.max(maxSum,currentWindowSum);
        }

        return maxSum/k;
    }
}