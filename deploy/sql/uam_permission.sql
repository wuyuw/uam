USE `uam`;

/* 权限条目 */
DROP TABLE IF EXISTS `uam_permission`;
CREATE TABLE `uam_permission`
(
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL DEFAULT 0 COMMENT '所属客户端ID',
  `type` varchar(32) NOT NULL DEFAULT '' COMMENT '权限类型',
  `key` varchar(64) NOT NULL COMMENT '唯一标识',
  `name` varchar(64) NOT NULL COMMENT '权限名称',
  `desc` varchar(256) NOT NULL DEFAULT '' COMMENT '权限描述',
  `editable` tinyint NOT NULL DEFAULT 1 COMMENT '是否允许通过后台编辑',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_client_key` (`client_id`, `key`)
) ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='权限条目';