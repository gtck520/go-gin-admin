/*
Navicat MySQL Data Transfer

Source Server         : 本地mysql7
Source Server Version : 50727
Source Host           : localhost:3306
Source Database       : gin-vue

Target Server Type    : MYSQL
Target Server Version : 50727
File Encoding         : 65001

Date: 2021-02-23 20:31:55
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for go_sys_admins
-- ----------------------------
DROP TABLE IF EXISTS `go_sys_admins`;
CREATE TABLE `go_sys_admins` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `updated_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `memo` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `real_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` char(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` char(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_admins_user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of go_sys_admins
-- ----------------------------
INSERT INTO `go_sys_admins` VALUES ('1', '2021-01-25 22:34:47', '2021-01-25 22:34:47', '0', '0', '', 'konger', 'kc', '900963658df8cd586cf9f31fe665acf7', 'kc', 'kc', '1');

-- ----------------------------
-- Table structure for go_sys_admins_role
-- ----------------------------
DROP TABLE IF EXISTS `go_sys_admins_role`;
CREATE TABLE `go_sys_admins_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `updated_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `admins_id` bigint(20) unsigned NOT NULL,
  `role_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_admins_role_admins_id` (`admins_id`,`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of go_sys_admins_role
-- ----------------------------
INSERT INTO `go_sys_admins_role` VALUES ('2', '2021-02-21 22:15:01', '2021-02-21 22:15:01', '0', '0', '1', '1');

-- ----------------------------
-- Table structure for go_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `go_sys_menu`;
CREATE TABLE `go_sys_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `updated_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL,
  `memo` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `parent_id` bigint(20) unsigned NOT NULL,
  `url` varchar(72) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sequence` int(11) NOT NULL,
  `menu_type` tinyint(1) NOT NULL,
  `code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `operate_type` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_menu_code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of go_sys_menu
-- ----------------------------
INSERT INTO `go_sys_menu` VALUES ('1', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '0', '', 'TOP', '1', '1', 'TOP', '', 'none');
INSERT INTO `go_sys_menu` VALUES ('2', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '1', '', '系统管理', '1', '1', 'Sys', 'lock', 'none');
INSERT INTO `go_sys_menu` VALUES ('3', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '2', '/icon', '图标管理', '10', '2', 'Icon', 'icon', 'none');
INSERT INTO `go_sys_menu` VALUES ('4', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '2', '/menu', '菜单管理', '20', '2', 'Menu', 'documentation', 'none');
INSERT INTO `go_sys_menu` VALUES ('5', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '4', '/menu/create', '新增', '1', '3', 'MenuAdd', '', 'add');
INSERT INTO `go_sys_menu` VALUES ('6', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '4', '/menu/delete', '删除', '2', '3', 'MenuDel', '', 'del');
INSERT INTO `go_sys_menu` VALUES ('7', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '4', '/menu/detail', '查看', '3', '3', 'MenuView', '', 'view');
INSERT INTO `go_sys_menu` VALUES ('8', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '4', '/menu/update', '编辑', '4', '3', 'MenuUpdate', '', 'update');
INSERT INTO `go_sys_menu` VALUES ('9', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '4', '/menu/list', '分页api', '5', '3', 'MenuList', '', 'list');
INSERT INTO `go_sys_menu` VALUES ('10', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '2', '/role', '角色管理', '30', '2', 'Role', 'tree', 'none');
INSERT INTO `go_sys_menu` VALUES ('11', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '10', '/role/create', '新增', '1', '3', 'RoleAdd', '', 'add');
INSERT INTO `go_sys_menu` VALUES ('12', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '10', '/role/delete', '删除', '2', '3', 'RoleDel', '', 'del');
INSERT INTO `go_sys_menu` VALUES ('13', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '10', '/role/detail', '查看', '3', '3', 'RoleView', '', 'view');
INSERT INTO `go_sys_menu` VALUES ('14', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '10', '/role/update', '编辑', '4', '3', 'RoleUpdate', '', 'update');
INSERT INTO `go_sys_menu` VALUES ('15', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '10', '/role/list', '分页api', '5', '3', 'RoleList', '', 'list');
INSERT INTO `go_sys_menu` VALUES ('16', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '10', '/role/setrole', '分配角色菜单', '6', '3', 'RoleSetrolemenu', '', 'setrolemenu');
INSERT INTO `go_sys_menu` VALUES ('17', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '2', '/admins', '后台用户管理', '40', '2', 'Admins', 'user', 'none');
INSERT INTO `go_sys_menu` VALUES ('18', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '17', '/admins/create', '新增', '1', '3', 'AdminsAdd', '', 'add');
INSERT INTO `go_sys_menu` VALUES ('19', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '17', '/admins/delete', '删除', '2', '3', 'AdminsDel', '', 'del');
INSERT INTO `go_sys_menu` VALUES ('20', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '17', '/admins/detail', '查看', '3', '3', 'AdminsView', '', 'view');
INSERT INTO `go_sys_menu` VALUES ('21', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '17', '/admins/update', '编辑', '4', '3', 'AdminsUpdate', '', 'update');
INSERT INTO `go_sys_menu` VALUES ('22', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '17', '/admins/list', '分页api', '5', '3', 'AdminsList', '', 'list');
INSERT INTO `go_sys_menu` VALUES ('23', '2021-01-25 21:50:40', '2021-01-25 21:50:40', '0', '0', '1', '', '17', '/admins/setrole', '分配角色', '6', '3', 'AdminsSetrole', '', 'setadminrole');

-- ----------------------------
-- Table structure for go_sys_role
-- ----------------------------
DROP TABLE IF EXISTS `go_sys_role`;
CREATE TABLE `go_sys_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `updated_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `memo` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sequence` int(11) NOT NULL,
  `parent_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of go_sys_role
-- ----------------------------
INSERT INTO `go_sys_role` VALUES ('1', '2021-01-25 22:36:00', '2021-01-25 22:36:00', '0', '0', '', '超级管理员', '1', '0');

-- ----------------------------
-- Table structure for go_sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `go_sys_role_menu`;
CREATE TABLE `go_sys_role_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `updated_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `role_id` bigint(20) unsigned NOT NULL,
  `menu_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_menu_role_id` (`role_id`,`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of go_sys_role_menu
-- ----------------------------
INSERT INTO `go_sys_role_menu` VALUES ('24', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '1');
INSERT INTO `go_sys_role_menu` VALUES ('25', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '2');
INSERT INTO `go_sys_role_menu` VALUES ('26', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '3');
INSERT INTO `go_sys_role_menu` VALUES ('27', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '4');
INSERT INTO `go_sys_role_menu` VALUES ('28', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '5');
INSERT INTO `go_sys_role_menu` VALUES ('29', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '6');
INSERT INTO `go_sys_role_menu` VALUES ('30', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '7');
INSERT INTO `go_sys_role_menu` VALUES ('31', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '8');
INSERT INTO `go_sys_role_menu` VALUES ('32', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '9');
INSERT INTO `go_sys_role_menu` VALUES ('33', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '10');
INSERT INTO `go_sys_role_menu` VALUES ('34', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '11');
INSERT INTO `go_sys_role_menu` VALUES ('35', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '12');
INSERT INTO `go_sys_role_menu` VALUES ('36', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '13');
INSERT INTO `go_sys_role_menu` VALUES ('37', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '14');
INSERT INTO `go_sys_role_menu` VALUES ('38', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '15');
INSERT INTO `go_sys_role_menu` VALUES ('39', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '16');
INSERT INTO `go_sys_role_menu` VALUES ('40', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '17');
INSERT INTO `go_sys_role_menu` VALUES ('41', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '18');
INSERT INTO `go_sys_role_menu` VALUES ('42', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '19');
INSERT INTO `go_sys_role_menu` VALUES ('43', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '20');
INSERT INTO `go_sys_role_menu` VALUES ('44', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '21');
INSERT INTO `go_sys_role_menu` VALUES ('45', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '22');
INSERT INTO `go_sys_role_menu` VALUES ('46', '2021-02-22 21:13:10', '2021-02-22 21:13:10', '0', '0', '1', '23');
