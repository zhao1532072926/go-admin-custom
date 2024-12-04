# first-user && first-dd-detail CreateTableSQL

```sql
-- 创建数据库
CREATE DATABASE dev
DEFAULT CHARACTER SET utf8
DEFAULT COLLATE utf8_general_ci;

-- 创建数据表
CREATE TABLE first_users (
  id int unsigned NOT NULL AUTO_INCREMENT,
  phone varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  password varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE first_dd_detail (
  id int unsigned NOT NULL AUTO_INCREMENT,
  sender varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  spu_id varchar(20) DEFAULT NULL,
  spu_name text,
  shop_id varchar(15) DEFAULT NULL,
  shop_name text,
  send_time timestamp NULL DEFAULT NULL,
  data_body text,
  remove int DEFAULT NULL,
  youhui text,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```