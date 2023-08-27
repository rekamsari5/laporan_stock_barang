/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:3306
 Source Schema         : InventoryManagement

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 27/08/2023 19:18:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for MasterCustomer
-- ----------------------------
DROP TABLE IF EXISTS `MasterCustomer`;
CREATE TABLE `MasterCustomer` (
  `CustomerPK` int NOT NULL,
  `CustomerName` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`CustomerPK`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of MasterCustomer
-- ----------------------------
BEGIN;
INSERT INTO `MasterCustomer` (`CustomerPK`, `CustomerName`) VALUES (1, 'Customer A');
INSERT INTO `MasterCustomer` (`CustomerPK`, `CustomerName`) VALUES (2, 'Customer B');
INSERT INTO `MasterCustomer` (`CustomerPK`, `CustomerName`) VALUES (3, 'Customer C');
COMMIT;

-- ----------------------------
-- Table structure for MasterProduct
-- ----------------------------
DROP TABLE IF EXISTS `MasterProduct`;
CREATE TABLE `MasterProduct` (
  `ProductPK` int NOT NULL,
  `ProductName` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`ProductPK`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of MasterProduct
-- ----------------------------
BEGIN;
INSERT INTO `MasterProduct` (`ProductPK`, `ProductName`) VALUES (1, 'Product A');
INSERT INTO `MasterProduct` (`ProductPK`, `ProductName`) VALUES (2, 'Product B');
INSERT INTO `MasterProduct` (`ProductPK`, `ProductName`) VALUES (3, 'Product C');
COMMIT;

-- ----------------------------
-- Table structure for MasterSupplier
-- ----------------------------
DROP TABLE IF EXISTS `MasterSupplier`;
CREATE TABLE `MasterSupplier` (
  `SupplierPK` int NOT NULL,
  `SupplierName` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`SupplierPK`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of MasterSupplier
-- ----------------------------
BEGIN;
INSERT INTO `MasterSupplier` (`SupplierPK`, `SupplierName`) VALUES (1, 'Supplier A');
INSERT INTO `MasterSupplier` (`SupplierPK`, `SupplierName`) VALUES (2, 'Supplier B');
INSERT INTO `MasterSupplier` (`SupplierPK`, `SupplierName`) VALUES (3, 'Supplier C');
COMMIT;

-- ----------------------------
-- Table structure for MasterWarehouse
-- ----------------------------
DROP TABLE IF EXISTS `MasterWarehouse`;
CREATE TABLE `MasterWarehouse` (
  `WhsPK` int NOT NULL,
  `WhsName` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  PRIMARY KEY (`WhsPK`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of MasterWarehouse
-- ----------------------------
BEGIN;
INSERT INTO `MasterWarehouse` (`WhsPK`, `WhsName`) VALUES (1, 'Gudang A');
INSERT INTO `MasterWarehouse` (`WhsPK`, `WhsName`) VALUES (2, 'Gudang B');
INSERT INTO `MasterWarehouse` (`WhsPK`, `WhsName`) VALUES (3, 'Gudang C');
COMMIT;

-- ----------------------------
-- Table structure for TransaksiPenerimaanBarangDetail
-- ----------------------------
DROP TABLE IF EXISTS `TransaksiPenerimaanBarangDetail`;
CREATE TABLE `TransaksiPenerimaanBarangDetail` (
  `TrxInDPK` int NOT NULL AUTO_INCREMENT,
  `TrxInIDF` varchar(255) DEFAULT NULL,
  `TrxInDProductIdf` int DEFAULT NULL,
  `TrxInDQtyDus` int DEFAULT NULL,
  `TrxInDQtyPcs` int DEFAULT NULL,
  PRIMARY KEY (`TrxInDPK`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of TransaksiPenerimaanBarangDetail
-- ----------------------------
BEGIN;
INSERT INTO `TransaksiPenerimaanBarangDetail` (`TrxInDPK`, `TrxInIDF`, `TrxInDProductIdf`, `TrxInDQtyDus`, `TrxInDQtyPcs`) VALUES (19, 'TRXIN0001', 1, 40, 70);
INSERT INTO `TransaksiPenerimaanBarangDetail` (`TrxInDPK`, `TrxInIDF`, `TrxInDProductIdf`, `TrxInDQtyDus`, `TrxInDQtyPcs`) VALUES (20, 'TRXIN0001', 2, 23, 30);
INSERT INTO `TransaksiPenerimaanBarangDetail` (`TrxInDPK`, `TrxInIDF`, `TrxInDProductIdf`, `TrxInDQtyDus`, `TrxInDQtyPcs`) VALUES (21, 'TRXIN0001', 3, 15, 20);
INSERT INTO `TransaksiPenerimaanBarangDetail` (`TrxInDPK`, `TrxInIDF`, `TrxInDProductIdf`, `TrxInDQtyDus`, `TrxInDQtyPcs`) VALUES (22, 'TRXIN0008', 1, 40, 70);
INSERT INTO `TransaksiPenerimaanBarangDetail` (`TrxInDPK`, `TrxInIDF`, `TrxInDProductIdf`, `TrxInDQtyDus`, `TrxInDQtyPcs`) VALUES (23, 'TRXIN0008', 2, 23, 30);
INSERT INTO `TransaksiPenerimaanBarangDetail` (`TrxInDPK`, `TrxInIDF`, `TrxInDProductIdf`, `TrxInDQtyDus`, `TrxInDQtyPcs`) VALUES (24, 'TRXIN0008', 3, 15, 20);
COMMIT;

-- ----------------------------
-- Table structure for TransaksiPenerimaanBarangHeader
-- ----------------------------
DROP TABLE IF EXISTS `TransaksiPenerimaanBarangHeader`;
CREATE TABLE `TransaksiPenerimaanBarangHeader` (
  `TrxInPK` int NOT NULL AUTO_INCREMENT,
  `TrxInNo` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `WhsIdf` int DEFAULT NULL,
  `TrxInDate` date DEFAULT NULL,
  `TrxInSuppIdf` int DEFAULT NULL,
  `TrxInCustomer` int DEFAULT NULL,
  PRIMARY KEY (`TrxInPK`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of TransaksiPenerimaanBarangHeader
-- ----------------------------
BEGIN;
INSERT INTO `TransaksiPenerimaanBarangHeader` (`TrxInPK`, `TrxInNo`, `WhsIdf`, `TrxInDate`, `TrxInSuppIdf`, `TrxInCustomer`) VALUES (7, 'TRXIN0001', 1, '2023-08-27', 2, 2);
INSERT INTO `TransaksiPenerimaanBarangHeader` (`TrxInPK`, `TrxInNo`, `WhsIdf`, `TrxInDate`, `TrxInSuppIdf`, `TrxInCustomer`) VALUES (8, 'TRXIN0008', 1, '2023-08-27', 2, 2);
COMMIT;

-- ----------------------------
-- Table structure for TransaksiPengeluaranBarangDetail
-- ----------------------------
DROP TABLE IF EXISTS `TransaksiPengeluaranBarangDetail`;
CREATE TABLE `TransaksiPengeluaranBarangDetail` (
  `TrxOutDPK` int NOT NULL AUTO_INCREMENT,
  `TrxOutIDF` varchar(255) DEFAULT NULL,
  `TrxOutDProductIdf` int DEFAULT NULL,
  `TrxOutDQtyDus` int DEFAULT NULL,
  `TrxOutDQtyPcs` int DEFAULT NULL,
  PRIMARY KEY (`TrxOutDPK`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of TransaksiPengeluaranBarangDetail
-- ----------------------------
BEGIN;
INSERT INTO `TransaksiPengeluaranBarangDetail` (`TrxOutDPK`, `TrxOutIDF`, `TrxOutDProductIdf`, `TrxOutDQtyDus`, `TrxOutDQtyPcs`) VALUES (25, 'TRXOUT0001', 1, 2, 5);
INSERT INTO `TransaksiPengeluaranBarangDetail` (`TrxOutDPK`, `TrxOutIDF`, `TrxOutDProductIdf`, `TrxOutDQtyDus`, `TrxOutDQtyPcs`) VALUES (26, 'TRXOUT0001', 2, 10, 6);
INSERT INTO `TransaksiPengeluaranBarangDetail` (`TrxOutDPK`, `TrxOutIDF`, `TrxOutDProductIdf`, `TrxOutDQtyDus`, `TrxOutDQtyPcs`) VALUES (27, 'TRXOUT0001', 3, 8, 5);
COMMIT;

-- ----------------------------
-- Table structure for TransaksiPengeluaranBarangHeader
-- ----------------------------
DROP TABLE IF EXISTS `TransaksiPengeluaranBarangHeader`;
CREATE TABLE `TransaksiPengeluaranBarangHeader` (
  `TrxOutPK` int NOT NULL AUTO_INCREMENT,
  `TrxOutNo` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `WhsIdf` int DEFAULT NULL,
  `TrxOutDate` date DEFAULT NULL,
  `TrxOutSuppIdf` int DEFAULT NULL,
  `TrxOutCustomer` int DEFAULT NULL,
  PRIMARY KEY (`TrxOutPK`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of TransaksiPengeluaranBarangHeader
-- ----------------------------
BEGIN;
INSERT INTO `TransaksiPengeluaranBarangHeader` (`TrxOutPK`, `TrxOutNo`, `WhsIdf`, `TrxOutDate`, `TrxOutSuppIdf`, `TrxOutCustomer`) VALUES (10, 'TRXOUT0001', 1, '2023-08-27', 2, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
