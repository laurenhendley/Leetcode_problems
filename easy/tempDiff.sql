# Write your MySQL query statement below
select c1.id from Weather c1, Weather c2
where c1.temperature > c2.temperature and DATEDIFF(c1.recordDate, c2.recordDate) = 1
