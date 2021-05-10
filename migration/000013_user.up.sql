CREATE TABLE IF NOT EXISTS `user` (
            `id` int(11) NOT NULL AUTO_INCREMENT,
            `first_name` varchar(50) NOT NULL,
            `last_name` varchar(50) NOT NULL,
            `username` varchar(50) NOT NULL,
            `email` varchar(50) NOT NULL,
            `password` varchar(500) NOT NULL,
            `created_at` datetime NOT NULL,
            `updated_at` datetime DEFAULT NULL, 
            PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
