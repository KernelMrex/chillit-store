DROP DATABASE IF EXISTS `chillit_store`;
CREATE DATABASE `chillit_store`;
ALTER DATABASE `chillit_store` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `chillit_store`;

CREATE TABLE IF NOT EXISTS `city`
(
    `id`    BIGINT NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(64) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `place`
(
    `id`          BIGINT NOT NULL AUTO_INCREMENT,
    `title`       VARCHAR(64) NOT NULL,
    `address`     VARCHAR(64) NOT NULL,
    `city_id`     BIGINT NOT NULL,
    `description` TEXT NOT NULL,
    `image_url`   VARCHAR(255),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`city_id`) REFERENCES `city`(`id`),
    CONSTRAINT `place_info` UNIQUE(`title`, `address`, `city_id`)
);
