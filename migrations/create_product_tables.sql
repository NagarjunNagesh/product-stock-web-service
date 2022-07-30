DROP TABLE IF EXISTS `product` CASCADE;

CREATE TABLE `product` (
  `id`           UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
  `name`         varchar(45)                 NOT NULL CHECK ( first_name <> '' ),
  `stock`        int(11)                              DEFAULT '0',
  `created_at`   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
  `updated_at`   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
) DEFAULT CHARSET=utf8