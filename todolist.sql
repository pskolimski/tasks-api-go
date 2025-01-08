/*
 Navicat Premium Dump SQL

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80040 (8.0.40-0ubuntu0.24.10.1)
 Source Host           : localhost:3306
 Source Schema         : todolist

 Target Server Type    : MySQL
 Target Server Version : 80040 (8.0.40-0ubuntu0.24.10.1)
 File Encoding         : 65001

 Date: 08/01/2025 22:06:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
  `uuid` varchar(255) COLLATE utf8mb3_polish_ci NOT NULL,
  `title` varchar(255) COLLATE utf8mb3_polish_ci DEFAULT NULL,
  `isDone` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_polish_ci;

SET FOREIGN_KEY_CHECKS = 1;
