INSERT INTO users (email, password) VALUES ('john_doe@example.com', '123');
INSERT INTO sites (title, description) VALUES ('My new blog', 'It is amazing!');


INSERT INTO posts (body, site_id) VALUES ('This is the first post', 1);
INSERT INTO posts (body, site_id) VALUES ('This is the second post', 1);
INSERT INTO posts (body, site_id) VALUES ('Today I am going to the beach', 1);
INSERT INTO posts (body, site_id) VALUES ('The weather is nice today', 1);


INSERT INTO pages (title, body, site_id, slug) VALUES ('About', 'This is the about page', 1, 'about');
INSERT INTO pages (title, body, site_id, slug) VALUES ('Contact', 'This is the contact page', 1, 'contact');



INSERT INTO events (title, body, event_date, site_id) VALUES ('Birthday party', 'This is the birthday party', '2023-01-01', 1);
INSERT INTO events (title, body, event_date, site_id) VALUES ('Christmas party', 'This is the christmas party', '2023-12-25', 1);
INSERT INTO events (title, body, event_date, site_id) VALUES ('New year party', 'This is the new year party', '2024-01-01', 1);
INSERT INTO events (title, body, event_date, site_id) VALUES ('Halloween party', 'This is the halloween party', '2024-10-31', 1);
INSERT INTO events (title, body, event_date, site_id) VALUES ('Easter party', 'This is the easter party', '2024-04-20', 1);
INSERT INTO events (title, body, event_date, site_id) VALUES ('Valentine party', 'This is the valentine party', '2024-02-14', 1);
INSERT INTO events (title, body, event_date, site_id) VALUES ('St. Patrick''s Day', 'This is the St. Patrick''s Day', '2024-03-17', 1);
