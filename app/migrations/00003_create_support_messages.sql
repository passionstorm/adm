-- +goose Up
-- SQL in this section is executed when the migration is applied.
DROP TABLE IF EXISTS `support_messages`;
CREATE TABLE support_messages (
   id int(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    email VARCHAR(255),
    subject VARCHAR(255),
    content TEXT,
    inserted_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_flg TINYINT(1) NOT NULL DEFAULT 0,

   PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE support_messages;