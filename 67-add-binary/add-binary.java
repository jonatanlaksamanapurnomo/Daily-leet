class Solution {
    public String addBinary(String a, String b) {
        int lastIdxA = a.length()-1;
        int lastIdxB = b.length()-1;
        int carry = 0;
        StringBuilder sb = new StringBuilder();

        while(lastIdxA >= 0 || lastIdxB >= 0 || carry > 0){
            int aElm = lastIdxA >= 0 ? a.charAt(lastIdxA) - '0' : 0;
            int bElm = lastIdxB >= 0 ? b.charAt(lastIdxB) - '0' : 0;

            int sum = ( carry + aElm + bElm) % 2;
            carry = (carry + aElm + bElm) / 2;

            sb.append(Integer.toString(sum));

            if(lastIdxA >= 0){
                lastIdxA--;
            }

            if(lastIdxB >= 0){
                lastIdxB--;
            }


        }
        return sb.reverse().toString();
    }
}