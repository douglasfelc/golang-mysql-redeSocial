/* Sample data for testing */
/* The unhashed value of passwords is 123456 */

INSERT INTO users (`name`, `nick`, `email`, `password`) VALUES 
("Alex Green", "alexg", "alexgreen@email.com", "$2a$10$zyJnI6rc1RKv942mBV1fzeKEi.Q5NREUEbxzDGjhZY51nh0mh6YzG"),
("Emily Brown", "ebrown", "ebrown@email.com", "$2a$10$zyJnI6rc1RKv942mBV1fzeKEi.Q5NREUEbxzDGjhZY51nh0mh6YzG"),
("Jason Red", "jsonred", "jsonred@email.com", "$2a$10$zyJnI6rc1RKv942mBV1fzeKEi.Q5NREUEbxzDGjhZY51nh0mh6YzG");

INSERT INTO followers (user_id, follower_id) VALUES 
(1, 2),
(3, 1),
(1, 3);