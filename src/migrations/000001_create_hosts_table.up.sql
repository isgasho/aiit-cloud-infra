CREATE TABLE IF NOT EXISTS `hosts` (
  `id` SERIAL PRIMARY KEY,             -- ホストID
  `name` VARCHAR(128) NOT NULL,        -- ホスト名
  `limit` INTEGER NOT NULL DEFAULT 0,  -- 払出容量
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
