class Solution {
    static int countPrimeSetBits(int left, int right) {
        int resp = 0;
        for (int i = left; i <= right; i++) {
            int currTotalOne = countOnes(i);
            if (isPrime(currTotalOne)) {
                resp++;
            }
        }

        return resp;
    }

    private static int countOnes(int n) {
        int count = 0;
        while (n > 0) {
            count += n & 1;
            n >>= 1;
        }

        return count;
    }

    private static boolean isPrime(int n) {
        if (n <= 1) return false;
        if (n <= 3) return true;
        if (n % 2 == 0 || n % 3 == 0) return false;
        for (int i = 5; i * i <= n; i += 6) if (n % i == 0 || n % (i + 2) == 0) return false;
        return true;
    }
}