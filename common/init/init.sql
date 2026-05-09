CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sys_user
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username      VARCHAR(50) UNIQUE NOT NULL,
    password      VARCHAR(255)       NOT NULL,
    nickname      VARCHAR(50),
    avatar        TEXT,
    email         VARCHAR(100),
    phone         VARCHAR(20),

    status        SMALLINT         DEFAULT 1,
    is_super      BOOLEAN          DEFAULT FALSE,
    is_sys_user   BOOLEAN          DEFAULT FALSE,

    last_login_at TIMESTAMP,
    last_login_ip VARCHAR(45),

    created_at    TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP
);

INSERT INTO sys_user (id,
                      username,
                      password,
                      nickname,
                      is_super,
                      is_sys_user)
VALUES (uuid_generate_v4(),
        'admin',
        '3BMHz_rEwHpi1xrAG8K9PQ', -- bcrypt加密
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
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sys_user_role
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    role_id    BIGINT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (user_id, role_id)
);