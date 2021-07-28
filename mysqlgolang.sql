-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Jul 28, 2021 at 10:04 AM
-- Server version: 5.7.32
-- PHP Version: 7.4.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `mysqlgolang`
--

-- --------------------------------------------------------

--
-- Table structure for table `contacts`
--

CREATE TABLE `contacts` (
  `id` int(11) NOT NULL,
  `contactname` varchar(255) NOT NULL,
  `phonenumber` varchar(12) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `contacts`
--

INSERT INTO `contacts` (`id`, `contactname`, `phonenumber`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Lucky', '081210295355', 'iamjustbe.lucky@gmail.com', '2021-07-28 15:42:05', '2021-07-28 15:42:05', '2021-07-28 16:04:07'),
(2, 'Abel', '082100005555', 'christabelvw10@gyahoo.com', '2021-07-28 15:42:40', '2021-07-28 15:55:13', NULL),
(3, 'ZZZZZ', '098712341234', 'zzzzzzzz@gmail.com', '2021-07-28 15:56:24', '2021-07-28 15:56:24', '2021-07-28 16:04:53'),
(4, 'YYYYY', '123456789101', 'yyyyyyyy@gmail.com', '2021-07-28 16:03:31', '2021-07-28 16:03:51', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `code` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `price` decimal(16,2) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `code`, `name`, `price`, `created_at`, `updated_at`, `deleted_at`) VALUES
(4, 'AH2489195', 'Product aaaa', '125000.00', '2021-07-28 14:31:31', '2021-07-28 14:41:59', '2021-07-28 16:49:12'),
(5, 'AH2489195', 'Product Two', '325000.00', '2021-07-28 14:32:22', '2021-07-28 14:32:22', '2021-07-28 16:47:26'),
(6, 'AL787481', 'Product Three', '75000.00', '2021-07-28 14:40:19', '2021-07-28 14:40:19', NULL),
(7, 'AR124952', 'Product Four', '20000.00', '2021-07-28 14:42:23', '2021-07-28 14:42:23', '2021-07-28 14:42:54'),
(8, 'A', 'Product aaaa', '125000.00', '2021-07-28 16:47:00', '2021-07-28 16:47:00', NULL),
(9, 'A', 'Product aaaa', '125000.00', '2021-07-28 16:47:38', '2021-07-28 16:47:38', '2021-07-28 16:48:38');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `contacts`
--
ALTER TABLE `contacts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `contacts`
--
ALTER TABLE `contacts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
