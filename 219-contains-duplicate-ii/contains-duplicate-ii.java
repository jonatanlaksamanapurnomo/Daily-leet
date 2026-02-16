class Solution {
      public boolean containsNearbyDuplicate(int[] nums, int k) {
        int left = 0;
        HashMap<Integer,Integer> cache = new HashMap<>();

        for(int right = 0 ; right < nums.length ; right++){

            //Shrink window if already > k
            if(right - left + 1 > k+1){
                if(cache.containsKey(nums[left])){
                    cache.put(nums[left] , cache.get(nums[left]) - 1);
                    left++;
                }
            }

           //Check elm already exist or not
            int currentElm = nums[right];
            if(cache.containsKey(currentElm) && cache.getOrDefault(currentElm,0) >= 1) {
                return true;
            }

            cache.put(nums[right] ,1);
        }

        return false;
    }
}