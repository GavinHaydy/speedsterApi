CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sys_user
(
    id            UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    username      VARCHAR(50) UNIQUE NOT NULL,
    password      VARCHAR(255)       NOT NULL,
    nickname      VARCHAR(50),
    avatar        TEXT,
    email         VARCHAR(100) UNIQUE,
    phone         VARCHAR(20) UNIQUE,

    status        SMALLINT           NOT NULL DEFAULT 1,
    is_super      BOOLEAN                     DEFAULT FALSE,
    is_sys_user   BOOLEAN                     DEFAULT FALSE,

    last_login_at TIMESTAMP,
    last_login_ip VARCHAR(45),

    created_at    TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP
);

INSERT INTO sys_user (id,
                      username,
                      password,
                      nickname,
                      is_super,
                      is_sys_user)
VALUES ('b62e9977-1bba-438e-983d-5e5e8344eced',
        'admin',
        'RPI8A-9CIUKMZ0-pxnkVNRJGaX6yS-2FqZEQ4sfCP-bIsFcceoFZ4aXz9JwxPEP5', -- bcrypt加密
        '超级管理员',
        TRUE,
        TRUE);


CREATE TABLE role
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(50)        NOT NULL,
    code        VARCHAR(50) UNIQUE NOT NULL, -- 如：admin/operator

    description TEXT,

    status      SMALLINT  DEFAULT 1,

    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    delete_at   TIMESTAMP
);

INSERT INTO role(name,code)
VALUES ('超级管理员','admin');

CREATE TABLE sys_user_role
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    UUID NOT NULL,
    role_id    BIGINT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (user_id, role_id)
);

INSERT INTO sys_user_role(user_id, role_id)
VALUES ('b62e9977-1bba-438e-983d-5e5e8344eced',1);

CREATE TABLE sys_permission (
                                id            BIGSERIAL PRIMARY KEY,
                                parent_id     BIGINT DEFAULT 0,

                                name          VARCHAR(100) NOT NULL,
                                code          VARCHAR(100) NOT NULL UNIQUE,

                                path          VARCHAR(255),
                                method        VARCHAR(20),

                                type          SMALLINT NOT NULL, -- 1菜单 2按钮 3接口
                                icon          VARCHAR(100),
                                sort          INT DEFAULT 0,

                                status        SMALLINT DEFAULT 1,
                                created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO sys_permission(name,code,path,method,type)
VALUES
    ('用户列表','user:list','/user/userlist','POST',3),
    ('退出登录','user:logout','/user/logout','DELETE',3);

CREATE TABLE sys_role_permission (
                                     id             BIGSERIAL PRIMARY KEY,
                                     role_id        BIGINT NOT NULL,
                                     permission_id  BIGINT NOT NULL,
                                     UNIQUE(role_id, permission_id)
);

INSERT INTO sys_role_permission(role_id, permission_id)
VALUES (1,1),(1,2);

CREATE INDEX idx_user_role_user_id ON sys_user_role(user_id);
CREATE INDEX idx_user_role_role_id ON sys_user_role(role_id);

CREATE INDEX idx_role_perm_role_id ON sys_role_permission(role_id);
CREATE INDEX idx_role_perm_perm_id ON sys_role_permission(permission_id);


-- ==========================================
-- 商家表
-- ==========================================

CREATE TABLE merchant (
                          id                  BIGSERIAL PRIMARY KEY,

                          merchant_code       VARCHAR(32) NOT NULL,

                          name                VARCHAR(100) NOT NULL,

                          logo                TEXT,

                          cover               TEXT,

                          description         TEXT,

                          notice              TEXT,

                          owner_id            BIGINT NOT NULL,

                          contact_name        VARCHAR(50) NOT NULL,

                          contact_phone       VARCHAR(20) NOT NULL,

                          email               VARCHAR(100),

                          license_no          VARCHAR(100) NOT NULL,

                          license_image       TEXT NOT NULL,

                          score               NUMERIC(2,1) NOT NULL DEFAULT 5.0,

                          avg_delivery_time   SMALLINT NOT NULL DEFAULT 30,

                          min_order_amount    NUMERIC(10,2) NOT NULL DEFAULT 0,

                          delivery_fee        NUMERIC(10,2) NOT NULL DEFAULT 0,

                          monthly_sales       INTEGER NOT NULL DEFAULT 0,

                          favorite_count      INTEGER NOT NULL DEFAULT 0,

                          view_count          INTEGER NOT NULL DEFAULT 0,

                          status              SMALLINT NOT NULL DEFAULT 1,

                          audit_status        SMALLINT NOT NULL DEFAULT 0,

                          created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                          updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                          deleted_at          TIMESTAMPTZ
);

-- ==========================================
-- Check Constraint
-- ==========================================

ALTER TABLE merchant
    ADD CONSTRAINT chk_merchant_status
        CHECK (status IN (1,2));

ALTER TABLE merchant
    ADD CONSTRAINT chk_merchant_audit
        CHECK (audit_status IN (0,1,2));

ALTER TABLE merchant
    ADD CONSTRAINT chk_merchant_score
        CHECK (score >= 0 AND score <= 5);

-- ==========================================
-- 索引
-- ==========================================

CREATE UNIQUE INDEX uk_merchant_code
    ON merchant(merchant_code);

CREATE UNIQUE INDEX uk_merchant_name
    ON merchant(name)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_merchant_owner
    ON merchant(owner_id);

CREATE INDEX idx_merchant_status
    ON merchant(status);

CREATE INDEX idx_merchant_audit
    ON merchant(audit_status);

-- ==========================================
-- 自动更新时间
-- ==========================================

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trg_merchant_updated
    BEFORE UPDATE
    ON merchant
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- ==========================================
-- 表注释
-- ==========================================

COMMENT ON TABLE merchant IS '商家信息';

COMMENT ON COLUMN merchant.id IS '主键ID';

COMMENT ON COLUMN merchant.merchant_code IS '商家编号';

COMMENT ON COLUMN merchant.name IS '商家名称';

COMMENT ON COLUMN merchant.logo IS '商家Logo';

COMMENT ON COLUMN merchant.cover IS '店铺封面图';

COMMENT ON COLUMN merchant.description IS '商家简介';

COMMENT ON COLUMN merchant.notice IS '商家公告';

COMMENT ON COLUMN merchant.owner_id IS '负责人用户ID';

COMMENT ON COLUMN merchant.contact_name IS '联系人';

COMMENT ON COLUMN merchant.contact_phone IS '联系电话';

COMMENT ON COLUMN merchant.email IS '联系邮箱';

COMMENT ON COLUMN merchant.license_no IS '营业执照编号';

COMMENT ON COLUMN merchant.license_image IS '营业执照图片';

COMMENT ON COLUMN merchant.score IS '商家评分';

COMMENT ON COLUMN merchant.avg_delivery_time IS '平均配送时间(分钟)';

COMMENT ON COLUMN merchant.min_order_amount IS '起送金额';

COMMENT ON COLUMN merchant.delivery_fee IS '配送费';

COMMENT ON COLUMN merchant.monthly_sales IS '月销量';

COMMENT ON COLUMN merchant.favorite_count IS '收藏数量';

COMMENT ON COLUMN merchant.view_count IS '浏览次数';

COMMENT ON COLUMN merchant.status IS '商家状态(1:正常 2:停业)';

COMMENT ON COLUMN merchant.audit_status IS '审核状态(0:待审核 1:审核通过 2:审核拒绝)';

COMMENT ON COLUMN merchant.created_at IS '创建时间';

COMMENT ON COLUMN merchant.updated_at IS '更新时间';

COMMENT ON COLUMN merchant.deleted_at IS '逻辑删除时间';



-- ==========================================
-- 门店表
-- ==========================================

CREATE TABLE merchant_store (
                                id                  BIGSERIAL PRIMARY KEY,

                                merchant_id         BIGINT NOT NULL,

                                store_code          VARCHAR(32) NOT NULL,

                                name                VARCHAR(100) NOT NULL,

                                phone               VARCHAR(20),

                                province            VARCHAR(50) NOT NULL,

                                city                VARCHAR(50) NOT NULL,

                                district            VARCHAR(50) NOT NULL,

                                address             VARCHAR(255) NOT NULL,

                                longitude           NUMERIC(10,6) NOT NULL,

                                latitude            NUMERIC(10,6) NOT NULL,

                                delivery_radius     INTEGER NOT NULL DEFAULT 3000,

                                min_order_amount    NUMERIC(10,2) NOT NULL DEFAULT 0,

                                delivery_fee        NUMERIC(10,2) NOT NULL DEFAULT 0,

                                packaging_fee       NUMERIC(10,2) NOT NULL DEFAULT 0,

                                avg_delivery_time   SMALLINT NOT NULL DEFAULT 30,

                                status              SMALLINT NOT NULL DEFAULT 1,

                                created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                deleted_at          TIMESTAMPTZ,

                                CONSTRAINT fk_store_merchant
                                    FOREIGN KEY (merchant_id)
                                        REFERENCES merchant(id)
);

-- ==========================================
-- Check Constraint
-- ==========================================

ALTER TABLE merchant_store
    ADD CONSTRAINT chk_store_status
        CHECK (status IN (1,2));

-- ==========================================
-- 索引
-- ==========================================

CREATE UNIQUE INDEX uk_store_code
    ON merchant_store(store_code);

CREATE UNIQUE INDEX uk_store_name
    ON merchant_store(merchant_id, name)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_store_merchant
    ON merchant_store(merchant_id);

CREATE INDEX idx_store_status
    ON merchant_store(status);

CREATE INDEX idx_store_location
    ON merchant_store(longitude, latitude);

-- ==========================================
-- Trigger
-- ==========================================

CREATE TRIGGER trg_store_updated
    BEFORE UPDATE
    ON merchant_store
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- ==========================================
-- 注释
-- ==========================================

COMMENT ON TABLE merchant_store IS '门店信息';

COMMENT ON COLUMN merchant_store.id IS '主键ID';

COMMENT ON COLUMN merchant_store.merchant_id IS '商家ID';

COMMENT ON COLUMN merchant_store.store_code IS '门店编号';

COMMENT ON COLUMN merchant_store.name IS '门店名称';

COMMENT ON COLUMN merchant_store.phone IS '门店联系电话';

COMMENT ON COLUMN merchant_store.province IS '省';

COMMENT ON COLUMN merchant_store.city IS '市';

COMMENT ON COLUMN merchant_store.district IS '区';

COMMENT ON COLUMN merchant_store.address IS '详细地址';

COMMENT ON COLUMN merchant_store.longitude IS '经度';

COMMENT ON COLUMN merchant_store.latitude IS '纬度';

COMMENT ON COLUMN merchant_store.delivery_radius IS '配送半径(米)';

COMMENT ON COLUMN merchant_store.min_order_amount IS '起送金额';

COMMENT ON COLUMN merchant_store.delivery_fee IS '配送费';

COMMENT ON COLUMN merchant_store.packaging_fee IS '打包费';

COMMENT ON COLUMN merchant_store.avg_delivery_time IS '平均配送时间(分钟)';

COMMENT ON COLUMN merchant_store.status IS '门店状态(1:营业 2:停业)';

COMMENT ON COLUMN merchant_store.created_at IS '创建时间';

COMMENT ON COLUMN merchant_store.updated_at IS '更新时间';

COMMENT ON COLUMN merchant_store.deleted_at IS '逻辑删除时间';


-- ==========================================
-- 门店营业时间
-- ==========================================

CREATE TABLE merchant_business_time (
                                        id                  BIGSERIAL PRIMARY KEY,

                                        store_id            BIGINT NOT NULL,

                                        week_day            SMALLINT NOT NULL,

                                        start_time          TIME NOT NULL,

                                        end_time            TIME NOT NULL,

                                        status              SMALLINT NOT NULL DEFAULT 1,

                                        created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                        updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                        CONSTRAINT fk_business_time_store
                                            FOREIGN KEY (store_id)
                                                REFERENCES merchant_store(id)
                                                ON DELETE CASCADE
);

-- ==========================================
-- Check Constraint
-- ==========================================

ALTER TABLE merchant_business_time
    ADD CONSTRAINT chk_business_week
        CHECK (week_day BETWEEN 1 AND 7);

ALTER TABLE merchant_business_time
    ADD CONSTRAINT chk_business_status
        CHECK (status IN (1,2));

ALTER TABLE merchant_business_time
    ADD CONSTRAINT chk_business_time
        CHECK (start_time < end_time);

-- ==========================================
-- 索引
-- ==========================================

CREATE INDEX idx_business_store
    ON merchant_business_time(store_id);

CREATE INDEX idx_business_week
    ON merchant_business_time(week_day);

CREATE UNIQUE INDEX uk_business_time
    ON merchant_business_time(
                              store_id,
                              week_day,
                              start_time,
                              end_time
        );

-- ==========================================
-- Trigger
-- ==========================================

CREATE TRIGGER trg_business_time_updated
    BEFORE UPDATE
    ON merchant_business_time
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- ==========================================
-- 注释
-- ==========================================

COMMENT ON TABLE merchant_business_time IS '门店营业时间';

COMMENT ON COLUMN merchant_business_time.id IS '主键ID';

COMMENT ON COLUMN merchant_business_time.store_id IS '门店ID';

COMMENT ON COLUMN merchant_business_time.week_day IS '星期(1=周一...7=周日)';

COMMENT ON COLUMN merchant_business_time.start_time IS '开始营业时间';

COMMENT ON COLUMN merchant_business_time.end_time IS '结束营业时间';

COMMENT ON COLUMN merchant_business_time.status IS '状态(1:营业 2:休息)';

COMMENT ON COLUMN merchant_business_time.created_at IS '创建时间';

COMMENT ON COLUMN merchant_business_time.updated_at IS '更新时间';


-- ==========================================
-- 商家结算配置
-- ==========================================

CREATE TABLE merchant_settlement (
                                     id                      BIGSERIAL PRIMARY KEY,

                                     merchant_id             BIGINT NOT NULL,

                                     settlement_type         SMALLINT NOT NULL DEFAULT 1,

                                     settlement_cycle        SMALLINT NOT NULL DEFAULT 1,

                                     commission_rate         NUMERIC(5,2) NOT NULL DEFAULT 0.00,

                                     min_withdraw_amount     NUMERIC(10,2) NOT NULL DEFAULT 100.00,

                                     bank_name               VARCHAR(100),

                                     bank_branch             VARCHAR(100),

                                     account_name            VARCHAR(100),

                                     account_no              VARCHAR(64),

                                     alipay_account          VARCHAR(100),

                                     wechat_account          VARCHAR(100),

                                     status                  SMALLINT NOT NULL DEFAULT 1,

                                     remark                  TEXT,

                                     created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                     updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                     CONSTRAINT fk_settlement_merchant
                                         FOREIGN KEY (merchant_id)
                                             REFERENCES merchant(id)
                                             ON DELETE CASCADE
);

-- ==========================================
-- Check Constraint
-- ==========================================

ALTER TABLE merchant_settlement
    ADD CONSTRAINT chk_settlement_type
        CHECK (settlement_type IN (1,2,3));

ALTER TABLE merchant_settlement
    ADD CONSTRAINT chk_settlement_cycle
        CHECK (settlement_cycle IN (1,7,15,30));

ALTER TABLE merchant_settlement
    ADD CONSTRAINT chk_settlement_status
        CHECK (status IN (1,2));

ALTER TABLE merchant_settlement
    ADD CONSTRAINT chk_commission_rate
        CHECK (
            commission_rate >= 0
                AND
            commission_rate <= 100
            );

-- ==========================================
-- 索引
-- ==========================================

CREATE UNIQUE INDEX uk_settlement_merchant
    ON merchant_settlement(merchant_id);

CREATE INDEX idx_settlement_status
    ON merchant_settlement(status);

-- ==========================================
-- Trigger
-- ==========================================

CREATE TRIGGER trg_settlement_updated
    BEFORE UPDATE
    ON merchant_settlement
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- ==========================================
-- 注释
-- ==========================================

COMMENT ON TABLE merchant_settlement IS '商家结算配置';

COMMENT ON COLUMN merchant_settlement.id IS '主键ID';

COMMENT ON COLUMN merchant_settlement.merchant_id IS '商家ID';

COMMENT ON COLUMN merchant_settlement.settlement_type IS '结算方式(1:银行卡 2:支付宝 3:微信)';

COMMENT ON COLUMN merchant_settlement.settlement_cycle IS '结算周期(1:每日 7:每周 15:半月 30:每月)';

COMMENT ON COLUMN merchant_settlement.commission_rate IS '平台抽成比例(%)';

COMMENT ON COLUMN merchant_settlement.min_withdraw_amount IS '最低提现金额';

COMMENT ON COLUMN merchant_settlement.bank_name IS '开户银行';

COMMENT ON COLUMN merchant_settlement.bank_branch IS '开户支行';

COMMENT ON COLUMN merchant_settlement.account_name IS '开户人';

COMMENT ON COLUMN merchant_settlement.account_no IS '银行卡号';

COMMENT ON COLUMN merchant_settlement.alipay_account IS '支付宝账号';

COMMENT ON COLUMN merchant_settlement.wechat_account IS '微信收款账号';

COMMENT ON COLUMN merchant_settlement.status IS '状态(1:启用 2:停用)';

COMMENT ON COLUMN merchant_settlement.remark IS '备注';

COMMENT ON COLUMN merchant_settlement.created_at IS '创建时间';

COMMENT ON COLUMN merchant_settlement.updated_at IS '更新时间';



-- ==========================================
-- 商家结算配置
-- ==========================================

CREATE TABLE merchant_settlement (
                                     id                      BIGSERIAL PRIMARY KEY,

                                     merchant_id             BIGINT NOT NULL,

                                     settlement_type         SMALLINT NOT NULL DEFAULT 1,

                                     settlement_cycle        SMALLINT NOT NULL DEFAULT 1,

                                     commission_rate         NUMERIC(5,2) NOT NULL DEFAULT 0.00,

                                     min_withdraw_amount     NUMERIC(10,2) NOT NULL DEFAULT 100.00,

                                     bank_name               VARCHAR(100),

                                     bank_branch             VARCHAR(100),

                                     account_name            VARCHAR(100),

                                     account_no              VARCHAR(64),

                                     alipay_account          VARCHAR(100),

                                     wechat_account          VARCHAR(100),

                                     status                  SMALLINT NOT NULL DEFAULT 1,

                                     remark                  TEXT,

                                     created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                     updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                     CONSTRAINT fk_settlement_merchant
                                         FOREIGN KEY (merchant_id)
                                             REFERENCES merchant(id)
                                             ON DELETE CASCADE
);

-- ==========================================
-- Check Constraint
-- ==========================================

ALTER TABLE merchant_settlement
    ADD CONSTRAINT chk_settlement_type
        CHECK (settlement_type IN (1,2,3));

ALTER TABLE merchant_settlement
    ADD CONSTRAINT chk_settlement_cycle
        CHECK (settlement_cycle IN (1,7,15,30));

ALTER TABLE merchant_settlement
    ADD CONSTRAINT chk_settlement_status
        CHECK (status IN (1,2));

ALTER TABLE merchant_settlement
    ADD CONSTRAINT chk_commission_rate
        CHECK (
            commission_rate >= 0
                AND
            commission_rate <= 100
            );

-- ==========================================
-- 索引
-- ==========================================

CREATE UNIQUE INDEX uk_settlement_merchant
    ON merchant_settlement(merchant_id);

CREATE INDEX idx_settlement_status
    ON merchant_settlement(status);

-- ==========================================
-- Trigger
-- ==========================================

CREATE TRIGGER trg_settlement_updated
    BEFORE UPDATE
    ON merchant_settlement
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- ==========================================
-- 注释
-- ==========================================

COMMENT ON TABLE merchant_settlement IS '商家结算配置';

COMMENT ON COLUMN merchant_settlement.id IS '主键ID';

COMMENT ON COLUMN merchant_settlement.merchant_id IS '商家ID';

COMMENT ON COLUMN merchant_settlement.settlement_type IS '结算方式(1:银行卡 2:支付宝 3:微信)';

COMMENT ON COLUMN merchant_settlement.settlement_cycle IS '结算周期(1:每日 7:每周 15:半月 30:每月)';

COMMENT ON COLUMN merchant_settlement.commission_rate IS '平台抽成比例(%)';

COMMENT ON COLUMN merchant_settlement.min_withdraw_amount IS '最低提现金额';

COMMENT ON COLUMN merchant_settlement.bank_name IS '开户银行';

COMMENT ON COLUMN merchant_settlement.bank_branch IS '开户支行';

COMMENT ON COLUMN merchant_settlement.account_name IS '开户人';

COMMENT ON COLUMN merchant_settlement.account_no IS '银行卡号';

COMMENT ON COLUMN merchant_settlement.alipay_account IS '支付宝账号';

COMMENT ON COLUMN merchant_settlement.wechat_account IS '微信收款账号';

COMMENT ON COLUMN merchant_settlement.status IS '状态(1:启用 2:停用)';

COMMENT ON COLUMN merchant_settlement.remark IS '备注';

COMMENT ON COLUMN merchant_settlement.created_at IS '创建时间';

COMMENT ON COLUMN merchant_settlement.updated_at IS '更新时间';



-- ==========================================
-- 商家入驻申请表
-- ==========================================

CREATE TABLE merchant_apply (
                                id                      BIGSERIAL PRIMARY KEY,

                                apply_no                VARCHAR(32) NOT NULL,

                                user_id                 BIGINT NOT NULL,

                                merchant_name           VARCHAR(100) NOT NULL,

                                logo                    TEXT,

                                cover                   TEXT,

                                description             TEXT,

                                contact_name            VARCHAR(50) NOT NULL,

                                contact_phone           VARCHAR(20) NOT NULL,

                                email                   VARCHAR(100),

                                license_no              VARCHAR(100) NOT NULL,

                                license_image           TEXT NOT NULL,

                                province                VARCHAR(50) NOT NULL,

                                city                    VARCHAR(50) NOT NULL,

                                district                VARCHAR(50) NOT NULL,

                                address                 VARCHAR(255) NOT NULL,

                                longitude               NUMERIC(10,6),

                                latitude                NUMERIC(10,6),

                                status                  SMALLINT NOT NULL DEFAULT 0,

                                reject_reason           TEXT,

                                auditor_id              BIGINT,

                                auditor_name            VARCHAR(100),

                                audit_time              TIMESTAMPTZ,

                                created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                deleted_at              TIMESTAMPTZ
);

-- ==========================================
-- Check Constraint
-- ==========================================

ALTER TABLE merchant_apply
    ADD CONSTRAINT chk_apply_status
        CHECK (status IN (0,1,2,3));

-- ==========================================
-- 索引
-- ==========================================

CREATE UNIQUE INDEX uk_apply_no
    ON merchant_apply(apply_no);

CREATE INDEX idx_apply_user
    ON merchant_apply(user_id);

CREATE INDEX idx_apply_status
    ON merchant_apply(status);

CREATE INDEX idx_apply_created
    ON merchant_apply(created_at DESC);

-- ==========================================
-- Trigger
-- ==========================================

CREATE TRIGGER trg_apply_updated
    BEFORE UPDATE
    ON merchant_apply
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- ==========================================
-- 注释
-- ==========================================

COMMENT ON TABLE merchant_apply IS '商家入驻申请';

COMMENT ON COLUMN merchant_apply.id IS '主键ID';

COMMENT ON COLUMN merchant_apply.apply_no IS '申请单号';

COMMENT ON COLUMN merchant_apply.user_id IS '申请人用户ID';

COMMENT ON COLUMN merchant_apply.merchant_name IS '商家名称';

COMMENT ON COLUMN merchant_apply.logo IS '商家Logo';

COMMENT ON COLUMN merchant_apply.cover IS '商家封面';

COMMENT ON COLUMN merchant_apply.description IS '商家简介';

COMMENT ON COLUMN merchant_apply.contact_name IS '联系人';

COMMENT ON COLUMN merchant_apply.contact_phone IS '联系电话';

COMMENT ON COLUMN merchant_apply.email IS '联系邮箱';

COMMENT ON COLUMN merchant_apply.license_no IS '营业执照编号';

COMMENT ON COLUMN merchant_apply.license_image IS '营业执照图片';

COMMENT ON COLUMN merchant_apply.province IS '省';

COMMENT ON COLUMN merchant_apply.city IS '市';

COMMENT ON COLUMN merchant_apply.district IS '区';

COMMENT ON COLUMN merchant_apply.address IS '详细地址';

COMMENT ON COLUMN merchant_apply.longitude IS '经度';

COMMENT ON COLUMN merchant_apply.latitude IS '纬度';

COMMENT ON COLUMN merchant_apply.status IS '申请状态(0:待审核 1:审核通过 2:审核拒绝 3:已取消)';

COMMENT ON COLUMN merchant_apply.reject_reason IS '审核拒绝原因';

COMMENT ON COLUMN merchant_apply.auditor_id IS '审核人ID';

COMMENT ON COLUMN merchant_apply.auditor_name IS '审核人';

COMMENT ON COLUMN merchant_apply.audit_time IS '审核时间';

COMMENT ON COLUMN merchant_apply.created_at IS '申请时间';

COMMENT ON COLUMN merchant_apply.updated_at IS '更新时间';

COMMENT ON COLUMN merchant_apply.deleted_at IS '逻辑删除时间';