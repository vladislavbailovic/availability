SELECT sources.site_id, url, IFNULL(err, 0)
FROM sources
	LEFT JOIN (
		SELECT site_id, MAX(recorded) AS recorded FROM probes GROUP BY site_id
	) AS p1
	ON sources.site_id=p1.site_id
	LEFT JOIN probes AS p2
	ON p2.site_id=p1.site_id AND p2.recorded=p1.recorded
WHERE active=1
AND (TIMESTAMPDIFF(SECOND, p1.recorded, NOW()) > ? OR p1.recorded IS NULL)
LIMIT ?
