class Solution {
    public int minPartitions(String n) {
        int maxNumber = 0;

        for(int i = 0 ; i<n.length() ; i++){
            int currentValue = Integer.parseInt(n.charAt(i) + "");
            maxNumber = Integer.max(maxNumber , currentValue);
        }

        return maxNumber;
    }
}
