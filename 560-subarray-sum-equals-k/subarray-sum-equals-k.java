class Solution {
    public int subarraySum(int[] nums, int k) {
        int sum = 0;
        HashMap<Integer, Integer> cache = new HashMap<>();
        int ans = 0;

        cache.put(0, 1);

        for (int i = 0; i < nums.length; i++) {
            sum += nums[i];

            ans += cache.getOrDefault(sum - k, 0);

            cache.put(sum, cache.getOrDefault(sum, 0) + 1);
        }

        return ans;
    }
}
