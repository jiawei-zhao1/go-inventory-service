CREATE TABLE `products` (
  `ID` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(250) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` int unsigned NOT NULL,
  `updated_at` int unsigned NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `products_name_IDX` (`name`) USING BTREE,
  KEY `products_created_at_IDX` (`created_at`) USING BTREE,
  KEY `products_updated_at_IDX` (`updated_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `inventory` (
  `product_id` int unsigned NOT NULL,
  `available_stock` int unsigned NOT NULL DEFAULT '0',
  `reserved_stock` int unsigned NOT NULL DEFAULT '0',
  `created_at` int unsigned NOT NULL DEFAULT '0',
  `updated_at` int unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`product_id`),
  KEY `inventory_created_at_IDX` (`created_at`) USING BTREE,
  KEY `inventory_updated_at_IDX` (`updated_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `inventory_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int unsigned NOT NULL DEFAULT '0',
  `quantity` int NOT NULL DEFAULT '0',
  `source` tinyint NOT NULL DEFAULT '0' COMMENT '1 order;',
  `source_id` bigint unsigned NOT NULL DEFAULT '0',
  `created_at` int unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `inventory_logs_source_id_IDX` (`source_id`,`source`,`product_id`) USING BTREE,
  KEY `inventory_logs_product_id_IDX` (`product_id`) USING BTREE,
  KEY `inventory_logs_created_at_IDX` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;