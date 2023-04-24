SELECT sources.site_id, url, IFNULL(err, 0) as err FROM sources LEFT JOIN (
	SELECT site_id, err, recorded FROM probes ORDER BY recorded DESC LIMIT 1
) AS probe ON sources.site_id=probe.site_id
WHERE sources.active=1 AND
TIMESTAMPDIFF(SECOND, probe.recorded, NOW()) > ?
LIMIT ?
