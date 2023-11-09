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

const GetEventDetails string = `SELECT
e.id,
e.title,
e.start_at,
e.end_at,
COUNT(w.id) AS total_workshops
FROM event_management.events AS e
LEFT JOIN event_management.workshops AS w ON e.id = w.event_id
WHERE e.id = ?
GROUP BY e.id;`

const WorkshopListByEventId string = `SELECT 
id, start_at, end_at, title, description
FROM
event_management.workshops
WHERE
event_id = ?
AND start_at > now();`

const GetActiveEventList string = `SELECT 
id, title, start_at, end_at
FROM
event_management.events
where start_at > now()
order by id desc
limit 10 offset ?;`

const GetActiveEventCount string = `SELECT 
COUNT(id)
FROM
event_management.events
WHERE
start_at > NOW();`
