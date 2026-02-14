class Solution {
    public int lengthOfLongestSubstring(String s) {
        int left = 0;
        Map<Character,Integer> cache = new HashMap<>();
        int longestSubstringLength = Integer.MIN_VALUE;

        for(int right = 0 ; right < s.length() ; right++){
            cache.put(s.charAt(right) , cache.getOrDefault(s.charAt(right) , 0) + 1);

             while (cache.get(s.charAt(right)) > 1){
                cache.put(s.charAt(left), cache.get(s.charAt(left)) - 1);
                left++;
            }


            longestSubstringLength = Integer.max(longestSubstringLength , right - left + 1);
        }

        if (longestSubstringLength == Integer.MIN_VALUE) {
            return 0;
        }

        return longestSubstringLength;
    }
}