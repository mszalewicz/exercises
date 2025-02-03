/*

Table: Students

+---------------+---------+
| Column Name   | Type    |
+---------------+---------+
| student_id    | int     |
| student_name  | varchar |
+---------------+---------+

student_id is the primary key (column with unique values) for this table.
Each row of this table contains the ID and the name of one student in the school.

Table: Subjects

+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| subject_name | varchar |
+--------------+---------+

subject_name is the primary key (column with unique values) for this table.
Each row of this table contains the name of one subject in the school.

Table: Examinations

+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| student_id   | int     |
| subject_name | varchar |
+--------------+---------+

There is no primary key (column with unique values) for this table. It may contain duplicates.
Each student from the Students table takes every course from the Subjects table.
Each row of this table indicates that a student with ID student_id attended the exam of subject_name.

Write a solution to find the number of times each student attended each exam.

Return the result table ordered by student_id and subject_name.

*/

SELECT
    t1.student_id,
    t1.student_name,
    t2.subject_name,
    COALESCE(COUNT(t3.subject_name),0) AS attended_exams
FROM
    Students as t1
CROSS JOIN
    Subjects as t2
LEFT JOIN
    Examinations as t3
ON
    t1.student_id = t3.student_id AND
    t2.subject_name = t3.subject_name
GROUP BY
    t1.student_id,
    t1.student_name,
    t2.subject_name
ORDER BY
    t1.student_id,
    t2.subject_nam
;