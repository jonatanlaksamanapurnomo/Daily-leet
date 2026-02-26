class Solution {
    public int numSteps(String s) {
        int count = 0;
        while (s.length() > 1) {
            if(s.charAt(s.length()-1) == '1') {
                s = addOne(s);
            } else {
                s = divideTwo(s);
            }
            count++;
        }
        return count;
    }

    public static String addOne(String s){
        char[] arr = s.toCharArray();
        for(int i = s.length()-1; i >= 0; i--){
            if(s.charAt(i) == '0'){
                arr[i] = '1';
                return new String(arr);
            }
            else{
                arr[i] = '0';
            }
        }

        return "1" + new String(arr);
    }

    public static String divideTwo(String bits) {
        return bits.substring(0, bits.length() - 1);
    }
}