USE narfs;
CREATE TABLE IF NOT EXISTS `sources` (

	site_id INT(32)      NOT NULL,
	url     VARCHAR(255) NOT NULL,
	created DATETIME     NOT NULL,
	changed DATETIME     NOT NULL,
	active  INT(1)       NOT NULL,

	PRIMARY KEY(site_id)
);
