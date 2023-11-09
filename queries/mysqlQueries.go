package queries

const GetWorkShopDetails string = `SELECT
w.id,
w.title,
w.description,
w.start_at,
w.end_at,
COUNT(r.id) AS total_reservations
FROM workshops AS w
LEFT JOIN reservations AS r ON w.id = r.workshop_id
WHERE w.id = ?
GROUP BY w.id;`
