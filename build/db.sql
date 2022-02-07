CREATE DATABASE  IF NOT EXISTS `unohouse` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
use `unohouse`
DROP TABLE IF EXISTS `notification`;
CREATE TABLE `notification` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `body` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `click_action` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `channel_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data` json DEFAULT NULL,
  `device_id` int NOT NULL DEFAULT 0,
  `sent` tinyint DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_BF5476CAA76ED395` (`user_id`),
  CONSTRAINT `FK_BF5476CAA76ED395` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `notification` VALUES 
(1,1,'Registration','New User registered','Notifications','270','{\"userId\": \"2\"}', 69,1,'2020-11-25 11:21:54',NULL),
(2,1,'Login','User logged in','Notifications','220','{\"userId\": \"2\"}',69,1,'2020-11-25 11:30:47',NULL),
(3,1,'Change','User another_user changes sensor settings','Admin','110','{\"sensor\": \"27\", \"userId\": \"2\"}', 69, 0,'2020-11-25 11:31:33',NULL);

DROP TABLE IF EXISTS `app_description`;
CREATE TABLE `app_description` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `body` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `app_description` VALUES (
  1, 'main','This main section describes app. You can use it freely for interview purposes.', '2022-02-01 12:00:00'
)

CREATE DATABASE  IF NOT EXISTS `users` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
use `users`

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(180) COLLATE utf8mb4_unicode_ci NOT NULL,
  `enabled` tinyint NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_login` datetime DEFAULT NULL,
  `roles` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '(DC2Type:array)',
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `firebase_token` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQ_8D93D64992FC23A8` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `user` VALUES (
  1,
  'user1',
  1,
  '$2a$13$idXK61n71KvV26xsXOt8F.zHlQjquyfLdVHPlBMDeb.3DdpL.aRDO',
  '2022-02-04 09:16:12',
  'a:1:{i:0;s:16:"ROLE_SUPER_ADMIN";}',
  'TestName1',
  'TestLastName1',
  '2021-11-01 22:14:09',
  NULL,
  'Admin account token',
  NULL),
  (
  2,
  'another_user',
  1,
  '$2a$13$idXK61n71KvV26xsXOt8F.zHlQjquyfLdVHPlBMDeb.3DdpL.aRDO',
  '2022-02-06 20:46:12',
  'a:1:{i:0;s:13:\"ROLE_EMPLOYEE\";}',
  'Multi',
  'Roles account',
  '2021-11-01 22:14:09',
  NULL,
  'Multi Roles account token',
  NULL),
  (3,
  'test_user2',
  1,
  '$2a$13$TQWml9TszLLVuZ62rMjbxeXSkJWeiYN51tAp4Dnb3zq6Q/c4OM196',
  '2020-12-08 09:25:32',
  'a:1:{i:0;s:9:"ROLE_USER";}',
  'FirstName2',
  'LastName2',
  '2020-11-02 04:33:20',
  '2022-01-22 19:11:34',
  'Token of normal user',
  NULL),
  (5,
  'incognito',
  1,
  '$2y$13$r25oAW8DjZmWUuKWK/Lt4eM3WhSJ.PabHO4.HGtQ/O2zZpOLrE5QG',
  '2021-11-04 19:23:40',
  'a:1:{i:0;s:9:"ROLE_USER";}',
  'incognito',
  'incognito',
  '2020-02-04 10:41:59',
  '2020-11-02 10:41:59',
  'Token with deleted user, should NOT BE LISTED ANYWHERE!',
  '2020-11-02 10:41:59');

DROP TABLE IF EXISTS `user_device`;
CREATE TABLE `user_device` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `firebase_token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `os_type` int NOT NULL,
  `sdk_version` int NOT NULL,
  `model` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `brand` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `last_login` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_vals` (`user_id`,`firebase_token`),
  KEY `IDX_6C7DADB3A76ED395` (`user_id`),
  CONSTRAINT `FK_6C7DADB3A76ED395` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `user_device` VALUES (
  69, 1, 'token #1 of user1', 1, 30, 'Mi 9 SE', 'Xiaomi', '2022-01-27 14:57:36', NULL, '2022-01-28 11:30:17'),
  (355,1,'token #3 for user #1',1, 30, 'M200000', 'POCO', '2022-02-04 11:21:35', NULL, '2022-02-04 11:55:26'),
  (361,1,'another token for UserId 1', 1, 26,'Android SDK built for x86', 'google', '2022-02-04 11:56:44', NULL, '2022-02-04 12:00:57'),
  (365,1,'System will hold more (tokens!) A. Diatlow', 1, 28, 'Mi A2 Lite', 'xiaomi', '2022-02-04 12:19:53', NULL, '2022-02-04 12:33:13'),
  (72, 1, 'Second token of user #1', 1, 26, 'Android SDK built for x86', 'google', '2022-01-27 18:00:31', NULL, '2022-01-28 09:49:25'),
  (87, 2, 'First token of user #2', 1, 29, 'motorola one action', 'motorola',' 2022-01-28 09:29:13', NULL, '2022-02-04 14:28:09'),
  (93,2,'Second token for user #2', 1, 26, 'Android SDK built for x86', 'google', '2022-01-28 11:44:29', NULL, '2022-02-03 12:30:01'),
  (98, 2, 'Third token for user #2', 1, 26,'Android SDK built for x86', 'google', '2022-01-28 11:54:11', NULL, '2022-01-28 12:10:13'),
  (99,3,'First token for user #3', 1, 31, 'LE2121', 'OnePlus', '2022-01-28 11:54:18', NULL, '2022-01-30 02:47:15');