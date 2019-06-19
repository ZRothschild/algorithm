CREATE TABLE `users` (
     `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
     `name` VARCHAR(225) NULL DEFAULT '' COMMENT '名字' COLLATE 'utf8mb4_bin',
     `age` TINYINT(3) UNSIGNED NULL DEFAULT '18' COMMENT '年龄',
     `update_time` INT(11) UNSIGNED NULL DEFAULT '0',
     `add_time` INT(11) UNSIGNED NULL DEFAULT '0',
     PRIMARY KEY (`id`)
)
    COMMENT='用户表'
    COLLATE='utf8mb4_bin'
    ENGINE=InnoDB
;

CREATE TABLE `courses` (
   `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
   `course_name` VARCHAR(225) NULL DEFAULT '' COMMENT '课程名字' COLLATE 'utf8mb4_bin',
   `teacher` VARCHAR(225) NULL DEFAULT '' COMMENT '老师' COLLATE 'utf8mb4_bin',
   `update_time` INT(11) UNSIGNED NULL DEFAULT '0',
   `add_time` INT(11) UNSIGNED NULL DEFAULT '0',
   PRIMARY KEY (`id`)
)
    COMMENT='课程表'
    COLLATE='utf8mb4_bin'
    ENGINE=InnoDB
;