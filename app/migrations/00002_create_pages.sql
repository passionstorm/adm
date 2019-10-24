-- +goose Up
-- SQL in this section is executed when the migration is applied.
DROP TABLE IF EXISTS `pages`;
CREATE TABLE pages (
    id int(11) NOT NULL AUTO_INCREMENT,
    title VARCHAR(255),
    page_title VARCHAR(255),
    meta_description TEXT,
    content TEXT,
    slug VARCHAR(255),
    layout VARCHAR(255) DEFAULT 'two-col',
    inserted_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_flg TINYINT(1) NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;
ALTER TABLE `pages` ADD UNIQUE INDEX `pages_slug_index` (`slug`);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE pages;