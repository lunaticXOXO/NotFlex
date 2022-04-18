-- phpMyAdmin SQL Dump
-- version 5.0.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 01, 2021 at 02:59 AM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_tubes_pbp`
--

-- --------------------------------------------------------

--
-- Table structure for table `historymovie`
--

CREATE TABLE `historymovie` (
  `user_id` int(10) NOT NULL,
  `movie_id` int(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `historymovie`
--

INSERT INTO `historymovie` (`user_id`, `movie_id`) VALUES
(575568, 101),
(575568, 103),
(921839, 103);

-- --------------------------------------------------------

--
-- Table structure for table `membership`
--

CREATE TABLE `membership` (
  `id` int(10) NOT NULL,
  `user_id` int(10) NOT NULL,
  `subsID` int(10) DEFAULT NULL,
  `CCNumber` varchar(16) DEFAULT NULL,
  `validThru` varchar(5) NOT NULL,
  `CVC` varchar(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `membership`
--

INSERT INTO `membership` (`id`, `user_id`, `subsID`, `CCNumber`, `validThru`, `CVC`) VALUES
(201, 575568, 11, '3027164822456721', '06/25', '848');

-- --------------------------------------------------------

--
-- Table structure for table `movie`
--

CREATE TABLE `movie` (
  `id` int(10) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `releaseDate` date DEFAULT NULL,
  `genre` varchar(255) DEFAULT NULL,
  `actor` varchar(255) NOT NULL,
  `director` varchar(255) DEFAULT NULL,
  `synopsis` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `movie`
--

INSERT INTO `movie` (`id`, `title`, `releaseDate`, `genre`, `actor`, `director`, `synopsis`) VALUES
(101, 'Kygo World Tour 2022', '2021-03-04', 'Comedy, Action', 'Dimitri, Martin, Brown', 'Tim Berg', 'This is kygo, over the world.'),
(102, 'Garrix World Tour 2021', '2021-05-17', 'Action', 'Alexander, Bob', 'Griffin', 'This is MG, over the world.'),
(103, 'Avengers', '2020-03-04', 'Action', 'Robert Downey Jr, Chris Evans', 'Calvin', 'The new Avengers.');

-- --------------------------------------------------------

--
-- Table structure for table `subscription`
--

CREATE TABLE `subscription` (
  `id` int(10) NOT NULL,
  `user_id` int(10) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `duration` varchar(255) DEFAULT NULL,
  `price` float DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `subscription`
--

INSERT INTO `subscription` (`id`, `user_id`, `type`, `duration`, `price`) VALUES
(11, 575568, 'Premium', '2021-05-30 17:22:15.9109615', 100000);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) NOT NULL,
  `fullName` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `dateofbirth` date DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `gender` varchar(10) DEFAULT NULL,
  `user_type` char(2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `fullName`, `email`, `password`, `dateofbirth`, `country`, `gender`, `user_type`) VALUES
(321777, 'Tim Berg', 'avicii@gmail.com', 'legends', '1997-02-13', 'Sweden', 'Male', 'A'),
(575568, 'Kygo Brown', 'kygo@gmail.com', 'oslo', '1997-04-14', 'Sweden', 'Male', 'M'),
(921839, 'ABCD', 'abcd@gmail.com', 'abcd', '2003-04-12', 'Sweden', 'Male', 'M');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `historymovie`
--
ALTER TABLE `historymovie`
  ADD KEY `movie_id` (`movie_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `membership`
--
ALTER TABLE `membership`
  ADD PRIMARY KEY (`id`),
  ADD KEY `subsID` (`subsID`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `movie`
--
ALTER TABLE `movie`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `subscription`
--
ALTER TABLE `subscription`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `subscription`
--
ALTER TABLE `subscription`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `historymovie`
--
ALTER TABLE `historymovie`
  ADD CONSTRAINT `historymovie_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `historymovie_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`id`);

--
-- Constraints for table `membership`
--
ALTER TABLE `membership`
  ADD CONSTRAINT `membership_ibfk_1` FOREIGN KEY (`subsID`) REFERENCES `subscription` (`id`),
  ADD CONSTRAINT `membership_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `subscription`
--
ALTER TABLE `subscription`
  ADD CONSTRAINT `subscription_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
