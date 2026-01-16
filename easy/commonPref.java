package easy;
class commonPref {
    public String longestCommonPrefix(String[] strs) {
        String pref = strs[0];

        for(int i = 0; i < strs.length; i++){
            while(!strs[i].startsWith(pref)){
                pref = pref.substring(0, pref.length() - 1);
                if(pref.isEmpty()) return "";
            }
        }

        return pref;
    }
}