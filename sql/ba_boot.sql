/*
 Navicat Premium Data Transfer

 Source Server         : 本地Mysql
 Source Server Type    : MySQL
 Source Server Version : 80012 (8.0.12)
 Source Host           : localhost:3306
 Source Schema         : ba_boot

 Target Server Type    : MySQL
 Target Server Version : 80012 (8.0.12)
 File Encoding         : 65001

 Date: 25/04/2025 10:15:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for mon_logs_file
-- ----------------------------
DROP TABLE IF EXISTS `mon_logs_file`;
CREATE TABLE `mon_logs_file`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `file_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件物理路径',
  `file_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件网络路径',
  `file_type` tinyint(1) NOT NULL COMMENT '文件类型',
  `file_size` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件大小',
  `error_msg` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '异常日志',
  `status` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '1' COMMENT '上传状态(0:失败 1:成功)',
  `create_user_id` bigint(20) NOT NULL COMMENT '创建用户ID',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_user_id` bigint(20) NULL DEFAULT NULL COMMENT '修改用户ID',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '文件上传日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of mon_logs_file
-- ----------------------------

-- ----------------------------
-- Table structure for mon_logs_operation
-- ----------------------------
DROP TABLE IF EXISTS `mon_logs_operation`;
CREATE TABLE `mon_logs_operation`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'IP',
  `ip_addr` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'IP所属地',
  `user_id` bigint(20) NOT NULL COMMENT '此操作用户ID',
  `user_agent` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录代理',
  `request_uri` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求URI',
  `request_path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求PATH',
  `request_method` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求方式',
  `method_params` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求参数',
  `use_time` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求耗时',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of mon_logs_operation
-- ----------------------------
INSERT INTO `mon_logs_operation` VALUES (1, '::1', '', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.base$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '5020682100', '2025-04-24 09:39:42', '2025-04-24 09:39:42', 0);
INSERT INTO `mon_logs_operation` VALUES (2, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.base$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '115171900', '2025-04-24 09:40:29', '2025-04-24 09:40:29', 0);
INSERT INTO `mon_logs_operation` VALUES (3, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysRole/updateRoleMenu', '/sysRole/updateRoleMenu', 'PUT', '{\"menuIds\":[1,2,3,4,5,6,7,8,9,10],\"roleId\":1}', '100806100', '2025-04-24 09:41:01', '2025-04-24 09:41:01', 0);
INSERT INTO `mon_logs_operation` VALUES (4, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '119365500', '2025-04-24 10:23:04', '2025-04-24 10:23:04', 0);
INSERT INTO `mon_logs_operation` VALUES (5, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '115908700', '2025-04-24 10:23:17', '2025-04-24 10:23:17', 0);
INSERT INTO `mon_logs_operation` VALUES (6, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '84740900', '2025-04-24 10:23:27', '2025-04-24 10:23:27', 0);
INSERT INTO `mon_logs_operation` VALUES (7, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"2\",\"menuName\":\"操作日志\",\"routeName\":\"monitor_logs_operation\",\"routePath\":\"/monitor/logs/operation\",\"component\":\"view.monitor_logs_operation\",\"i18nKey\":\"route.monitor_logs_operation\",\"icon\":\"ic:round-list-alt\",\"iconType\":\"1\",\"parentId\":8,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[{\"code\":\"/monitor/monLogsOperation/page\",\"desc\":\"分页获取操作日志\"}],\"id\":10,\"children\":null}', '35161300', '2025-04-24 10:24:06', '2025-04-24 10:24:06', 0);
INSERT INTO `mon_logs_operation` VALUES (8, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysRole/updatePermissionIdsByRoleId', '/sysRole/updatePermissionIdsByRoleId', 'PUT', '{\"roleId\":1,\"permissionIds\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,608,436,262,496,657,394,351,335,42]}', '47839500', '2025-04-24 10:24:14', '2025-04-24 10:24:14', 0);
INSERT INTO `mon_logs_operation` VALUES (9, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '114306000', '2025-04-24 10:24:31', '2025-04-24 10:24:31', 0);
INSERT INTO `mon_logs_operation` VALUES (10, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '123980600', '2025-04-24 10:25:48', '2025-04-24 10:25:48', 0);
INSERT INTO `mon_logs_operation` VALUES (11, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '54440900', '2025-04-24 10:26:00', '2025-04-24 10:26:00', 0);
INSERT INTO `mon_logs_operation` VALUES (12, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.base$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '58.6335ms', '2025-04-24 10:38:13', '2025-04-24 10:38:13', 0);
INSERT INTO `mon_logs_operation` VALUES (13, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.base$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '141.6482ms', '2025-04-24 10:39:34', '2025-04-24 10:39:34', 0);
INSERT INTO `mon_logs_operation` VALUES (14, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.base$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '99.8312ms', '2025-04-24 10:40:58', '2025-04-24 10:40:58', 0);
INSERT INTO `mon_logs_operation` VALUES (15, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.base$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '72.6707ms', '2025-04-24 10:43:33', '2025-04-24 10:43:33', 0);
INSERT INTO `mon_logs_operation` VALUES (16, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.base$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '6.2131318s', '2025-04-24 10:45:13', '2025-04-24 10:45:13', 0);
INSERT INTO `mon_logs_operation` VALUES (17, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"日志管理\",\"routeName\":\"monitor_logs\",\"routePath\":\"/monitor/logs\",\"component\":\"layout.base\",\"i18nKey\":\"route.monitor_logs\",\"icon\":\"mdi:math-log\",\"iconType\":\"1\",\"parentId\":7,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":8,\"children\":[{\"id\":10,\"parentId\":8,\"menuType\":\"2\",\"menuName\":\"操作日志\",\"i18nKey\":\"route.monitor_logs_operation\",\"routeName\":\"monitor_logs_operation\",\"routePath\":\"/monitor/logs/operation\",\"icon\":\"ic:round-list-alt\",\"iconType\":\"1\",\"component\":\"view.monitor_logs_operation\",\"keepAlive\":false,\"hideInMenu\":false,\"constant\":false,\"href\":\"\",\"activeMenu\":\"\",\"order\":1,\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[{\"code\":\"/monitor/monLogsOperation/page\",\"desc\":\"分页获取操作日志\"}],\"status\":\"1\",\"children\":null},{\"id\":9,\"parentId\":8,\"menuType\":\"2\",\"menuName\":\"文件日志\",\"i18nKey\":\"route.monitor_logs_file\",\"routeName\":\"monitor_logs_file\",\"routePath\":\"/monitor/logs/file\",\"icon\":\"basil:document-outline\",\"iconType\":\"1\",\"component\":\"view.monitor_logs_file\",\"keepAlive\":false,\"hideInMenu\":false,\"constant\":false,\"href\":\"\",\"activeMenu\":\"\",\"order\":2,\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[{\"code\":\"/monitor/monLogsFile/page\",\"desc\":\"分页获取文件日志\"},{\"code\":\"/monitor/monLogsFile/remove\",\"desc\":\"删除日志并删除文件\"}],\"status\":\"1\",\"children\":null}]}', '122.6751ms', '2025-04-24 12:44:36', '2025-04-24 12:44:36', 0);
INSERT INTO `mon_logs_operation` VALUES (18, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.blank$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":false,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '79.798ms', '2025-04-24 14:53:34', '2025-04-24 14:53:34', 0);
INSERT INTO `mon_logs_operation` VALUES (19, '::1', '未知地址|IP异常', 0, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '144.1906ms', '2025-04-24 15:12:05', '2025-04-24 15:12:05', 0);
INSERT INTO `mon_logs_operation` VALUES (20, '::1', '未知地址|IP异常', 0, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '102.2565ms', '2025-04-24 15:15:48', '2025-04-24 15:15:48', 0);
INSERT INTO `mon_logs_operation` VALUES (21, '::1', '未知地址|IP异常', 0, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '107.19ms', '2025-04-24 15:19:01', '2025-04-24 15:19:01', 0);
INSERT INTO `mon_logs_operation` VALUES (22, '::1', '未知地址|IP异常', 1, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/sysMenu/update', '/sysMenu/update', 'PUT', '{\"menuType\":\"1\",\"menuName\":\"首页\",\"routeName\":\"home\",\"routePath\":\"/home\",\"component\":\"layout.blank$view.home\",\"i18nKey\":\"route.home\",\"icon\":\"icon-park-outline:data-sheet\",\"iconType\":\"1\",\"parentId\":0,\"status\":\"1\",\"keepAlive\":false,\"constant\":true,\"order\":1,\"href\":\"\",\"hideInMenu\":false,\"activeMenu\":\"\",\"multiTab\":false,\"fixedIndexInTab\":0,\"query\":[],\"buttons\":[],\"id\":5,\"children\":null,\"index\":1}', '91.6824ms', '2025-04-24 15:19:21', '2025-04-24 15:19:21', 0);
INSERT INTO `mon_logs_operation` VALUES (23, '::1', '未知地址|IP异常', 0, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '76.1743ms', '2025-04-24 15:23:28', '2025-04-24 15:23:28', 0);
INSERT INTO `mon_logs_operation` VALUES (24, '::1', '未知地址|IP异常', 0, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36', '/auth/login', '/auth/login', 'POST', '{\"userName\":\"admin\",\"password\":\"123456\"}', '129.4819ms', '2025-04-24 17:52:21', '2025-04-24 17:52:21', 0);

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典名称',
  `code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字典编码',
  `type` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '字典类型(1:系统字典,2:业务字典)',
  `sort` int(11) NULL DEFAULT 999 COMMENT '排序值',
  `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典描述',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否启用(0:禁用,1:启用)',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '数据字典管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
INSERT INTO `sys_dict` VALUES (1, '系统是否', 'yes_no', '1', 1, '通用性', '2025-04-15 10:40:42', '2025-04-15 07:01:51', '1', 0);
INSERT INTO `sys_dict` VALUES (2, '用户性别', 'gender', '1', 2, '用户性别', '2025-04-15 12:50:55', '2025-04-15 12:50:57', '1', 0);
INSERT INTO `sys_dict` VALUES (3, '启用状态', 'status', '1', 3, '系统通用的启用状态', '2025-04-15 12:55:02', '2025-04-15 12:55:04', '1', 0);
INSERT INTO `sys_dict` VALUES (4, '字典类型', 'dict_type', '1', 4, '字典类型', '2025-04-15 05:40:27', '2025-04-15 05:40:27', '1', 0);
INSERT INTO `sys_dict` VALUES (5, '菜单图标', 'menu_icon_type', '1', 4, '菜单图标', '2025-04-15 06:29:58', '2025-04-15 07:12:23', '1', 0);
INSERT INTO `sys_dict` VALUES (6, '系统状态', 'sys_status', '1', 5, '系统状态', '2025-04-22 08:06:04', '2025-04-22 08:06:04', '1', 0);

-- ----------------------------
-- Table structure for sys_dict_item
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_item`;
CREATE TABLE `sys_dict_item`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dict_id` bigint(20) NULL DEFAULT NULL COMMENT '父字典ID',
  `dict_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '父字典编码',
  `value` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '数据值',
  `zh_cn` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '中文名称',
  `en_us` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '英文名称',
  `type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类型(前端渲染类型)',
  `sort` int(11) NULL DEFAULT 999 COMMENT '排序值',
  `description` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典描述',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否启用(0:禁用,1:启用)',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '数据字典子项管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict_item
-- ----------------------------
INSERT INTO `sys_dict_item` VALUES (1, 1, 'yes_no', '1', '是', 'Yes', 'success', 1, '是', '2025-04-15 10:45:08', '2025-04-15 10:45:10', '1', 0);
INSERT INTO `sys_dict_item` VALUES (2, 1, 'yes_no', '2', '否', 'No', 'error', 2, '否', '2025-04-15 12:51:38', '2024-08-21 22:50:35', '1', 0);
INSERT INTO `sys_dict_item` VALUES (3, 2, 'gender', '1', '保密', 'Confidential', NULL, 1, '保密', '2025-04-15 12:52:37', '2024-11-04 16:53:33', '1', 0);
INSERT INTO `sys_dict_item` VALUES (4, 2, 'gender', '1', '男', 'Male', 'primary', 2, '男', '2025-04-15 12:53:05', '2025-04-15 12:53:07', '1', 0);
INSERT INTO `sys_dict_item` VALUES (5, 2, 'gender', '2', '女', 'Female', 'error', 3, '女', '2025-04-15 12:53:41', '2025-04-15 12:53:44', '1', 0);
INSERT INTO `sys_dict_item` VALUES (6, 3, 'status', '1', '启用', 'Enable', 'success', 0, '启用', '2025-04-15 12:55:43', '2025-04-15 12:55:46', '1', 0);
INSERT INTO `sys_dict_item` VALUES (7, 3, 'status', '0', '禁用', 'Disable', 'error', 1, '禁用', '2025-04-15 12:56:06', '2025-04-15 12:56:09', '1', 0);
INSERT INTO `sys_dict_item` VALUES (8, 4, 'dict_type', '1', '系统字典', 'System', 'default', 1, '系统字典', '2025-04-15 06:09:58', '2025-04-15 06:09:58', '1', 0);
INSERT INTO `sys_dict_item` VALUES (9, 4, 'dict_type', '2', '业务字典', 'Business', 'default', 2, '业务字典', '2025-04-15 06:12:15', '2025-04-15 06:12:15', '1', 0);
INSERT INTO `sys_dict_item` VALUES (10, 5, 'menu_icon_type', '1', 'iconify图标', 'Iconify Icon', 'default', 1, 'iconify图标', '2025-04-15 06:31:09', '2025-04-15 06:31:09', '1', 0);
INSERT INTO `sys_dict_item` VALUES (11, 5, 'menu_icon_type', '2', '本地图标', 'Local Icon', 'default', 1, '本地图标', '2025-04-15 06:32:16', '2025-04-15 06:32:16', '1', 0);
INSERT INTO `sys_dict_item` VALUES (12, 6, 'sys_status', '0', '失败', 'fail', 'error', 1, '', '2025-04-22 08:06:29', '2025-04-22 08:06:29', '1', 0);
INSERT INTO `sys_dict_item` VALUES (13, 6, 'sys_status', '1', '成功', 'success', 'success', 1, '', '2025-04-22 08:06:42', '2025-04-22 08:06:42', '1', 0);

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父菜单ID',
  `menu_type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单类型 1:目录 2:菜单',
  `menu_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `i18n_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '多语言标题',
  `route_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '路由名称',
  `route_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单路径',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `icon_type` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图标类型 1:iconify icon 2:local icon',
  `component` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由组件',
  `keep_alive` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '缓存页面(Y:是,N:否)',
  `hide_in_menu` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '是否隐藏(Y:是,N:否)',
  `constant` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否为常量路由(Y:是,N:否)',
  `href` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '外部链接',
  `order` int(11) NULL DEFAULT 999 COMMENT '排序值',
  `multi_tab` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '支持多标签(Y:是,N:否)',
  `fixed_index_in_tab` int(11) NULL DEFAULT NULL COMMENT '固定在页签中的序号',
  `active_menu` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '内链URL',
  `query` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '路由查询参数',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否启用(0:禁用,1:启用)',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, 0, '1', '系统管理', 'route.manage', 'manage', '/manage', 'solar:settings-linear', '1', 'layout.base', 'N', 'N', 'N', '', 2, 'N', NULL, NULL, '[]', '2023-12-28 22:22:50', '2025-04-22 07:49:15', '1', 0);
INSERT INTO `sys_menu` VALUES (2, 1, '2', '用户管理', 'route.manage_user', 'manage_user', '/manage/user', 'ic:round-manage-accounts', '1', 'view.manage_user', 'N', 'N', 'N', '', 1, 'N', NULL, NULL, '[]', '2023-12-28 22:22:50', '2025-04-22 07:49:35', '1', 0);
INSERT INTO `sys_menu` VALUES (3, 1, '2', '角色管理', 'route.manage_role', 'manage_role', '/manage/role', 'ic:round-people', '1', 'view.manage_role', 'N', 'N', 'N', '', 2, 'N', NULL, NULL, '[]', '2023-12-28 22:22:50', '2024-11-04 14:27:36', '1', 0);
INSERT INTO `sys_menu` VALUES (4, 1, '2', '菜单管理', 'route.manage_menu', 'manage_menu', '/manage/menu', 'ic:round-menu', '1', 'view.manage_menu', 'N', 'N', 'N', '', 3, 'N', NULL, NULL, '[]', '2023-12-28 22:22:50', '2024-11-04 14:30:22', '1', 0);
INSERT INTO `sys_menu` VALUES (5, 0, '1', '首页', 'route.home', 'home', '/home', 'icon-park-outline:data-sheet', '1', 'layout.blank$view.home', 'N', 'N', 'Y', '', 1, 'N', NULL, NULL, '[]', '2024-02-05 22:30:15', '2025-04-24 15:19:21', '1', 0);
INSERT INTO `sys_menu` VALUES (6, 1, '2', '字典管理', 'route.manage_dict', 'manage_dict', '/manage/dict', 'ic:round-menu-book', '1', 'view.manage_dict', 'N', 'N', 'N', '', 4, 'N', NULL, NULL, '[]', '2025-04-15 01:51:05', '2025-04-22 08:25:07', '1', 0);
INSERT INTO `sys_menu` VALUES (7, 0, '1', '监控管理', 'route.monitor', 'monitor', '/monitor', 'icon-park-solid:monitor-one', '1', 'layout.base', 'N', 'N', 'N', NULL, 3, 'N', NULL, NULL, '[]', '2025-04-21 17:20:50', '2025-04-21 20:23:16', '1', 0);
INSERT INTO `sys_menu` VALUES (8, 7, '1', '日志管理', 'route.monitor_logs', 'monitor_logs', '/monitor/logs', 'mdi:math-log', '1', 'layout.base', 'N', 'N', 'N', NULL, 1, 'N', NULL, NULL, '[]', '2025-04-21 17:22:08', '2025-04-24 12:44:36', '1', 0);
INSERT INTO `sys_menu` VALUES (9, 8, '2', '文件日志', 'route.monitor_logs_file', 'monitor_logs_file', '/monitor/logs/file', 'basil:document-outline', '1', 'view.monitor_logs_file', 'N', 'N', 'N', NULL, 2, 'N', NULL, NULL, '[]', '2025-04-21 17:23:41', '2025-04-24 09:24:57', '1', 0);
INSERT INTO `sys_menu` VALUES (10, 8, '2', '操作日志', 'route.monitor_logs_operation', 'monitor_logs_operation', '/monitor/logs/operation', 'ic:round-list-alt', '1', 'view.monitor_logs_operation', 'N', 'N', 'N', '', 1, 'N', 0, '', '[]', '2025-04-24 09:24:38', '2025-04-24 10:24:06', '1', 0);

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `code` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限资源',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否启用(0:禁用,1:启用)',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限(按钮)管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
INSERT INTO `sys_permission` VALUES (1, 2, '用户管理', '/sysUser/page', '获取用户列表分页', '2025-03-12 09:17:32', '2025-03-12 09:17:36', '1', 0);
INSERT INTO `sys_permission` VALUES (2, 1, '系统管理', '/auth/getUserInfo', '获取用户详细详细', '2025-03-12 15:18:50', '2025-03-12 15:18:53', '1', 0);
INSERT INTO `sys_permission` VALUES (3, 1, '系统管理', '/route/getUserRoutes', '获取用户路由列表', '2025-03-12 15:21:43', '2025-03-12 15:21:45', '1', 0);
INSERT INTO `sys_permission` VALUES (4, 2, '用户管理', '/sysUser/removeByIds', '删除多个用户', '2025-03-12 15:23:46', '2025-03-12 15:23:48', '1', 0);
INSERT INTO `sys_permission` VALUES (5, 3, '角色管理', '/sysRole/getAllRoles', '获取角色列表', '2025-03-12 15:25:19', '2025-03-12 15:25:22', '1', 0);
INSERT INTO `sys_permission` VALUES (6, 2, '用户管理', '/sysUser/save', '添加用户', '2025-03-12 15:27:03', '2025-03-12 15:27:05', '1', 0);
INSERT INTO `sys_permission` VALUES (7, 2, '用户管理', '/sysUser/remove', '删除单个用户', '2025-03-12 15:39:54', '2025-03-12 15:39:57', '1', 0);
INSERT INTO `sys_permission` VALUES (8, 2, '用户管理', '/sysUser/update', '修改用户信息', '2025-03-12 07:42:21', '2025-03-12 07:42:23', '1', 0);
INSERT INTO `sys_permission` VALUES (9, 3, '角色管理', '/sysRole/page', '获取角色列表分页', '2025-03-13 09:56:42', '2025-03-13 09:56:45', '1', 0);
INSERT INTO `sys_permission` VALUES (10, 3, '角色管理', '/sysRole/save', '新增角色', '2025-03-13 10:09:31', '2025-03-13 10:09:33', '1', 0);
INSERT INTO `sys_permission` VALUES (11, 3, '角色管理', '/sysRole/removeByIds', '删除多个角色', '2025-03-13 10:34:51', '2025-03-13 10:34:53', '1', 0);
INSERT INTO `sys_permission` VALUES (12, 3, '角色管理', '/sysRole/remove', '删除单个角色', '2025-03-13 10:35:17', '2025-03-13 10:35:20', '1', 0);
INSERT INTO `sys_permission` VALUES (13, 3, '角色管理', '/sysRole/update', '修改角色信息', '2025-03-13 10:57:39', '2025-03-13 10:57:41', '1', 0);
INSERT INTO `sys_permission` VALUES (14, 4, '菜单管理', '/sysMenu/getAllMenus', '获取菜单名称列表', '2025-03-13 11:33:27', '2025-03-13 11:33:30', '1', 0);
INSERT INTO `sys_permission` VALUES (15, 4, '菜单管理', '/sysMenu/getMenuTree', '获取菜单树', '2025-03-13 13:45:54', '2025-03-13 13:45:56', '1', 0);
INSERT INTO `sys_permission` VALUES (16, 4, '菜单管理', '/sysMenu/getMenuIdsByRoleId', '获取菜单ID列表', '2025-03-13 14:09:04', '2025-03-13 14:09:06', '1', 0);
INSERT INTO `sys_permission` VALUES (17, 3, '角色管理', '/sysRole/updateRoleMenu', '修改角色菜单权限', '2025-03-13 15:30:28', '2025-03-13 15:30:30', '1', 0);
INSERT INTO `sys_permission` VALUES (18, 3, '角色管理', '/sysRole/getPermissionTree', '获取按钮树', '2025-03-13 17:07:51', '2025-03-13 17:07:54', '1', 0);
INSERT INTO `sys_permission` VALUES (19, 3, '角色管理', '/sysRole/getPermissionIdsByRoleId', '获取角色按钮列表', '2025-03-15 12:18:52', '2025-03-15 12:18:55', '1', 0);
INSERT INTO `sys_permission` VALUES (20, 3, '角色管理', '/sysRole/updatePermissionIdsByRoleId', '修改角色按钮列表', '2025-03-15 12:40:11', '2025-03-15 12:40:12', '1', 0);
INSERT INTO `sys_permission` VALUES (21, 4, '菜单管理', '/sysMenu/page', '获取菜单列表分页', '2025-03-15 17:00:17', '2025-03-15 17:00:20', '1', 0);
INSERT INTO `sys_permission` VALUES (22, 4, '菜单管理', '/sysMenu/save', '添加菜单', '2025-03-15 18:45:51', '2025-03-15 18:45:53', '1', 0);
INSERT INTO `sys_permission` VALUES (23, 4, '菜单管理', '/sysMenu/removeByIds', '删除多个菜单', '2025-03-20 01:21:53', '2025-03-20 01:21:53', '1', 0);
INSERT INTO `sys_permission` VALUES (24, 4, '菜单管理', '/sysMenu/remove', '删除单个菜单', '2025-03-20 01:38:17', '2025-03-20 01:38:17', '1', 0);
INSERT INTO `sys_permission` VALUES (25, 4, '菜单管理', '/sysMenu/update', '修改菜单', '2025-03-20 10:54:34', '2025-03-20 10:54:36', '1', 0);
INSERT INTO `sys_permission` VALUES (26, 6, '字典管理', '/sysDict/allDict', '获取全部字典Options', '2025-04-15 01:51:05', '2025-04-15 01:51:05', '1', 0);
INSERT INTO `sys_permission` VALUES (27, 6, '字典管理', '/sysDict/dictAll', '获取全部字典Vo', '2025-04-15 02:32:51', '2025-04-15 02:32:51', '1', 0);
INSERT INTO `sys_permission` VALUES (28, 6, '字典管理', '/sysDict/itemPage', '分页获取子字典Vo', '2025-04-15 04:30:35', '2025-04-15 04:30:35', '1', 0);
INSERT INTO `sys_permission` VALUES (29, 6, '字典管理', '/sysDict/update', '修改字典', '2025-04-15 04:43:43', '2025-04-15 04:43:43', '1', 0);
INSERT INTO `sys_permission` VALUES (30, 6, '字典管理', '/sysDict/delete', '删除字段', '2025-04-15 04:43:43', '2025-04-15 04:43:43', '1', 0);
INSERT INTO `sys_permission` VALUES (31, 6, '字典管理', '/sysDict/save', '新增字段', '2025-04-15 04:43:43', '2025-04-15 04:43:43', '1', 0);
INSERT INTO `sys_permission` VALUES (32, 6, '字典管理', '/sysDict/getInfo', '获取字典详细', '2025-04-15 05:55:39', '2025-04-15 05:55:39', '1', 0);
INSERT INTO `sys_permission` VALUES (33, 6, '字典管理', '/sysDict/removeItemByIds', '删除多个子字典', '2025-04-15 06:21:46', '2025-04-15 06:21:46', '1', 0);
INSERT INTO `sys_permission` VALUES (34, 6, '字典管理', '/sysDict/removeItemById', '删除单个子字典', '2025-04-15 06:21:46', '2025-04-15 06:21:46', '1', 0);
INSERT INTO `sys_permission` VALUES (35, 6, '字典管理', '/sysDict/removeById', '删除单个字典', '2025-04-15 06:21:46', '2025-04-15 06:21:46', '1', 0);
INSERT INTO `sys_permission` VALUES (36, 6, '字典管理', '/sysDict/updateItem', '修改子字典', '2025-04-15 06:22:58', '2025-04-15 06:22:58', '1', 0);
INSERT INTO `sys_permission` VALUES (37, 7, '文件管理', '/sysFile/upload', '上传文件', '2025-04-18 07:39:21', '2025-04-18 07:39:21', '1', 0);
INSERT INTO `sys_permission` VALUES (38, 9, '文件日志', '/monitor/monLogsFile/page', '分页获取文件日志', '2025-04-22 07:26:40', '2025-04-22 07:26:40', '1', 0);
INSERT INTO `sys_permission` VALUES (39, 2, '用户管理', '/sysUser/allUserName', '获取所有用户枚举', '2025-04-22 07:49:35', '2025-04-22 07:49:35', '1', 0);
INSERT INTO `sys_permission` VALUES (40, 6, '字典管理', '/sysDict/getItemInfo', '获取子字典详细', '2025-04-22 08:25:07', '2025-04-22 08:25:07', '1', 0);
INSERT INTO `sys_permission` VALUES (41, 9, '文件日志', '/monitor/monLogsFile/remove', '删除日志并删除文件', '2025-04-22 09:15:21', '2025-04-22 09:15:21', '1', 0);
INSERT INTO `sys_permission` VALUES (42, 10, '操作日志', '/monitor/monLogsOperation/page', '分页获取操作日志', '2025-04-24 10:24:06', '2025-04-24 10:24:06', '1', 0);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
  `role_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色编码',
  `role_desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否启用(0:禁用,1:启用)',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', 'R_SUPER', '系统管理员', '2023-12-31 17:28:23', '2025-03-13 03:00:32', '1', 0);
INSERT INTO `sys_role` VALUES (2, '管理员', 'R_MANAGER', '管理员', '2025-03-13 02:13:10', '2025-03-13 02:13:10', '1', 0);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色菜单管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (1, 1, 1, '2025-02-18 07:21:03', '2025-02-18 07:21:03', 0);
INSERT INTO `sys_role_menu` VALUES (2, 1, 2, '2025-02-17 23:21:03', '2025-02-17 23:21:03', 0);
INSERT INTO `sys_role_menu` VALUES (3, 1, 3, '2025-02-17 23:21:03', '2025-02-17 23:21:03', 0);
INSERT INTO `sys_role_menu` VALUES (4, 1, 4, '2025-02-17 23:21:03', '2025-02-17 23:21:03', 0);
INSERT INTO `sys_role_menu` VALUES (5, 1, 5, '2025-02-17 23:21:03', '2025-02-17 23:21:03', 0);
INSERT INTO `sys_role_menu` VALUES (6, 2, 1, '2025-03-14 13:42:03', '2025-03-14 13:42:05', 0);
INSERT INTO `sys_role_menu` VALUES (7, 2, 2, '2025-03-15 05:48:09', '2025-03-15 05:48:09', 0);
INSERT INTO `sys_role_menu` VALUES (8, 2, 3, '2025-03-15 05:48:10', '2025-03-15 05:48:10', 0);
INSERT INTO `sys_role_menu` VALUES (9, 2, 4, '2025-03-15 05:48:10', '2025-03-15 05:48:10', 0);
INSERT INTO `sys_role_menu` VALUES (10, 2, 5, '2025-03-14 14:29:09', '2025-03-14 14:29:09', 0);
INSERT INTO `sys_role_menu` VALUES (11, 1, 6, '2025-04-15 01:51:21', '2025-04-15 01:51:21', 0);
INSERT INTO `sys_role_menu` VALUES (12, 1, 7, '2025-04-18 07:41:43', '2025-04-18 07:41:43', 0);
INSERT INTO `sys_role_menu` VALUES (13, 1, 8, '2025-04-21 09:14:42', '2025-04-21 09:14:42', 0);
INSERT INTO `sys_role_menu` VALUES (14, 1, 9, '2025-04-21 09:24:16', '2025-04-21 09:24:16', 0);
INSERT INTO `sys_role_menu` VALUES (15, 1, 10, '2025-04-24 09:41:00', '2025-04-24 09:41:00', 0);

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `permission_id` bigint(20) NOT NULL COMMENT '权限ID',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色权限管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_permission
-- ----------------------------
INSERT INTO `sys_role_permission` VALUES (1, 1, 1, '2025-03-12 16:23:14', '2025-03-12 16:23:16', 0);
INSERT INTO `sys_role_permission` VALUES (2, 1, 2, '2025-03-12 16:23:14', '2025-03-12 16:23:16', 0);
INSERT INTO `sys_role_permission` VALUES (3, 1, 3, '2025-03-12 16:25:45', '2025-03-12 16:25:47', 0);
INSERT INTO `sys_role_permission` VALUES (4, 1, 4, '2025-03-12 16:25:52', '2025-03-12 16:25:55', 0);
INSERT INTO `sys_role_permission` VALUES (5, 1, 5, '2025-03-12 16:26:00', '2025-03-12 16:26:03', 0);
INSERT INTO `sys_role_permission` VALUES (6, 1, 6, '2025-03-12 16:26:13', '2025-03-12 16:26:15', 0);
INSERT INTO `sys_role_permission` VALUES (7, 1, 7, '2025-03-12 16:26:22', '2025-03-12 16:26:24', 0);
INSERT INTO `sys_role_permission` VALUES (8, 1, 8, '2025-03-12 16:26:31', '2025-03-12 16:26:34', 0);
INSERT INTO `sys_role_permission` VALUES (9, 1, 9, '2025-03-13 09:58:47', '2025-03-13 09:58:50', 0);
INSERT INTO `sys_role_permission` VALUES (10, 1, 10, '2025-03-13 10:09:44', '2025-03-13 10:09:47', 0);
INSERT INTO `sys_role_permission` VALUES (11, 1, 11, '2025-03-13 10:35:37', '2025-03-13 10:35:39', 0);
INSERT INTO `sys_role_permission` VALUES (12, 1, 12, '2025-03-13 10:35:45', '2025-03-13 10:35:47', 0);
INSERT INTO `sys_role_permission` VALUES (13, 1, 13, '2025-03-13 10:57:49', '2025-03-13 10:57:52', 0);
INSERT INTO `sys_role_permission` VALUES (14, 1, 14, '2025-03-13 11:33:48', '2025-03-13 11:33:50', 0);
INSERT INTO `sys_role_permission` VALUES (15, 1, 15, '2025-03-13 13:46:06', '2025-03-13 13:46:08', 0);
INSERT INTO `sys_role_permission` VALUES (16, 1, 16, '2025-03-13 14:09:13', '2025-03-13 14:09:15', 0);
INSERT INTO `sys_role_permission` VALUES (17, 1, 17, '2025-03-13 15:30:38', '2025-03-13 15:30:40', 0);
INSERT INTO `sys_role_permission` VALUES (18, 1, 18, '2025-03-13 17:08:06', '2025-03-13 17:08:09', 0);
INSERT INTO `sys_role_permission` VALUES (19, 1, 19, '2025-03-15 12:19:07', '2025-03-15 12:19:09', 0);
INSERT INTO `sys_role_permission` VALUES (20, 1, 20, '2025-03-15 12:40:26', '2025-03-15 12:40:29', 0);
INSERT INTO `sys_role_permission` VALUES (21, 2, 1, '2025-03-15 06:38:44', '2025-03-15 06:38:44', 0);
INSERT INTO `sys_role_permission` VALUES (22, 2, 4, '2025-03-15 06:48:29', '2025-03-15 06:48:29', 0);
INSERT INTO `sys_role_permission` VALUES (23, 2, 6, '2025-03-15 06:48:29', '2025-03-15 06:48:29', 0);
INSERT INTO `sys_role_permission` VALUES (24, 2, 7, '2025-03-15 06:50:58', '2025-03-15 06:50:58', 0);
INSERT INTO `sys_role_permission` VALUES (25, 2, 8, '2025-03-15 06:50:58', '2025-03-15 06:50:58', 0);
INSERT INTO `sys_role_permission` VALUES (26, 2, 9, '2025-03-15 06:50:58', '2025-03-15 06:50:58', 0);
INSERT INTO `sys_role_permission` VALUES (27, 2, 5, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (28, 2, 10, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (29, 2, 11, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (30, 2, 12, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (31, 2, 13, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (32, 2, 17, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (33, 2, 18, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (34, 2, 19, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (35, 2, 20, '2025-03-15 07:07:08', '2025-03-15 07:07:08', 0);
INSERT INTO `sys_role_permission` VALUES (36, 1, 21, '2025-03-15 17:00:43', '2025-03-15 17:00:46', 0);
INSERT INTO `sys_role_permission` VALUES (37, 1, 22, '2025-03-15 10:55:30', '2025-03-15 10:55:30', 0);
INSERT INTO `sys_role_permission` VALUES (38, 1, 23, '2025-03-20 01:22:25', '2025-03-20 01:22:25', 0);
INSERT INTO `sys_role_permission` VALUES (39, 1, 24, '2025-03-20 01:38:55', '2025-03-20 01:38:55', 0);
INSERT INTO `sys_role_permission` VALUES (40, 1, 15, '2025-03-20 10:54:56', '2025-03-20 10:54:57', 0);
INSERT INTO `sys_role_permission` VALUES (41, 1, 25, '2025-03-20 02:57:02', '2025-03-20 02:57:02', 0);
INSERT INTO `sys_role_permission` VALUES (42, 1, 26, '2025-04-15 02:38:07', '2025-04-15 02:38:07', 0);
INSERT INTO `sys_role_permission` VALUES (43, 1, 27, '2025-04-15 02:38:07', '2025-04-15 02:38:07', 0);
INSERT INTO `sys_role_permission` VALUES (44, 1, 28, '2025-04-15 04:30:47', '2025-04-15 04:30:47', 0);
INSERT INTO `sys_role_permission` VALUES (45, 1, 29, '2025-04-15 04:44:02', '2025-04-15 04:44:02', 0);
INSERT INTO `sys_role_permission` VALUES (46, 1, 30, '2025-04-15 04:44:02', '2025-04-15 04:44:02', 0);
INSERT INTO `sys_role_permission` VALUES (47, 1, 31, '2025-04-15 04:44:02', '2025-04-15 04:44:02', 0);
INSERT INTO `sys_role_permission` VALUES (48, 1, 32, '2025-04-15 05:55:47', '2025-04-15 05:55:47', 0);
INSERT INTO `sys_role_permission` VALUES (49, 1, 33, '2025-04-15 06:21:54', '2025-04-15 06:21:54', 0);
INSERT INTO `sys_role_permission` VALUES (50, 1, 34, '2025-04-15 06:21:54', '2025-04-15 06:21:54', 0);
INSERT INTO `sys_role_permission` VALUES (51, 1, 35, '2025-04-15 06:21:54', '2025-04-15 06:21:54', 0);
INSERT INTO `sys_role_permission` VALUES (52, 1, 36, '2025-04-15 06:23:08', '2025-04-15 06:23:08', 0);
INSERT INTO `sys_role_permission` VALUES (53, 1, 37, '2025-04-18 07:41:52', '2025-04-18 07:41:52', 0);
INSERT INTO `sys_role_permission` VALUES (54, 1, 38, '2025-04-22 07:26:57', '2025-04-22 07:26:57', 0);
INSERT INTO `sys_role_permission` VALUES (55, 1, 39, '2025-04-22 07:49:47', '2025-04-22 07:49:47', 0);
INSERT INTO `sys_role_permission` VALUES (56, 1, 40, '2025-04-22 08:25:18', '2025-04-22 08:25:18', 0);
INSERT INTO `sys_role_permission` VALUES (57, 1, 41, '2025-04-22 09:15:29', '2025-04-22 09:15:29', 0);
INSERT INTO `sys_role_permission` VALUES (58, 1, 42, '2025-04-24 10:24:14', '2025-04-24 10:24:14', 0);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名称',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `nick_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `user_email` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `user_gender` tinyint(4) NULL DEFAULT NULL COMMENT '用户性别',
  `user_phone` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `salt` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'MD5的盐值',
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否启用(0:禁用,1:启用)',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id` DESC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '2d4b97b5f46ac80c82e8f8a5812ad89b082583129fa7b7fd7575b286647c9bb5', 'admin', '', 1, '', '2024-09-30 23:08:05', '2025-04-24 17:52:21', '2025-04-24 17:52:21', 'VECaJx', '1', 0);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `status` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否启用(0:禁用,1:启用)',
  `is_deleted` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除(0:否,1:是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户角色管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1, 1, '2025-03-12 16:24:27', '2025-03-12 16:24:29', '1', 0);

SET FOREIGN_KEY_CHECKS = 1;
