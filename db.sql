
DROP TABLE IF EXISTS `tinyurl_maps`;

CREATE TABLE `tinyurl_maps` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `origin_url` varchar(1024) COLLATE utf8mb4_bin NOT NULL,
  `short_url` varchar(14) COLLATE utf8mb4_bin NOT NULL,
  `created_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_idx_short_url` (`short_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
