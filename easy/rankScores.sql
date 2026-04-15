# Write your MySQL query statement below
select S.score, count(S2.score)  as 'rank' from Scores S,
(select distinct score from Scores) as S2 where S.score <= s2.score 
group by S.id order by S.score desc
