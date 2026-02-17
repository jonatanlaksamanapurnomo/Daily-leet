class Solution {
    public int totalFruit(int[] fruits) {
        HashMap<Integer,Integer> cache = new HashMap<>();
        int left = 0;
        int maxFruit = 0;

        for(int right = 0 ; right < fruits.length ; right++){

            cache.put(fruits[right] , cache.getOrDefault(fruits[right] , 0) +1);

            //Shrink window
               while (cache.size() > 2){
                cache.put(fruits[left] , cache.get(fruits[left])-1);
                if(cache.get(fruits[left]) <= 0 ){
                    cache.remove(fruits[left]);
                }
                left++;
            }
            //Calc response
            maxFruit = Integer.max(maxFruit , right - left +1);
            
        }


        return maxFruit;
    }
}