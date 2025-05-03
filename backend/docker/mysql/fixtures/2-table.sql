-- users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime     NOT NULL,
  `updated_at` datetime     NOT NULL,
  `username`   varchar(64)  NOT NULL     COMMENT 'username',
  `password`   varchar(128) NOT NULL     COMMENT 'password',
  `nickname`   varchar(256) DEFAULT NULL COMMENT 'nickname',
  `avatar`     varchar(256) DEFAULT NULL COMMENT 'avatar',
  `email`      varchar(128) DEFAULT NULL COMMENT 'email',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- roles
CREATE TABLE IF NOT EXISTS `roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `name` varchar(64) NOT NULL COMMENT 'role_name',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- menus
CREATE TABLE IF NOT EXISTS `menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `parent_id` int(10) unsigned DEFAULT 0    COMMENT 'parent_menu_id',
  `name`      varchar(32)      NOT NULL     COMMENT 'manu_name',
  `path`      varchar(128)     DEFAULT NULL COMMENT 'path',
  `component` varchar(128)     DEFAULT NULL COMMENT 'component',
  `sort`      int(10) unsigned DEFAULT 0    COMMENT 'sort',
  `icon`      varchar(256)     DEFAULT NULL COMMENT 'icon',
  `title`     varchar(32)      DEFAULT NULL COMMENT 'title',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- user_role_bindings
CREATE TABLE IF NOT EXISTS `user_role_bindings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `user_id` int(10) unsigned NOT NULL COMMENT 'user_id',
  `role_id` int(10) unsigned NOT NULL COMMENT 'role_id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_role` (`user_id`, `role_id`),
  KEY `idx_role_id` (`role_id`),
  CONSTRAINT `fk_urb_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_urb_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- role_menu_bindings
CREATE TABLE IF NOT EXISTS `role_menu_bindings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `role_id` int(10) unsigned NOT NULL COMMENT 'role_id',
  `menu_id` int(10) unsigned NOT NULL COMMENT 'menu_id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_role_menu` (`role_id`, `menu_id`),
  KEY `idx_menu_id` (`menu_id`),
  CONSTRAINT `fk_rmb_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_rmb_menu_id` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;