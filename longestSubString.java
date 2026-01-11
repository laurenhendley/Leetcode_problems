import java.util.HashSet;
import java.util.Set;

class longestSubString {
    public int lengthOfLongestSubstring(String s) {
        int maxL = 0;
        int l = 0;
        Set<Character> cSet = new HashSet<>();

        for(int r = 0; r < s.length(); r++){
            while(cSet.contains(s.charAt(r))){
                cSet.remove(s.charAt(l));
                l++;
            }
            cSet.add(s.charAt(r));
            maxL = Math.max(maxL, r-l+1);
        }

        return maxL;
    }
}