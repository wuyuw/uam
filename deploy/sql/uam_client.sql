USE `uam`;


/* 接入客户端 */
DROP TABLE IF EXISTS `uam_client`;
CREATE TABLE `uam_client` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '客户端名称',
  `app_code` varchar(16) NOT NULL DEFAULT '' COMMENT '客户端代码',
  `private_key` varchar(64) NOT NULL DEFAULT '' COMMENT '秘钥',
  `department` varchar(64) NOT NULL DEFAULT '' COMMENT '所属部门',
  `maintainer` varchar(64) NOT NULL DEFAULT '' COMMENT '对接人',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态: 0-正常, 1-禁用, 2-删除',
  `type` tinyint NOT NULL DEFAULT 1 COMMENT '客户端类型: 1-普通客户端, 2-系统后台',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_app_code` (`app_code`)
)ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='接入客户端';
