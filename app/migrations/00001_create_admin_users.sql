-- +goose Up
-- SQL in this section is executed when the migration is applied.
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE admin_users (
    id int(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR (255),
    email VARCHAR (255),
    password_hash CHARACTER(255),
    inserted_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_flg TINYINT(1) NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

ALTER TABLE `admin_users` ADD UNIQUE INDEX `admin_users_unique_mail` (`email`);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE admin_users;