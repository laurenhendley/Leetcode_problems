package easy;
import java.util.HashMap;

class RomanInt {
    public int romanToInt(String s) {
        int ans = 0;
        HashMap<Character, Integer> ri = new HashMap<>();

        ri.put('I', 1);
        ri.put('V', 5);
        ri.put('X', 10);
        ri.put('L', 50);
        ri.put('C', 100);
        ri.put('D', 500);
        ri.put('M', 1000);


        for(int i = 0; i < s.length();i++){
            char c = s.charAt(i);
            int val = ri.get(c);

            if(i+1 < s.length()){
                char cNext = s.charAt(i+1);
                if(c == 'I' && cNext == 'V'){ val = 4; i++; }
                else if(c == 'I' && cNext == 'X'){ val = 9; i++; }
                else if(c == 'X' && cNext == 'L'){ val = 40; i++; }
                else if(c == 'X' && cNext == 'C'){ val = 90; i++; }
                else if(c == 'C' && cNext == 'D'){ val = 400; i++; }
                else if(c == 'C' && cNext == 'M'){ val = 900; i++; }
            }

            ans += val;
        }



        return ans;
    }
}