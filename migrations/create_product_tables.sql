DROP TABLE IF EXISTS `product` CASCADE;

CREATE TABLE `product` (
  `id`           int(11)                     NOT NULL AUTO_INCREMENT,
  `name`         varchar(45)                 NOT NULL CHECK ( first_name <> '' ),
  `stock`        int(11)                              DEFAULT '0',
  `created_at`   TIMESTAMP                            DEFAULT CURRENT_TIMESTAMP ,
  `updated_at`   DATETIME                             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;