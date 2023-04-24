USE narfs;
CREATE TABLE IF NOT EXISTS `probes` (

	probe_id      INT(32)     AUTO_INCREMENT NOT NULL,
	site_id       INT(32)     NOT NULL,
	recorded      DATETIME    NOT NULL,
	response_time INT(64)     NOT NULL,
	err           INT(8)      NOT NULL,
	msg           VARCHAR(32) NOT NULL,

	PRIMARY KEY(probe_id),
	FOREIGN KEY(site_id) REFERENCES sources(site_id) ON DELETE CASCADE
);

