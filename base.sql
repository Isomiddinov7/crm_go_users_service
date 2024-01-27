SELECT
  	s.full_name,
    s.phone,
    s.paid_sum,
    COUNT(c.id) AS course_count,
    (s.paid_sum * COUNT(c.id)) AS total_sum
FROM
    student as s
JOIN
    "group" as g ON s.group_id = g.id
JOIN
    "group" as c ON g.id = c.id
WHERE
    s.branch_id = '4e894e71-0b02-4fe6-8598-fe63bd6b56c4'
GROUP BY
    s.id, s.full_name, s.phone, s.paid_sum;



SELECT
	full_name,
	phone,
	salary,
	months_worked,
	(salary * months_worked) AS total_sum,
	ielts
FROM teacher
WHERE branch_id = '4e894e71-0b02-4fe6-8598-fe63bd6b56c4'

SELECT 
    t.full_name as teacher,
    ARRAY_AGG(s.full_name) as student,
    COUNT(s.full_name) as student_count
FROM
    "teacher" as t
JOIN  "group" as g ON g.teacher_id = t.id
JOIN "student" as s ON s.group_id = g.id 
GROUP BY t.full_name;


