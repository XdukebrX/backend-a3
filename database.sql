SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

--
-- Banco de dados: `a3-clavison`
--
CREATE DATABASE IF NOT EXISTS `a3-clavison` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `a3-clavison`;

-- --------------------------------------------------------

--
-- Estrutura da tabela `commodities`
--

CREATE TABLE IF NOT EXISTS `commodities` (
  `id_products` int(11) NOT NULL,
  `id_raw_materials` int(11) NOT NULL,
  `quantity` int(11) DEFAULT NULL,
  KEY `id_product_idx` (`id_products`),
  KEY `id_raw_material_idx` (`id_raw_materials`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Extraindo dados da tabela `commodities`
--

INSERT INTO `commodities` (`id_products`, `id_raw_materials`, `quantity`) VALUES(26, 10, 20);
INSERT INTO `commodities` (`id_products`, `id_raw_materials`, `quantity`) VALUES(1, 2, 1);
INSERT INTO `commodities` (`id_products`, `id_raw_materials`, `quantity`) VALUES(1, 2, 1);
INSERT INTO `commodities` (`id_products`, `id_raw_materials`, `quantity`) VALUES(11, 34, 1);

-- --------------------------------------------------------

--
-- Estrutura da tabela `products`
--

CREATE TABLE IF NOT EXISTS `products` (
  `id_product` int(11) NOT NULL AUTO_INCREMENT,
  `pname` varchar(45) DEFAULT NULL,
  `pvalue` decimal(7,2) DEFAULT NULL,
  PRIMARY KEY (`id_product`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4;

--
-- Extraindo dados da tabela `products`
--

INSERT INTO `products` (`id_product`, `pname`, `pvalue`) VALUES(33, 'Lápis', '1.00');
INSERT INTO `products` (`id_product`, `pname`, `pvalue`) VALUES(34, 'Caneta de Madeira Azul', '2.00');
INSERT INTO `products` (`id_product`, `pname`, `pvalue`) VALUES(35, 'Casa de Cachorro', '200.00');

-- --------------------------------------------------------

--
-- Estrutura da tabela `raw_materials`
--

CREATE TABLE IF NOT EXISTS `raw_materials` (
  `idraw_material` int(11) NOT NULL AUTO_INCREMENT,
  `rname` varchar(45) DEFAULT NULL,
  `stock` int(11) DEFAULT NULL,
  PRIMARY KEY (`idraw_material`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;

--
-- Extraindo dados da tabela `raw_materials`
--

INSERT INTO `raw_materials` (`idraw_material`, `rname`, `stock`) VALUES(11, 'Tinta Azul', 10);
INSERT INTO `raw_materials` (`idraw_material`, `rname`, `stock`) VALUES(12, 'Madeira para Lapis', 5);
INSERT INTO `raw_materials` (`idraw_material`, `rname`, `stock`) VALUES(13, 'Prego', 50);
INSERT INTO `raw_materials` (`idraw_material`, `rname`, `stock`) VALUES(14, 'Tábua', 100);
INSERT INTO `raw_materials` (`idraw_material`, `rname`, `stock`) VALUES(15, 'Telha', 100);

--
-- Restrições para despejos de tabelas
--

--
-- Limitadores para a tabela `commodities`
--
ALTER TABLE `commodities`
  ADD CONSTRAINT `id_product` FOREIGN KEY (`id_products`) REFERENCES `products` (`id_product`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `id_raw_material` FOREIGN KEY (`id_raw_materials`) REFERENCES `raw_materials` (`idraw_material`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;
