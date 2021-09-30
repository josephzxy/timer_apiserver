CREATE TABLE IF NOT EXISTS `timer` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    `name` VARCHAR(255) NOT NULL,
    `trigger_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_timer_deleted_at` (`deleted_at`),
    KEY `idx_timer_trigger_at` (`trigger_at`),
    CONSTRAINT `uniq_timer_name` UNIQUE (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;
