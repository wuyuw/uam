USE `uam`;


/* 用户 */
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          int         NOT NULL AUTO_INCREMENT,
    `uid`         int         NOT NULL DEFAULT 0 COMMENT 'uid',
    `nickname`    varchar(32) NOT NULL DEFAULT '' COMMENT '昵称',
    `email`       varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone`       varchar(32) NOT NULL DEFAULT '' COMMENT '手机',
    `create_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uidx_uid` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='用户';


/* 本地认证 */
DROP TABLE IF EXISTS `auth_local`;
CREATE TABLE `auth_local`
(
    `id`          int         NOT NULL AUTO_INCREMENT,
    `uid`         int         NOT NULL DEFAULT 0 COMMENT 'uid',
    `username`    varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
    `password`    varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
    `salt`        varchar(4)  NOT NULL DEFAULT '' COMMENT '盐',
    `create_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uidx_username` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='本地认证';

/* 三方认证 */
DROP TABLE IF EXISTS `auth_oauth`;
CREATE TABLE `auth_oauth`
(
    `id`          int         NOT NULL AUTO_INCREMENT,
    `uid`         int         NOT NULL DEFAULT 0 COMMENT 'uid',
    `oauth_type`  varchar(32) NOT NULL DEFAULT '' COMMENT 'OAuth类型',
    `oauth_id`    varchar(64) NOT NULL DEFAULT '' COMMENT 'OAuth ID',
    `oauth_token` varchar(42) NOT NULL DEFAULT '' COMMENT 'OAuth Token',
    `create_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uidx_type_id` (`oauth_type`, `oauth_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='OAuth认证';