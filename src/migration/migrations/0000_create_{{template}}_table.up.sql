CREATE TABLE IF NOT EXISTS `campaigns`
(
    `id`                   BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `code`                 VARCHAR(255),
    `name`                 VARCHAR(255),
    `pre_camp_info`        TEXT,
    `in_camp_info`         TEXT,
    `post_camp_info`       TEXT,
    `rule_expression`      VARCHAR(2000),
    `start_pre_camp_time`  BIGINT UNSIGNED,
    `start_in_camp_time`   BIGINT UNSIGNED,
    `start_post_camp_time` BIGINT UNSIGNED,
    `end_time`             BIGINT UNSIGNED,
    `created_at`           DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`           DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX `campaigns_code_idx` ON `campaigns` (`code`);
