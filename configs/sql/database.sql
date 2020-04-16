CREATE TABLE IF NOT EXISTS `place`
(
    `id`          BIGINT NOT NULL AUTO_INCREMENT,
    `title`       VARCHAR(64) NOT NULL,
    `address`     VARCHAR(64) NOT NULL ,
    `description` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

#DROP TABLE `place`;

