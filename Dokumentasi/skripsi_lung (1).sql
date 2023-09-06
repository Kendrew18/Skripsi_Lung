-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 06, 2023 at 07:42 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `skripsi_lung`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `id_admin` int(11) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`id_admin`, `username`, `password`) VALUES
(1, 'admin', 'admin');

-- --------------------------------------------------------

--
-- Table structure for table `detail_order`
--

CREATE TABLE `detail_order` (
  `id_detail_order` int(11) NOT NULL COMMENT 'Primary Key',
  `id_order` int(11) DEFAULT NULL,
  `id_stock` int(11) DEFAULT NULL,
  `jumlah` int(11) DEFAULT NULL,
  `satuan` varchar(255) DEFAULT 'NULL',
  `harga_jual` bigint(20) DEFAULT NULL,
  `sub_total` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `detail_order`
--

INSERT INTO `detail_order` (`id_detail_order`, `id_order`, `id_stock`, `jumlah`, `satuan`, `harga_jual`, `sub_total`) VALUES
(6, 7, 2, 20, 'pcs', 125000, 2500000);

-- --------------------------------------------------------

--
-- Table structure for table `detail_order_barang`
--

CREATE TABLE `detail_order_barang` (
  `id_detail_order_barang` int(11) NOT NULL COMMENT 'Primary Key',
  `id_detail_order` int(11) DEFAULT NULL,
  `id_ukuran` int(11) DEFAULT NULL,
  `jumlah` int(11) DEFAULT NULL,
  `satuan` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `detail_order_barang`
--

INSERT INTO `detail_order_barang` (`id_detail_order_barang`, `id_detail_order`, `id_ukuran`, `jumlah`, `satuan`) VALUES
(5, 6, 3, 10, 'pcs'),
(6, 6, 4, 10, 'pcs');

-- --------------------------------------------------------

--
-- Table structure for table `detail_retur`
--

CREATE TABLE `detail_retur` (
  `id_detail_retur` int(11) NOT NULL COMMENT 'Primary Key',
  `id_retur` int(11) DEFAULT NULL,
  `id_stock` int(11) DEFAULT NULL,
  `jumlah` int(11) DEFAULT NULL,
  `satuan` varchar(255) DEFAULT NULL,
  `harga_jual` bigint(20) DEFAULT NULL,
  `sub_total` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `detail_retur`
--

INSERT INTO `detail_retur` (`id_detail_retur`, `id_retur`, `id_stock`, `jumlah`, `satuan`, `harga_jual`, `sub_total`) VALUES
(3, 5, 2, 10, 'pcs', 125000, 1250000);

-- --------------------------------------------------------

--
-- Table structure for table `detail_retur_barang`
--

CREATE TABLE `detail_retur_barang` (
  `id_retur_barang` int(11) NOT NULL COMMENT 'Primary Key',
  `id_detail_retur` int(11) DEFAULT NULL,
  `id_ukuran` int(11) DEFAULT NULL,
  `jumlah` int(11) DEFAULT NULL,
  `satuan` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `detail_retur_barang`
--

INSERT INTO `detail_retur_barang` (`id_retur_barang`, `id_detail_retur`, `id_ukuran`, `jumlah`, `satuan`) VALUES
(3, 3, 3, 5, 'pcs'),
(4, 3, 4, 5, 'pcs');

-- --------------------------------------------------------

--
-- Table structure for table `order`
--

CREATE TABLE `order` (
  `id_order` int(11) NOT NULL COMMENT 'Primary Key',
  `id_pelanggan` int(11) NOT NULL,
  `id_sales` int(11) DEFAULT NULL,
  `tanggal_pesanan` date NOT NULL,
  `no_order` varchar(255) NOT NULL,
  `pembayaran` varchar(255) NOT NULL,
  `down_payment` varchar(255) NOT NULL,
  `tanggal_pembayaran` date NOT NULL,
  `tanggal_pengiriman` date NOT NULL,
  `catatan` mediumtext DEFAULT 'NULL'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `order`
--

INSERT INTO `order` (`id_order`, `id_pelanggan`, `id_sales`, `tanggal_pesanan`, `no_order`, `pembayaran`, `down_payment`, `tanggal_pembayaran`, `tanggal_pengiriman`, `catatan`) VALUES
(7, 1, 2, '2023-09-06', 'BRZ-1234567890123', 'CASH', 'DERFTT-1234565679', '2023-09-12', '0000-00-00', '');

-- --------------------------------------------------------

--
-- Table structure for table `pelanggan`
--

CREATE TABLE `pelanggan` (
  `id_pelanggan` int(11) NOT NULL COMMENT 'Primary Key',
  `nama_toko` varchar(255) NOT NULL,
  `no_telp` varchar(255) NOT NULL,
  `alamat` text NOT NULL,
  `kota` text NOT NULL,
  `provinsi` text NOT NULL,
  `nama_penanggungjawab` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pelanggan`
--

INSERT INTO `pelanggan` (`id_pelanggan`, `nama_toko`, `no_telp`, `alamat`, `kota`, `provinsi`, `nama_penanggungjawab`) VALUES
(1, 'toko maju jaya', '09867543799', 'Weleri', 'semarang', 'jawa tengah', 'Kepin');

-- --------------------------------------------------------

--
-- Table structure for table `retur`
--

CREATE TABLE `retur` (
  `id_retur` int(11) NOT NULL COMMENT 'Primary Key',
  `id_sales` int(11) DEFAULT NULL,
  `id_pelanggan` int(11) DEFAULT NULL,
  `no_order` varchar(255) DEFAULT NULL,
  `tanggal_retur` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `retur`
--

INSERT INTO `retur` (`id_retur`, `id_sales`, `id_pelanggan`, `no_order`, `tanggal_retur`) VALUES
(5, 2, 1, 'BRZ-1234567890123', '2023-09-18');

-- --------------------------------------------------------

--
-- Table structure for table `stock`
--

CREATE TABLE `stock` (
  `id` int(11) NOT NULL COMMENT 'Primary Key',
  `nama_barang` varchar(255) DEFAULT NULL,
  `harga` bigint(20) DEFAULT NULL,
  `jenis_barang` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `stock`
--

INSERT INTO `stock` (`id`, `nama_barang`, `harga`, `jenis_barang`) VALUES
(2, 'polo-shirt-merah', 125000, 'polo');

-- --------------------------------------------------------

--
-- Table structure for table `stock_to_ukuran`
--

CREATE TABLE `stock_to_ukuran` (
  `id_detail_ukuran` int(11) NOT NULL,
  `id_stock` int(11) DEFAULT NULL,
  `id_ukuran` int(11) DEFAULT NULL,
  `stock` double(225,1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `stock_to_ukuran`
--

INSERT INTO `stock_to_ukuran` (`id_detail_ukuran`, `id_stock`, `id_ukuran`, `stock`) VALUES
(5, 2, 1, 20.0),
(6, 2, 2, 20.0),
(7, 2, 3, 20.0),
(8, 2, 4, 20.0),
(9, 2, 5, 20.0),
(10, 2, 6, 20.0);

-- --------------------------------------------------------

--
-- Table structure for table `ukuran`
--

CREATE TABLE `ukuran` (
  `id_ukuran` int(11) NOT NULL,
  `jenis_ukuran` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ukuran`
--

INSERT INTO `ukuran` (`id_ukuran`, `jenis_ukuran`) VALUES
(1, 'XS'),
(2, 'S'),
(3, 'M'),
(4, 'L'),
(5, 'XL'),
(6, 'XXL');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id_user` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `nomor_hp` varchar(255) DEFAULT NULL,
  `bank` varchar(255) DEFAULT NULL,
  `no_rekening` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `longitude` double NOT NULL DEFAULT 0,
  `latitude` double NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id_user`, `name`, `alamat`, `nomor_hp`, `bank`, `no_rekening`, `username`, `password`, `longitude`, `latitude`) VALUES
(1, 'yong', 'abc', '081999999', 'bca', '82200000', 'yong123', 'anjay123', 0, 0),
(2, 'Denny', 'Babatan pantai', '1234567890', 'Mandiri', '2345678910', 'Levina', 'Charin', 0, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id_admin`);

--
-- Indexes for table `detail_order`
--
ALTER TABLE `detail_order`
  ADD PRIMARY KEY (`id_detail_order`),
  ADD KEY `id_order` (`id_order`),
  ADD KEY `id_stock` (`id_stock`);

--
-- Indexes for table `detail_order_barang`
--
ALTER TABLE `detail_order_barang`
  ADD PRIMARY KEY (`id_detail_order_barang`),
  ADD KEY `id_detail_order` (`id_detail_order`),
  ADD KEY `id_ukuran` (`id_ukuran`);

--
-- Indexes for table `detail_retur`
--
ALTER TABLE `detail_retur`
  ADD PRIMARY KEY (`id_detail_retur`);

--
-- Indexes for table `detail_retur_barang`
--
ALTER TABLE `detail_retur_barang`
  ADD PRIMARY KEY (`id_retur_barang`);

--
-- Indexes for table `order`
--
ALTER TABLE `order`
  ADD PRIMARY KEY (`id_order`),
  ADD UNIQUE KEY `no_order` (`no_order`),
  ADD KEY `id_pelanggan` (`id_pelanggan`);

--
-- Indexes for table `pelanggan`
--
ALTER TABLE `pelanggan`
  ADD PRIMARY KEY (`id_pelanggan`);

--
-- Indexes for table `retur`
--
ALTER TABLE `retur`
  ADD PRIMARY KEY (`id_retur`),
  ADD UNIQUE KEY `no_order` (`no_order`);

--
-- Indexes for table `stock`
--
ALTER TABLE `stock`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `stock_to_ukuran`
--
ALTER TABLE `stock_to_ukuran`
  ADD PRIMARY KEY (`id_detail_ukuran`),
  ADD KEY `id_stock` (`id_stock`),
  ADD KEY `id_ukuran` (`id_ukuran`);

--
-- Indexes for table `ukuran`
--
ALTER TABLE `ukuran`
  ADD PRIMARY KEY (`id_ukuran`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id_user`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `detail_order`
--
ALTER TABLE `detail_order`
  MODIFY `id_detail_order` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key', AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `detail_order_barang`
--
ALTER TABLE `detail_order_barang`
  MODIFY `id_detail_order_barang` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key', AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `detail_retur`
--
ALTER TABLE `detail_retur`
  MODIFY `id_detail_retur` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key', AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `detail_retur_barang`
--
ALTER TABLE `detail_retur_barang`
  MODIFY `id_retur_barang` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key', AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `order`
--
ALTER TABLE `order`
  MODIFY `id_order` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key', AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `pelanggan`
--
ALTER TABLE `pelanggan`
  MODIFY `id_pelanggan` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key', AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `retur`
--
ALTER TABLE `retur`
  MODIFY `id_retur` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key', AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `stock`
--
ALTER TABLE `stock`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key', AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `stock_to_ukuran`
--
ALTER TABLE `stock_to_ukuran`
  MODIFY `id_detail_ukuran` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `ukuran`
--
ALTER TABLE `ukuran`
  MODIFY `id_ukuran` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `detail_order`
--
ALTER TABLE `detail_order`
  ADD CONSTRAINT `detail_order_ibfk_1` FOREIGN KEY (`id_order`) REFERENCES `order` (`id_order`),
  ADD CONSTRAINT `detail_order_ibfk_2` FOREIGN KEY (`id_stock`) REFERENCES `stock` (`id`);

--
-- Constraints for table `detail_order_barang`
--
ALTER TABLE `detail_order_barang`
  ADD CONSTRAINT `detail_order_barang_ibfk_1` FOREIGN KEY (`id_detail_order`) REFERENCES `detail_order` (`id_detail_order`),
  ADD CONSTRAINT `detail_order_barang_ibfk_2` FOREIGN KEY (`id_ukuran`) REFERENCES `ukuran` (`id_ukuran`);

--
-- Constraints for table `order`
--
ALTER TABLE `order`
  ADD CONSTRAINT `order_ibfk_1` FOREIGN KEY (`id_pelanggan`) REFERENCES `pelanggan` (`id_pelanggan`);

--
-- Constraints for table `stock_to_ukuran`
--
ALTER TABLE `stock_to_ukuran`
  ADD CONSTRAINT `stock_to_ukuran_ibfk_1` FOREIGN KEY (`id_stock`) REFERENCES `stock` (`id`),
  ADD CONSTRAINT `stock_to_ukuran_ibfk_2` FOREIGN KEY (`id_ukuran`) REFERENCES `ukuran` (`id_ukuran`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
