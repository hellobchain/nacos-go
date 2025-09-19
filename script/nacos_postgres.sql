-- ==========================
-- 1. 服务实例表 instance
-- ==========================
CREATE TABLE IF NOT EXISTS instance (
    id          BIGSERIAL PRIMARY KEY,
    service_name VARCHAR(128) NOT NULL,
    group_name   VARCHAR(128) NOT NULL,
    cluster_name VARCHAR(128) DEFAULT 'DEFAULT',
    ip           VARCHAR(64)         NOT NULL,          
    port         INT          NOT NULL,
    weight       DOUBLE PRECISION DEFAULT 1,
    healthy      SMALLINT     DEFAULT 1,        -- pg 无 TINYINT
    ephemeral    SMALLINT     DEFAULT 1,
    metadata     TEXT,
    expire_time  BIGINT       DEFAULT 0,
    CONSTRAINT uk_instance UNIQUE (service_name, group_name, ip::TEXT, port)
);

-- ==========================
-- 2. 配置表 config
-- ==========================
CREATE TABLE IF NOT EXISTS config (
    id           BIGSERIAL PRIMARY KEY,
    data_id      VARCHAR(128) NOT NULL,
    group_id     VARCHAR(128) NOT NULL DEFAULT 'DEFAULT_GROUP',
    content      TEXT         NOT NULL,
    md5          CHAR(32)     NOT NULL,
    beta_ips     VARCHAR(1024),
    src_user     VARCHAR(128),
    src_ip       VARCHAR(64),                       
    app_name     VARCHAR(128),
    tenant_id    VARCHAR(128) DEFAULT '',
    type         VARCHAR(16)  DEFAULT 'yaml',
    gmt_create   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    gmt_modified TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_config UNIQUE (data_id, group_id, tenant_id)
);

-- 3. 模拟 MySQL 的 ON UPDATE CURRENT_TIMESTAMP
CREATE OR REPLACE FUNCTION set_gmt_modified()
RETURNS TRIGGER AS $$
BEGIN
    NEW.gmt_modified := CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_config_gmt_modified ON config;
CREATE TRIGGER trg_config_gmt_modified
BEFORE UPDATE ON config
FOR EACH ROW
EXECUTE FUNCTION set_gmt_modified();