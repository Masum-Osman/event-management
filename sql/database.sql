create database event_management;
use event_management;

CREATE TABLE events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_at DATETIME NOT NULL,
    end_at DATETIME NOT NULL
);

CREATE TABLE workshops (
    id INT AUTO_INCREMENT PRIMARY KEY,
    event_id INT,
    start_at DATETIME NOT NULL,
    end_at DATETIME NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE reservations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    workshop_id INT
);


-- Insert data into the "events" table
INSERT INTO events (title, start_at, end_at)
VALUES
    ('Event 1', '2023-10-31 09:00:00', '2023-10-31 17:00:00'),
    ('Event 2', '2023-11-01 10:00:00', '2023-11-01 18:00:00'),
	('Event 3', '2023-11-02 08:00:00', '2023-11-02 16:00:00'),
    ('Event 4', '2023-11-03 10:00:00', '2023-11-03 18:00:00');

-- Insert data into the "workshops" table
INSERT INTO workshops (event_id, start_at, end_at, title, description)
VALUES
    (1, '2023-10-31 10:00:00', '2023-10-31 12:00:00', 'Workshop A', 'Description of Workshop A'),
    (1, '2023-10-31 14:00:00', '2023-10-31 16:00:00', 'Workshop B', 'Description of Workshop B'),
	(2, '2023-11-02 11:00:00', '2023-11-02 13:00:00', 'Workshop C', 'Description of Workshop C'),
    (2, '2023-11-02 15:00:00', '2023-11-02 17:00:00', 'Workshop D', 'Description of Workshop D'),
    (3, '2023-11-03 09:00:00', '2023-11-03 11:00:00', 'Workshop E', 'Description of Workshop E'),
    (3, '2023-11-03 14:00:00', '2023-11-03 16:00:00', 'Workshop F', 'Description of Workshop F');

-- Insert data into the "reservations" table
INSERT INTO reservations (name, email, workshop_id)
VALUES
    ('John Doe', 'john@example.com', 1),
    ('Jane Smith', 'jane@example.com', 2),
	('Alice Johnson', 'alice@example.com', 3),
    ('Bob Williams', 'bob@example.com', 4),
    ('Eva Davis', 'eva@example.com', 5),
    ('George Wilson', 'george@example.com', 6);