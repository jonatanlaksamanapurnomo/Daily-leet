class Solution {
    public String findDifferentBinaryString(String[] nums) {
        StringBuilder diagonal = new StringBuilder();
        for(int i = 0; i < nums.length; i++){
            String currentNum = nums[i];
            for(int j = 0 ; j<currentNum.length() ; j++){
                if(j == i){
                    diagonal.append(flipChar(currentNum.charAt(j)));
                    break;
                }
            }
        }
        return diagonal.toString();
    }

    public char flipChar(char ch){
        if(ch == '1') return '0';
        return '1';
    }
}
