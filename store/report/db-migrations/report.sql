SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for report
-- ----------------------------
CREATE TABLE `report` (
  `report_id` VARCHAR(50) NOT NULL COMMENT '报告ID',
  `developer_id` varchar(20) NOT NULL COMMENT '测量开发者ID',
  `app_id` varchar(20) NOT NULL COMMENT '受试者使用的appid',
  `subject_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '主体 ID',
  `rev` INT NOT NULL DEFAULT 0 COMMENT '数据版本号字段',
  `created_at` TIMESTAMP NOT NULL COMMENT '数据记录创建时间',
  `updated_at` TIMESTAMP NOT NULL COMMENT '数据记录更新时间',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '数据记录伪删除时间',
  PRIMARY KEY (`report_id`),
  KEY `idx_subject_id` (`subject_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报告内容';
