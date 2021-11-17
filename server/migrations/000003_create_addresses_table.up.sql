CREATE TABLE IF NOT EXISTS "addresses" (
  "id" SERIAL PRIMARY KEY,              -- IPアドレスID
  "ip_address" VARCHAR(128) NOT NULL,   -- IPアドレス (cidr, inet は利用しない)
  "mac_address" VARCHAR(128) NOT NULL,  -- MACアドレス (macaddr は利用しない)
  "instance_id" INTEGER DEFAULT NULL,   -- インスタンスID
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
