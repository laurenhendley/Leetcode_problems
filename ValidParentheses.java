class ValidParentheses {
    public boolean isValid(String s) {
        int l = s.length();
        Stack<Character> stk = new Stack<>();

        for(int i = 0; i < l; i++){
            Character c = s.charAt(i);

            if(c.equals("(") || c.equals("{") || c.equals("[")){
                stk.push(c);
            } else if(c.equals(")") || c.equals("}") || c.equals("]")){
                if(stk.isEmpty()) { return false; }

                Character c2 = stk.pop();

                if( !(( c2.equals("[") && c.equals("]") ) ||  ( c2.equals("(") && c.equals(")") ) || ( c2.equals("{") && c.equals("}") ) )){
                    return false;
                }

            }
        }

        return stk.isEmpty();
    }
}