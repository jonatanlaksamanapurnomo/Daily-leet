class Solution {
    public double champagneTower(int poured, int query_row, int query_glass) {
        double[][] dp = new double[101][101];
        dp[0][0] = poured;

        for(int row = 0 ; row < 100 ; row++){
            for(int col = 0 ; col <= row ; col++){
                if(dp[row][col] > 1 ){
                    double overflow = dp[row][col] - 1;
                    dp[row][col] = 1;

                    double half = overflow / 2;
                    dp[row+1][col] += half;
                    dp[row+1][col+1] += half;

                }
            }
        }
        return dp[query_row][query_glass];
    }
}