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