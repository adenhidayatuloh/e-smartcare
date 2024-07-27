-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 27, 2024 at 10:03 AM
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
-- Database: `db_smartcare`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `email` varchar(100) NOT NULL,
  `nama_lengkap` varchar(100) DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `no_telepon` varchar(15) DEFAULT NULL,
  `foto_profil` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`email`, `nama_lengkap`, `alamat`, `no_telepon`, `foto_profil`) VALUES
('adminku@gmail.com', 'admin HT', 'Ajibarang', '0868977896', '/uploads/adminku@gmail.com.jpeg');

-- --------------------------------------------------------

--
-- Table structure for table `alarm`
--

CREATE TABLE `alarm` (
  `id` int(11) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `keterangan` varchar(45) DEFAULT NULL,
  `tanggal_mulai` date DEFAULT NULL,
  `jam` time DEFAULT NULL,
  `pengulangan` int(11) DEFAULT NULL,
  `status` varchar(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `alarm`
--

INSERT INTO `alarm` (`id`, `email`, `keterangan`, `tanggal_mulai`, `jam`, `pengulangan`, `status`) VALUES
(2, 'user3@gmail.com', 'Alarm Berangkat ngaji', '2023-11-10', '08:00:00', 10, '1'),
(3, 'user3@gmail.com', 'Alarm Bangun Makan', '2023-10-10', '03:00:00', 2, '1'),
(5, 'user3@gmail.com', 'Alarm Berangkat ngaji', '2023-11-10', '08:00:00', 10, '1'),
(6, 'user3@gmail.com', 'Alarm Berangkat ngaji', '2023-11-10', '08:00:00', 10, '1');

-- --------------------------------------------------------

--
-- Table structure for table `pakar`
--

CREATE TABLE `pakar` (
  `email` varchar(100) NOT NULL,
  `nama_lengkap` varchar(100) DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `no_telepon` varchar(15) DEFAULT NULL,
  `foto_profil` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pakar`
--

INSERT INTO `pakar` (`email`, `nama_lengkap`, `alamat`, `no_telepon`, `foto_profil`) VALUES
('pakarke2@gmail.com', 'pakar ku', 'Ajibarang', '0868977896', '/uploads/pakarke2@gmail.com.jpeg');

-- --------------------------------------------------------

--
-- Table structure for table `pemeriksaan`
--

CREATE TABLE `pemeriksaan` (
  `email` varchar(100) NOT NULL,
  `waktu` datetime DEFAULT NULL,
  `foto` text DEFAULT NULL,
  `tinggi` double DEFAULT NULL,
  `berat` double DEFAULT NULL,
  `keterangan` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pemeriksaan`
--

INSERT INTO `pemeriksaan` (`email`, `waktu`, `foto`, `tinggi`, `berat`, `keterangan`) VALUES
('samsul@gmail.com', '2024-07-27 14:13:08', '/uploads/pemeriksaan/samsul@gmail.com.jpg', 200, 90, 'stunting');

-- --------------------------------------------------------

--
-- Table structure for table `riwayat_tanya_jawab`
--

CREATE TABLE `riwayat_tanya_jawab` (
  `id` int(11) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `waktu` datetime DEFAULT NULL,
  `pertanyaan` text DEFAULT NULL,
  `jawaban` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `riwayat_tanya_jawab`
--

INSERT INTO `riwayat_tanya_jawab` (`id`, `email`, `waktu`, `pertanyaan`, `jawaban`) VALUES
(10, 'samsul@gmail.com', '2024-07-27 13:08:48', 'tes tambah  riwayat pertanyaan', 'tes tambah riwayat tanya jawdwab'),
(11, 'samsul@gmail.com', '2024-07-27 13:19:20', 'tes tambah  riwayat pertanyaan', 'tes tambah riwayat jawaban');

-- --------------------------------------------------------

--
-- Table structure for table `siswa`
--

CREATE TABLE `siswa` (
  `email` varchar(100) NOT NULL,
  `nis` varchar(45) DEFAULT NULL,
  `nama_lengkap` varchar(100) DEFAULT NULL,
  `tempat_lahir` varchar(100) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `no_telepon` varchar(15) DEFAULT NULL,
  `kelas` varchar(10) DEFAULT NULL,
  `agama` varchar(20) DEFAULT NULL,
  `foto_profil` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `siswa`
--

INSERT INTO `siswa` (`email`, `nis`, `nama_lengkap`, `tempat_lahir`, `tanggal_lahir`, `alamat`, `no_telepon`, `kelas`, `agama`, `foto_profil`) VALUES
('samsul@gmail.com', '21SA1251', 'Aden HT', 'purwokerto', '2003-10-10', 'Ajibarang', '0868977896', 'Ipa 1', 'islam', '/uploads/samsul@gmail.com.jpeg'),
('user3@gmail.com', '21SA1251', 'Aden', 'purwokerto', '2003-10-10', '', '', '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `tanya_jawab`
--

CREATE TABLE `tanya_jawab` (
  `id_tanya_jawab` int(11) NOT NULL,
  `pertanyaan` text DEFAULT NULL,
  `jawaban` text DEFAULT NULL,
  `validator` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `tanya_jawab`
--

INSERT INTO `tanya_jawab` (`id_tanya_jawab`, `pertanyaan`, `jawaban`, `validator`) VALUES
(1, 'Apa itu stunting?', 'Stunting adalah kondisi gagal tumbuh pada anak balita akibat kekurangan gizi kronis, infeksi berulang, dan stimulasi psikososial yang tidak memadai.', 'adminku@gmail.com'),
(2, 'Apa penyebab utama stunting?', 'Penyebab utama stunting adalah asupan gizi yang tidak memadai, terutama dalam 1.000 hari pertama kehidupan, infeksi berulang, dan kurangnya stimulasi psikososial.', 'adminku@gmail.com'),
(3, 'Bagaimana cara mencegah stunting?', 'Stunting dapat dicegah dengan memberikan asupan gizi yang baik sejak kehamilan hingga usia 2 tahun, akses ke layanan kesehatan yang memadai, sanitasi yang baik, dan stimulasi psikososial yang cukup.', 'adminku@gmail.com'),
(4, 'Apa dampak jangka panjang stunting?', 'Dampak jangka panjang stunting meliputi terganggunya perkembangan kognitif dan motorik, prestasi pendidikan yang rendah, dan risiko lebih tinggi untuk penyakit kronis di masa dewasa.', 'adminku@gmail.com'),
(5, 'Bagaimana cara mengatasi stunting pada anak?', 'Mengatasi stunting pada anak memerlukan intervensi gizi, pengobatan infeksi, peningkatan sanitasi dan kebersihan, serta stimulasi psikososial yang tepat.', 'adminku@gmail.com'),
(6, 'Apa program pemerintah untuk mengatasi stunting?', 'Program pemerintah untuk mengatasi stunting mencakup inisiatif seperti Pemberian Makanan Tambahan (PMT), imunisasi, peningkatan sanitasi, dan program pendidikan bagi orang tua mengenai pentingnya gizi dan perawatan anak.', 'adminku@gmail.com'),
(7, 'Apa tanda-tanda anak mengalami stunting?', 'Tanda-tanda anak mengalami stunting meliputi tinggi badan yang lebih pendek dari anak seusianya, perkembangan fisik yang lambat, dan sering mengalami infeksi.', 'adminku@gmail.com'),
(8, 'Apa peran ibu hamil dalam pencegahan stunting?', 'Peran ibu hamil dalam pencegahan stunting sangat penting, termasuk menjaga asupan gizi yang cukup, rutin memeriksakan kehamilan, dan menghindari faktor risiko seperti merokok dan konsumsi alkohol.', 'adminku@gmail.com'),
(9, 'Bagaimana peran ASI dalam pencegahan stunting?', 'ASI eksklusif selama 6 bulan pertama kehidupan sangat penting untuk pencegahan stunting, karena ASI mengandung semua nutrisi yang diperlukan bayi serta antibodi untuk melindungi dari infeksi.', 'adminku@gmail.com'),
(10, 'Apa itu intervensi 1000 HPK?', 'Intervensi 1000 Hari Pertama Kehidupan (HPK) adalah upaya untuk memastikan asupan gizi yang optimal dari masa kehamilan hingga anak berusia 2 tahun, periode yang sangat penting untuk mencegah stunting.', 'adminku@gmail.com'),
(11, 'Bagaimana stunting mempengaruhi perkembangan otak anak?', 'Stunting dapat menghambat perkembangan otak anak, mengurangi kemampuan belajar dan konsentrasi, serta meningkatkan risiko gangguan perilaku.', 'adminku@gmail.com'),
(12, 'Apa hubungan antara stunting dan kemiskinan?', 'Kemiskinan sering kali menyebabkan kurangnya akses terhadap makanan bergizi, layanan kesehatan, dan sanitasi, yang semuanya berkontribusi terhadap stunting.', 'adminku@gmail.com'),
(13, 'Apa peran ayah dalam pencegahan stunting?', 'Ayah memiliki peran penting dalam pencegahan stunting dengan mendukung nutrisi ibu dan anak, menyediakan lingkungan yang sehat, dan memastikan akses ke layanan kesehatan.', 'adminku@gmail.com'),
(14, 'Bagaimana sanitasi buruk berkontribusi pada stunting?', 'Sanitasi buruk meningkatkan risiko infeksi dan diare, yang dapat menghambat penyerapan nutrisi dan berkontribusi terhadap stunting.', 'adminku@gmail.com'),
(15, 'Apa itu wasting dan bagaimana hubungannya dengan stunting?', 'Wasting adalah kondisi kekurangan berat badan untuk tinggi badan yang sering terjadi akibat kekurangan gizi akut. Meskipun berbeda, wasting dan stunting sering terjadi bersamaan dan saling memperburuk kondisi satu sama lain.', 'adminku@gmail.com'),
(16, 'Bagaimana pola makan ibu hamil mempengaruhi risiko stunting?', 'Pola makan ibu hamil yang kurang gizi dapat menyebabkan pertumbuhan janin yang terhambat, meningkatkan risiko stunting pada anak setelah lahir.', 'adminku@gmail.com'),
(17, 'Bagaimana infeksi berulang dapat menyebabkan stunting?', 'Infeksi berulang, seperti diare dan infeksi saluran pernapasan, dapat mengganggu penyerapan nutrisi dan menghambat pertumbuhan anak, yang pada akhirnya menyebabkan stunting.', 'adminku@gmail.com'),
(18, 'Apa itu fortifikasi makanan dan bagaimana dapat membantu mencegah stunting?', 'Fortifikasi makanan adalah proses menambahkan mikronutrien penting ke dalam makanan. Ini dapat membantu mencegah kekurangan gizi yang berkontribusi terhadap stunting.', 'adminku@gmail.com'),
(19, 'Bagaimana kebijakan kesehatan masyarakat dapat membantu mengurangi stunting?', 'Kebijakan kesehatan masyarakat yang efektif dapat meningkatkan akses terhadap makanan bergizi, layanan kesehatan, sanitasi, dan pendidikan, yang semuanya penting untuk mencegah stunting.', 'adminku@gmail.com'),
(20, 'Apa peran pendidikan ibu dalam pencegahan stunting?', 'Ibu yang terdidik lebih cenderung memiliki pengetahuan tentang gizi, sanitasi, dan perawatan kesehatan, yang semuanya berkontribusi pada pencegahan stunting.', 'adminku@gmail.com'),
(21, 'Bagaimana akses ke air bersih dapat mencegah stunting?', 'Akses ke air bersih mengurangi risiko infeksi yang dapat menghambat penyerapan nutrisi dan pertumbuhan anak, sehingga membantu mencegah stunting.', 'adminku@gmail.com'),
(22, 'Apa peran vitamin dan mineral dalam pencegahan stunting?', 'Vitamin dan mineral seperti vitamin A, zinc, dan zat besi sangat penting untuk pertumbuhan dan perkembangan anak. Kekurangan mikronutrien ini dapat berkontribusi pada stunting.', 'adminku@gmail.com'),
(23, 'Bagaimana perubahan iklim dapat mempengaruhi stunting?', 'Perubahan iklim dapat mempengaruhi ketersediaan pangan, meningkatkan risiko bencana alam, dan menyulitkan akses ke air bersih, yang semuanya dapat meningkatkan risiko stunting.', 'adminku@gmail.com'),
(24, 'Bagaimana peran komunitas dalam pencegahan stunting?', 'Komunitas dapat berperan dengan meningkatkan kesadaran tentang pentingnya gizi, sanitasi, dan perawatan kesehatan, serta menyediakan dukungan sosial bagi keluarga yang berisiko stunting.', 'adminku@gmail.com'),
(25, 'Apa itu intervensi gizi sensitif dan spesifik dalam konteks stunting?', 'Intervensi gizi sensitif adalah tindakan yang mempengaruhi faktor penyebab stunting secara tidak langsung, seperti sanitasi dan pendidikan. Intervensi gizi spesifik langsung terkait dengan peningkatan status gizi, seperti suplementasi dan fortifikasi.', 'adminku@gmail.com'),
(26, 'Bagaimana stunting mempengaruhi produktivitas ekonomi di masa depan?', 'Stunting dapat mengurangi kemampuan kognitif dan kesehatan seseorang, yang berdampak pada produktivitas dan potensi penghasilan di masa dewasa, sehingga mempengaruhi ekonomi secara keseluruhan.', 'adminku@gmail.com'),
(27, 'Apa itu program PMTCT dan bagaimana kaitannya dengan stunting?', 'PMTCT (Prevention of Mother-To-Child Transmission) adalah program untuk mencegah penularan HIV dari ibu ke anak. Infeksi HIV dapat memperburuk stunting, jadi pencegahan penularan penting untuk mengurangi risiko stunting.', 'adminku@gmail.com'),
(28, 'Bagaimana teknologi dapat digunakan untuk mengatasi stunting?', 'Teknologi dapat digunakan untuk mengatasi stunting melalui aplikasi pemantauan gizi, edukasi digital untuk ibu, dan distribusi makanan fortifikasi yang lebih efisien.', 'adminku@gmail.com'),
(29, 'Bagaimana pandemi COVID-19 mempengaruhi upaya pencegahan stunting?', 'Pandemi COVID-19 telah mengganggu layanan kesehatan, memperburuk kemiskinan, dan meningkatkan kerawanan pangan, yang semuanya dapat meningkatkan risiko stunting.', 'adminku@gmail.com'),
(30, 'Bagaimana pemberian makan tambahan (PMT) dapat membantu mencegah stunting?', 'Pemberian Makan Tambahan (PMT) menyediakan asupan nutrisi tambahan untuk anak-anak yang kekurangan gizi, membantu memenuhi kebutuhan gizi mereka dan mencegah stunting.', 'adminku@gmail.com'),
(31, 'Apa peran lembaga internasional dalam pencegahan stunting?', 'Lembaga internasional seperti WHO dan UNICEF berperan dalam pencegahan stunting dengan memberikan pedoman, dukungan teknis, dan dana untuk program gizi dan kesehatan.', 'adminku@gmail.com'),
(32, 'Bagaimana pengaruh gizi buruk selama kehamilan terhadap stunting?', 'Gizi buruk selama kehamilan dapat menghambat pertumbuhan janin, meningkatkan risiko bayi lahir dengan berat badan rendah dan mengalami stunting di masa kecil.', 'adminku@gmail.com'),
(33, 'Bagaimana peran bidan dalam pencegahan stunting?', 'Bidan berperan penting dalam pencegahan stunting dengan memberikan edukasi gizi kepada ibu hamil, memantau pertumbuhan bayi, dan memberikan imunisasi serta intervensi kesehatan lainnya.', 'adminku@gmail.com'),
(34, 'Bagaimana masyarakat dapat mendukung program pencegahan stunting?', 'Masyarakat dapat mendukung program pencegahan stunting dengan terlibat dalam kampanye kesadaran, mendukung kebijakan kesehatan yang baik, dan membantu memastikan akses terhadap makanan bergizi dan layanan kesehatan bagi semua keluarga.', NULL),
(35, 'Apa hubungan antara imunisasi dan pencegahan stunting?', 'Imunisasi melindungi anak dari penyakit yang dapat mengganggu penyerapan nutrisi dan pertumbuhan, sehingga membantu mencegah stunting.', NULL),
(36, 'Bagaimana pengaruh pernikahan dini terhadap risiko stunting?', 'Pernikahan dini sering kali dikaitkan dengan kehamilan remaja, yang meningkatkan risiko kekurangan gizi ibu dan bayi, serta risiko stunting pada anak yang dilahirkan.', NULL),
(37, 'Apa itu Growth Monitoring Promotion (GMP) dan bagaimana kaitannya dengan pencegahan stunting?', 'Growth Monitoring Promotion (GMP) adalah program yang memantau pertumbuhan anak secara rutin dan memberikan edukasi kepada orang tua tentang gizi dan kesehatan, yang penting untuk mencegah stunting.', NULL),
(40, 'tes tanya', 'tes jawab', 'pakarke2@gmail.com'),
(41, 'tes update tanya jawdwab', 'tes update tanya jawdwab', 'pakarke2@gmail.com'),
(43, 'tes tanya', 'tes jawab', 'pakarke2@gmail.com'),
(45, 'tes tanya1', 'tes jawab1', ''),
(47, 'tes tanya3', 'tes jawab3', ''),
(48, 'tes tambah tanya jawdwab', 'tes tambah tanya jawdwab', ''),
(49, 'tes tambah tanya jawdwab', 'tes tambah tanya jawdwab', '');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `email` varchar(100) NOT NULL,
  `password` varchar(246) DEFAULT NULL,
  `jenis_akun` varchar(1) DEFAULT NULL,
  `request_jenis_akun` varchar(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`email`, `password`, `jenis_akun`, `request_jenis_akun`) VALUES
('admin1@gmail.com', '$2a$10$JBcTuKsFLYEtWegbwLa/VehDmW/gS8sfKMiyf3gedffOvOOfhKcRO', '1', '1'),
('admin3@gmail.com', '$2a$10$7FNRG14xpxV7z5ffdgRyDOAWhjRvU6ggHrCb.9ZwkmyEf0VJrj6Fm', '1', '1'),
('adminku@gmail.com', '$2a$10$N5mVUaL3Ht/mDjgZ6cVTzuUGOyGTMF8LjwLI7u6Ex0gFF11ZA8gg.', '1', '1'),
('pakarke2@gmail.com', '$2a$10$ckITqj/ZNSdKiwdFPlIS6uhvN27PnKe6f/JPe89dpC8FA56QnMaTa', '2', '2'),
('samsul@gmail.com', '$2a$10$ycE4AC58nr8cPhJOpb9YqeQBWZ9vjYiKYHmqrb7TGyAiqU8kYWQRi', '3', '3'),
('user3@gmail.com', '$2a$10$C/fHJQql8bzaLh29YsSUz.SUYpI2b9b262kQOM12iSmQzlfRDAK3i', '3', '3');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`email`);

--
-- Indexes for table `alarm`
--
ALTER TABLE `alarm`
  ADD PRIMARY KEY (`id`),
  ADD KEY `email` (`email`);

--
-- Indexes for table `pakar`
--
ALTER TABLE `pakar`
  ADD PRIMARY KEY (`email`);

--
-- Indexes for table `pemeriksaan`
--
ALTER TABLE `pemeriksaan`
  ADD PRIMARY KEY (`email`);

--
-- Indexes for table `riwayat_tanya_jawab`
--
ALTER TABLE `riwayat_tanya_jawab`
  ADD PRIMARY KEY (`id`),
  ADD KEY `email` (`email`);

--
-- Indexes for table `siswa`
--
ALTER TABLE `siswa`
  ADD PRIMARY KEY (`email`);

--
-- Indexes for table `tanya_jawab`
--
ALTER TABLE `tanya_jawab`
  ADD PRIMARY KEY (`id_tanya_jawab`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `alarm`
--
ALTER TABLE `alarm`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `riwayat_tanya_jawab`
--
ALTER TABLE `riwayat_tanya_jawab`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `tanya_jawab`
--
ALTER TABLE `tanya_jawab`
  MODIFY `id_tanya_jawab` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=50;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `admin`
--
ALTER TABLE `admin`
  ADD CONSTRAINT `admin_ibfk_1` FOREIGN KEY (`email`) REFERENCES `user` (`email`) ON DELETE CASCADE;

--
-- Constraints for table `alarm`
--
ALTER TABLE `alarm`
  ADD CONSTRAINT `alarm_ibfk_1` FOREIGN KEY (`email`) REFERENCES `siswa` (`email`) ON DELETE CASCADE;

--
-- Constraints for table `pakar`
--
ALTER TABLE `pakar`
  ADD CONSTRAINT `pakar_ibfk_1` FOREIGN KEY (`email`) REFERENCES `user` (`email`) ON DELETE CASCADE;

--
-- Constraints for table `pemeriksaan`
--
ALTER TABLE `pemeriksaan`
  ADD CONSTRAINT `pemeriksaan_ibfk_1` FOREIGN KEY (`email`) REFERENCES `siswa` (`email`) ON DELETE CASCADE;

--
-- Constraints for table `riwayat_tanya_jawab`
--
ALTER TABLE `riwayat_tanya_jawab`
  ADD CONSTRAINT `riwayat_tanya_jawab_ibfk_1` FOREIGN KEY (`email`) REFERENCES `siswa` (`email`) ON DELETE CASCADE;

--
-- Constraints for table `siswa`
--
ALTER TABLE `siswa`
  ADD CONSTRAINT `siswa_ibfk_1` FOREIGN KEY (`email`) REFERENCES `user` (`email`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
