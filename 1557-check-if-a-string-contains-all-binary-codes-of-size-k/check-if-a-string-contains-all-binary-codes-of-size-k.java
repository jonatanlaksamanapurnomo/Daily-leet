class Solution {
    public boolean hasAllCodes(String s, int k) {
        int left = 0;
        HashSet<String> set = new HashSet<>();

        for(int right = k-1 ; right < s.length() ; right++){
            String current = s.substring(left , right+1);
            set.add(current);
            if(set.size() == Math.pow(2,k)){
                return true;
            }
            left++;
        }

        return set.size() == Math.pow(2,k) ? true : false;
    }
}