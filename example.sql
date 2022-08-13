CREATE DATABASE IF NOT EXISTS `docker` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `docker`;

DROP TABLE IF EXISTS `users`;


mysql> show create table users\G
*************************** 1. row ***************************
       Table: users
Create Table: CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `email` longtext,
  `phone` longtext,
  `password` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_created_at` (`created_at`),
  KEY `idx_users_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
1 row in set (0.00 sec)

mysql> describe users;
+------------+---------------------+------+-----+---------+----------------+
| Field      | Type                | Null | Key | Default | Extra          |
+------------+---------------------+------+-----+---------+----------------+
| id         | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| name       | longtext            | YES  |     | NULL    |                |
| email      | longtext            | YES  |     | NULL    |                |
| phone      | longtext            | YES  |     | NULL    |                |
| password   | longtext            | YES  |     | NULL    |                |
| created_at | datetime(3)         | YES  | MUL | NULL    |                |
| updated_at | datetime(3)         | YES  | MUL | NULL    |                |
+------------+---------------------+------+-----+---------+----------------+
7 rows in set (0.00 sec)

mysql> insert into `users` (phone) 
values ("18018001801"),
       ("17845127845");

