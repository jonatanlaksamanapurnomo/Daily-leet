class Solution {
    public boolean hasAlternatingBits(int n) {
        int alternate =  (n >> 1);
        while(alternate > 0){
            int lastAlternateBit = alternate & 1;
            int lastNBit = n & 1;
            int bit = lastAlternateBit ^ lastNBit;
            if(bit != 1){
                return false;
            }
            alternate >>=1;
            n>>=1;
        }
        return true;
    }
}
