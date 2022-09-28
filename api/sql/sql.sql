CREATE DATABASE IF NOT EXISTS redeSocial;
USE redeSocial;

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';
GRANT ALL PRIVILEGES ON redeSocial.* TO 'golang'@'localhost'

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` INT auto_increment PRIMARY KEY,
  `name` VARCHAR(50) NOT NULL,
  `nick` VARCHAR(50) NOT NULL UNIQUE,
  `email` VARCHAR(50) NOT NULL UNIQUE,
  `password` VARCHAR(20) NOT NULL,
  `createdAt` TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;
