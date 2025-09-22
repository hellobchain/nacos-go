-- ==========================
-- 1. 服务实例表 instance
-- ==========================
CREATE TABLE IF NOT EXISTS instance (
    id           BIGSERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ,
    updated_at   TIMESTAMPTZ,
    deleted_at   TIMESTAMPTZ,
    service_name VARCHAR(128) NOT NULL,
    group_name   VARCHAR(128) NOT NULL,
    cluster_name VARCHAR(128) DEFAULT 'DEFAULT',
    ip           VARCHAR(64)  NOT NULL,
    port         BIGINT       NOT NULL,
    weight       DOUBLE PRECISION DEFAULT 1,
    healthy      SMALLINT     DEFAULT 1,        -- pg 无 TINYINT(1)
    ephemeral    SMALLINT     DEFAULT 1,
    metadata     TEXT,
    expire_time  BIGINT       DEFAULT 0,
    CONSTRAINT uk_instance UNIQUE (service_name, group_name, ip, port)
);
CREATE INDEX IF NOT EXISTS idx_instance_deleted_at ON instance (deleted_at);

-- ==========================
-- 2. 配置表 config
-- ==========================
CREATE TABLE IF NOT EXISTS config (
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    deleted_at  TIMESTAMPTZ,
    data_id     VARCHAR(128) NOT NULL,
    group_id    VARCHAR(128) NOT NULL DEFAULT 'DEFAULT_GROUP',
    tenant_id   VARCHAR(128) DEFAULT '',
    content     TEXT         NOT NULL,
    md5         VARCHAR(32)  NOT NULL,
    beta_ips    VARCHAR(1024),
    src_user    VARCHAR(128),
    src_ip      VARCHAR(64),
    app_name    VARCHAR(128),
    type        VARCHAR(16)  DEFAULT 'yaml',
    CONSTRAINT uk_config UNIQUE (data_id, group_id, tenant_id)
);
CREATE INDEX IF NOT EXISTS idx_config_deleted_at ON config (deleted_at);

-- 3. 自动更新 updated_at（等价 MySQL ON UPDATE）
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at := CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_instance_updated ON instance;
CREATE TRIGGER trg_instance_updated
    BEFORE UPDATE ON instance
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

DROP TRIGGER IF EXISTS trg_config_updated ON config;
CREATE TRIGGER trg_config_updated
    BEFORE UPDATE ON config
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();