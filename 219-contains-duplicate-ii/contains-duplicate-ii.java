class Solution {
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        int left = 0;
        HashMap<Integer,Integer> cache = new HashMap<>();


        for(int i = 0 ; i<Math.min(k,nums.length) ; i++){
            if(cache.getOrDefault(nums[i] , 0 ) > 0){
                return true;
            }
            cache.put(nums[i] , cache.getOrDefault(nums[i] , 0) + 1);
        }

        for(int right = k ; right < nums.length ; right++){

            //Shrink window if already > k
            if(right - left + 1 > k+1){
                cache.put(nums[left] , cache.get(nums[left]) - 1);
                left++;
            }

            //Check elm already exist or not
            int currentElm = nums[right];
            if (cache.getOrDefault(currentElm, 0) > 0) {
                return true;
            }

            cache.put(nums[right] ,cache.getOrDefault(nums[right],0) + 1);
        }

        return false;
    }
}