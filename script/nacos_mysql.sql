-- ==========================
-- 1. 服务实例表 instance
-- ==========================
CREATE TABLE IF NOT EXISTS `instance` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `service_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `group_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cluster_name` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT 'DEFAULT',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `port` bigint(20) unsigned NOT NULL,
  `weight` double DEFAULT '1',
  `healthy` tinyint(1) DEFAULT '1',
  `ephemeral` tinyint(1) DEFAULT '1',
  `metadata` text COLLATE utf8mb4_unicode_ci,
  `expire_time` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_instance` (`service_name`,`group_name`,`ip`,`port`),
  KEY `idx_instance_deleted_at` (`deleted_at`)
) ;

-- ==========================
-- 2. 配置表 config
-- ==========================

CREATE TABLE IF NOT EXISTS `config` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `data_id` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `group_id` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'DEFAULT_GROUP',
  `tenant_id` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `md5` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `beta_ips` varchar(1024) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `src_user` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `src_ip` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `app_name` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT 'yaml',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_config` (`data_id`,`group_id`,`tenant_id`),
  KEY `idx_config_deleted_at` (`deleted_at`)
);