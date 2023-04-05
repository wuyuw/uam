USE `uam`;

/* 角色-权限 */
DROP TABLE IF EXISTS `rel_role_permission`;
CREATE TABLE `rel_role_permission` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL DEFAULT 0 COMMENT '角色ID',
  `permission_id` int NOT NULL DEFAULT 0 COMMENT '权限ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_role_permission` (`role_id`, `permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色-权限关系';


/* 组-角色 */
DROP TABLE IF EXISTS `rel_group_role`;
CREATE TABLE `rel_group_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `group_id` int NOT NULL DEFAULT 0 COMMENT '组ID',
  `role_id` int NOT NULL DEFAULT 0 COMMENT '角色ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_group_role` (`group_id`, `role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='组-角色关系';


/* 用户-组 */
DROP TABLE IF EXISTS `rel_user_group`;
CREATE TABLE `rel_user_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL DEFAULT 0 COMMENT '客户端ID',
  `uid` int NOT NULL DEFAULT 0 COMMENT '员工UID',
  `group_id` int NOT NULL DEFAULT 0 COMMENT '组ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_user_group` (`uid`, `group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户-组关系';


/* 用户-角色 */
DROP TABLE IF EXISTS `rel_user_role`;
CREATE TABLE `rel_user_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL DEFAULT 0 COMMENT '客户端ID',
  `uid` int NOT NULL DEFAULT 0 COMMENT '员工UID',
  `role_id` int NOT NULL DEFAULT 0 COMMENT '角色ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_user_role` (`uid`, `role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户-角色关系';
