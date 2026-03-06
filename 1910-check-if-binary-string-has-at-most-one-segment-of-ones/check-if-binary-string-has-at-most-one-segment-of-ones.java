class Solution {
    public boolean checkOnesSegment(String s) {
        boolean seenZero = false;
        for(int i = 0; i < s.length(); i++) {
            if(s.charAt(i) == '0') {
                seenZero = true;
            } else {
                if(seenZero) {
                    return false;
                }
            }
        }
        return true;
    }
}
