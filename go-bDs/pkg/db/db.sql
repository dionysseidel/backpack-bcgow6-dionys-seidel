CREATE DATABASE IF NOT EXISTS `storage` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci; 
USE `storage`;

CREATE TABLE `products` ( 
    `id`INT(11) NOT NULL, 
    `name` VARCHAR(60) NOT NULL, 
    `type` VARCHAR(60) NOT NULL, 
    `count` INT(11) NOT NULL, 
    `price` FLOAT NOT NULL 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4; 

ALTER TABLE `products` 
  ADD PRIMARY KEY (`id`);

ALTER TABLE `products` 
  MODIFY `id` INT(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4; 
COMMIT;

CREATE TABLE `warehouses` (
  `id` int(11) NOT NULL,
  `name` varchar(40) NOT NULL,
  `address` varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Volcado de datos para la tabla `warehouses`
INSERT INTO `warehouses` (`id`, `name`, `address`) VALUES
(1, 'Main Warehouse', '221b Baker Street');

ALTER TABLE `warehouses`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `warehouses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

ALTER TABLE `products` ADD `id_warehouse` INT NOT NULL AFTER `price`;

UPDATE `products` SET `id_warehouse` = '1'
WHERE id > 0;