class Solution {
    public int[] sortByBits(int[] arr) {
        int [] resp = new int[arr.length];
        PriorityQueue<Integer> pq = new PriorityQueue<Integer>((a,b) -> {
            int bitsA = countOneInBit(a);
            int bitsB = countOneInBit(b);

            if(bitsA == bitsB){
                return a-b;
            }
            return bitsA - bitsB;
        });

        for(int i = 0 ; i<arr.length ; i++){
            pq.add(arr[i]);
        }

        int i = 0;
        while (pq.size() > 0){
            resp[i++] = pq.poll();
        }
        return resp;
    }

    public static int countOneInBit(int elm) {
        int totalOne = 0;

        while (elm > 0 ){
            if((elm & 1) == 1){
                totalOne++;
            }
            elm >>=1;
        }
        return totalOne;
    }
}