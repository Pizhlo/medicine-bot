CREATE VIEW drugs.drugs_view AS 
SELECT id, user_id, row_number()over(partition by user_id order by id) AS drug_number
FROM drugs.drugs;