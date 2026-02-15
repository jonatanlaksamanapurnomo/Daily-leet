class Solution {
    public int maxVowels(String s, int k) {
        int maxVowel = 0;
        int left = 0;

        for(int i = 0; i < k ; i++){
            char currentChar = s.charAt(i);
            if(isCharVowel(currentChar)){
                maxVowel++;
            }
        }

        int currentVowel = maxVowel;
        for(int right = k ; right < s.length() ; right++){
            int add = isCharVowel(s.charAt(right)) ? 1 : 0;
            int remove = isCharVowel(s.charAt(left)) ? 1:0;
            currentVowel += add - remove;
            maxVowel = Integer.max(maxVowel,currentVowel);
            left++;
        }

        return maxVowel;
    }

    public boolean isCharVowel(char currentChar){
        if(currentChar == 'a' || currentChar == 'i' || currentChar == 'e' || currentChar == 'o' || currentChar=='u'){
            return true;
        }

        return false;
    }
}