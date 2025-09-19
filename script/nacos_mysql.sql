-- ==========================
-- 1. 服务实例表 instance
-- ==========================
CREATE TABLE IF NOT EXISTS instance (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    service_name VARCHAR(128) NOT NULL,
    group_name   VARCHAR(128) NOT NULL,
    cluster_name VARCHAR(128) DEFAULT 'DEFAULT',
    ip           VARCHAR(64)  NOT NULL,
    port         INT          NOT NULL,
    weight       DOUBLE       DEFAULT 1,
    healthy      TINYINT(1)   DEFAULT 1,
    ephemeral    TINYINT(1)   DEFAULT 1,
    metadata     TEXT,
    expire_time  BIGINT       DEFAULT 0,
    UNIQUE KEY uk_instance (service_name, group_name, ip, port)
);

-- ==========================
-- 2. 配置表 config
-- ==========================
CREATE TABLE IF NOT EXISTS config (
    id            BIGINT AUTO_INCREMENT PRIMARY KEY,
    data_id       VARCHAR(128) NOT NULL,
    group_id      VARCHAR(128) NOT NULL DEFAULT 'DEFAULT_GROUP',
    content       TEXT         NOT NULL,
    md5           CHAR(32)     NOT NULL,
    beta_ips      VARCHAR(1024),
    src_user      VARCHAR(128),
    src_ip        VARCHAR(64),
    app_name      VARCHAR(128),
    tenant_id     VARCHAR(128) DEFAULT '',
    `type`        VARCHAR(16)  DEFAULT 'yaml',
    gmt_create    DATETIME DEFAULT CURRENT_TIMESTAMP,
    gmt_modified  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_config (data_id, group_id, tenant_id)
);