CREATE TABLE IF NOT EXISTS `notification_data` (
                                     `id` int(11) NOT NULL,
                                     `notification_id` int(11) NOT NULL,
                                     `user_id` varchar(32) NOT NULL,
                                     `is_read` tinyint(4) NOT NULL,
                                     `is_deleted` tinyint(4) NOT NULL,
                                     `created_at` datetime NOT NULL,
                                     `read_at` datetime DEFAULT NULL,
                                     `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
ALTER TABLE `notification_data`
    ADD PRIMARY KEY (`id`);
ALTER TABLE `notification_data`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;
