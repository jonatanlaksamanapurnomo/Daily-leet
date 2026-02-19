class Solution {
    public int countBinarySubstrings(String s) {
        ArrayList<Integer> group = new ArrayList<>();
        int currentGroupLen = 1;
        int resp = 0;

        for(int i = 1; i<s.length();i++){
            char prev = s.charAt(i-1);
            char current = s.charAt(i);

            if(prev == current){
                currentGroupLen++;
            }
            else {
                group.add(currentGroupLen);
                currentGroupLen = 1;
            }
        }
        group.add(currentGroupLen);

        for(int i = 1; i<group.size() ; i++){
            int current = group.get(i);
            int prev = group.get(i-1);
            resp += Integer.min(current , prev);
        }
        return resp;
    }
}