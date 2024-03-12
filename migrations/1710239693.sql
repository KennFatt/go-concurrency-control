-- Active: 1679834974662@@127.0.0.1@3306@flights
CREATE TABLE `seats` (`id` bigint NOT NULL AUTO_INCREMENT, `is_booked` bool NOT NULL DEFAULT false, `passenger_name` varchar(255) NULL, `version` bigint unsigned NULL DEFAULT 0, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
