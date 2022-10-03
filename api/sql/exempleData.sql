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

INSERT INTO posts (`title`, `content`, `author_id`) VALUES
("Post 1 of User 1", "Content 1 of User 1", 1),
("Post 2 of User 1", "Content 2 of User 1", 1),
("Post 3 of User 1", "Content 3 of User 1", 1),
("Post 1 of User 2", "Content 1 of User 2", 2),
("Post 1 of User 3", "Content 1 of User 3", 3);