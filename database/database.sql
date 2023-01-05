--database
CREATE DATABASE IF NOT EXISTS `apigin` 
USE `apigin`;

-- users table
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `last_name` varchar(100) NOT NULL DEFAULT '',
  `age` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
)
