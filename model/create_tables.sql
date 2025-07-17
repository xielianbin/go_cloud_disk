
-- 用户表
CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    open_id VARCHAR(255) NOT NULL UNIQUE,
    file_store_id INT NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    register_time DATETIME NOT NULL,
    image_path VARCHAR(255),
    INDEX idx_open_id (open_id)
);
-- 文件夹表
CREATE TABLE file_folder (
    id INT PRIMARY KEY AUTO_INCREMENT,
    file_folder_name VARCHAR(255) NOT NULL,
    parent_folder_id INT NOT NULL DEFAULT 0,
    file_store_id INT NOT NULL,
    time VARCHAR(50) NOT NULL,
    FOREIGN KEY (file_store_id) REFERENCES file_store(id),
    INDEX idx_file_store_id (file_store_id),
    INDEX idx_parent_folder_id (parent_folder_id)
);
-- 共享记录表
CREATE TABLE share (
    id INT PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(64) NOT NULL COMMENT '共享码',
    file_id INT NOT NULL COMMENT '文件ID',
    username VARCHAR(50) NOT NULL COMMENT '用户名',
    hash VARCHAR(128) NOT NULL COMMENT '文件哈希值',
    INDEX idx_code (code),
    INDEX idx_username (username),
    INDEX idx_file_id (file_id)
);
-- 文件仓库表
CREATE TABLE file_store (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    current_size BIGINT NOT NULL DEFAULT 0,
    max_size BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

-- 文件表
CREATE TABLE my_file (
    id INT PRIMARY KEY AUTO_INCREMENT,
    file_name VARCHAR(255) NOT NULL,
    file_hash VARCHAR(255) NOT NULL,
    file_store_id INT NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    download_num INT NOT NULL DEFAULT 0,
    upload_time VARCHAR(50) NOT NULL,
    parent_folder_id INT NOT NULL DEFAULT 0,
    size BIGINT NOT NULL,
    size_str VARCHAR(20) NOT NULL,
    type INT NOT NULL,
    postfix VARCHAR(20) NOT NULL,
    FOREIGN KEY (file_store_id) REFERENCES file_store(id),
    INDEX idx_file_store_id (file_store_id),
    INDEX idx_parent_folder_id (parent_folder_id)
);
