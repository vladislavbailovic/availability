USE narfs;
CREATE TABLE IF NOT EXISTS `incidents` (

	incident_id   INT(32) AUTO_INCREMENT NOT NULL,
	site_id       INT(32)                NOT NULL,
	down_probe_id INT(32) NULL,
	up_probe_id   INT(32) NULL,

	PRIMARY KEY(incident_id),
	FOREIGN KEY(site_id)       REFERENCES sources(site_id) ON DELETE CASCADE,
	FOREIGN KEY(down_probe_id) REFERENCES probes(probe_id) ON DELETE CASCADE,
	FOREIGN KEY(up_probe_id)   REFERENCES probes(probe_id) ON DELETE CASCADE
);


