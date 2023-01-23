CREATE TABLE IF NOT EXISTS `locked-slots`
(
    `id`                   BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `code`                 VARCHAR(255),
    `name`                 VARCHAR(255),
    `created_at`           DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`           DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX `locked-slots_code_idx` ON `locked-slots` (`code`);
