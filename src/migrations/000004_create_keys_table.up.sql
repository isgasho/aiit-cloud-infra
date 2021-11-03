CREATE TABLE IF NOT EXISTS "keys" (
  "id" SERIAL PRIMARY KEY,         -- 鍵ID
  "instance_id" INTEGER NOT NULL,  -- インスタンスID
  "data" TEXT NOT NULL,            -- 鍵情報
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
