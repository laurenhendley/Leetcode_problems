package hard;

class ArrayMedian {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int total = (nums1.length + nums2.length);
        int mid = total / 2;

        int curr = 0, prev = 0, cnt = 0, cur2 = 0, cur1 = 0;

        while( cnt <= mid ){
            prev = curr;

            if( cur1 < nums1.length && (cur2 >= nums2.length || nums1[cur1] <= nums2[cur2])){
                curr = nums1[cur1++];
            }
            else {
                curr = nums2[cur2++];
            }

            cnt++;
        }

        if(total % 2 == 0){
            return (prev + curr) / 2.0; 
        } else {
            return curr; 
        }
    }
}
