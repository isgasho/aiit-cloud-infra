CREATE TABLE IF NOT EXISTS "instances" (
  "id" SERIAL PRIMARY KEY,            -- インスタンスID
  "host_id" INTEGER NOT NULL,         -- ホストID
  "name" VARCHAR(128) NOT NULL,       -- インスタンス名
  "state" SMALLINT NOT NULL,          -- ステータス
  "size" INTEGER NOT NULL DEFAULT 0,  -- 容量
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
