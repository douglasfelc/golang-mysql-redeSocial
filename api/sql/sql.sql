CREATE DATABASE IF NOT EXISTS redeSocial;
USE redeSocial;

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';
GRANT ALL PRIVILEGES ON redeSocial.* TO 'golang'@'localhost';

DROP TABLE IF EXISTS `posts`;
DROP TABLE IF EXISTS `followers`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` INT auto_increment PRIMARY KEY,
  `name` VARCHAR(50) NOT NULL,
  `nick` VARCHAR(50) NOT NULL UNIQUE,
  `email` VARCHAR(50) NOT NULL UNIQUE,
  `password` VARCHAR(100) NOT NULL,
  `createdAt` TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;

CREATE TABLE `followers` (
  `user_id` INT NOT NULL,
  /* I define foreign key so that only valid users.id enter this table */
  FOREIGN KEY(user_id) /* Set the key name */
  REFERENCES users(id) /* Reference to column */
  ON DELETE CASCADE, /* Deletes records when the referenced user is deleted */

  `follower_id` INT NOT NULL,
  FOREIGN KEY(follower_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  /* Composite primary key: this way I guarantee that there will never be a duplicate record in the table */
  PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE `posts` (
  `id` INT auto_increment PRIMARY KEY,
  `title` VARCHAR(100) NOT NULL,
  `content` TEXT NOT NULL,

  `author_id` INT NOT NULL,
  FOREIGN KEY(author_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  `likes` INT DEFAULT 0,
  `createdAt` TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;
