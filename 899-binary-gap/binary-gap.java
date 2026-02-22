class Solution {
    public int binaryGap(int n) {
        int resp =0;
        int lastOnepos = 0;
        int currentPos = 1;
        while (n > 0){
            if((n&1) == 1){
                if(lastOnepos > 0 ){
                    //101
                    resp = Integer.max(resp , currentPos - lastOnepos);
                }
                lastOnepos = currentPos;
            }
            n >>=1;
            currentPos++;
        }

        return resp;
    }
}