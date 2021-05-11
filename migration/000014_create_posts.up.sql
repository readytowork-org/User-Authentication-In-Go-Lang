CREATE TABLE IF NOT EXISTS `post`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `body`  varchar(500) NOT NULL,
    `author_id` int(100)NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime DEFAULT NULL, 
    PRIMARY KEY (`id`),
    FOREIGN KEY(`author_id`)
        REFERENCES `user` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;