package easy;

class Solution {

    public class TreeNode {
       int val;
       TreeNode left;
       TreeNode right;
       TreeNode() {}
       TreeNode(int val) { this.val = val; }
       TreeNode(int val, TreeNode left, TreeNode right) {
           this.val = val;
           this.left = left;
           this.right = right;
       }
   }

   
    public boolean isSymmetric(TreeNode root) {
        if(root == null){
            return true;
        }

        return isMirror(root.left, root.right);
    }

    private boolean isMirror(TreeNode l, TreeNode r){
        if(r == null && l == null){
            return true;
        }

        if(r == null || l == null){
            return false;
        }

        if(r.val != l.val){
            return false;
        }

        return isMirror(l.left, r.right) && isMirror(l.right, r.left);
    }
}
