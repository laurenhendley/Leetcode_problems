import java.util.ArrayList;

class PalindromeNum {
    // Done using strings
    public boolean isPalindrome(int x) {
        String ns = Integer.toString(x);
        int lgth = ns.length() - 1;
        int i = 0;

        while(i <= lgth){
            if(ns.charAt(i) != ns.charAt(lgth)){
                return false;
            }
            i++;
            lgth--;
        }
        return true;
    }
}