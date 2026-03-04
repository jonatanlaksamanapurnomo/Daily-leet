class Solution {
    public int numSpecial(int[][] mat) {
        int [] rowCount = new int [mat.length];
        int [] colCount = new int [mat[0].length];

        // Hitung semua 1 per row dan per col dalam 1 loop
        for(int i = 0; i < mat.length; i++){
            for(int j = 0; j < mat[0].length; j++){
                rowCount[i] += mat[i][j];  // akumulasi per row
                colCount[j] += mat[i][j];  // akumulasi per col
            }
        }

        int count = 0;
        for(int i = 0; i < mat.length; i++){
            for(int j = 0; j < mat[0].length; j++){
                if(mat[i][j] == 1 && rowCount[i] == 1 && colCount[j] == 1){
                    count++;
                }
            }
        }

        return count;
    }
}