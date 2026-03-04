class Solution {
    public int numSpecial(int[][] mat) {
        int resp = 0;

        for(int i = 0; i < mat.length; i++){

            for(int j = 0; j < mat[0].length; j++){
                if(mat[i][j] == 1){
                    int totalRow = 0;
                    int totalCol = 0;
                    //check row and col
                    for(int row = 0; row < mat.length; row++){
                        totalRow += mat[row][j];
                    }

                    for(int col = 0; col < mat[0].length; col++){
                        totalCol += mat[i][col];
                    }

                    if (totalRow == 1 && totalCol == 1) {
                        resp++;
                    }
                }
            }
        }
        return resp;
    }
}