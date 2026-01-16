package easy;

class MinPointDist {
    public int minTimeToVisitAllPoints(int[][] points) {
        int dx = 0, dy = 0, sec = 0;

        for(int i = 0; i < points.length - 1; i++){
            dx = Math.abs(points[i][0] - points[i+1][0]);
            dy = Math.abs(points[i][1] - points[i+1][1]);

            sec += Math.max(dx,dy);
        }
        return sec;
    }
}