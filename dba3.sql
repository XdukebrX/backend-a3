-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema a3-clavison
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema a3-clavison
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `a3-clavison` DEFAULT CHARACTER SET utf8mb4 ;
USE `a3-clavison` ;

-- -----------------------------------------------------
-- Table `a3-clavison`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `a3-clavison`.`products` (
  `id_product` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL DEFAULT NULL,
  `value` DECIMAL(7,2) NULL DEFAULT NULL,
  PRIMARY KEY (`id_product`))
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `a3-clavison`.`raw_materials`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `a3-clavison`.`raw_materials` (
  `idraw_materials` INT(11) NOT NULL,
  `name` VARCHAR(45) NULL DEFAULT NULL,
  `stock` INT(11) NULL DEFAULT NULL,
  PRIMARY KEY (`idraw_materials`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `a3-clavison`.`commodities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `a3-clavison`.`commodities` (
  `id_product` INT(11) NOT NULL,
  `id_raw_material` INT(11) NOT NULL,
  `quantity` INT(11) NULL DEFAULT NULL,
  INDEX `id_product_idx` (`id_product` ASC) VISIBLE,
  INDEX `id_raw_material_idx` (`id_raw_material` ASC) VISIBLE,
  CONSTRAINT `id_product`
    FOREIGN KEY (`id_product`)
    REFERENCES `a3-clavison`.`products` (`id_product`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `id_raw_material`
    FOREIGN KEY (`id_raw_material`)
    REFERENCES `a3-clavison`.`raw_materials` (`idraw_materials`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
