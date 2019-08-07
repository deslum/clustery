package consts


const QueryTemplate = `
	select toFloat64(ContentID), toFloat64(SerieID) from history where PartnerID = 10 and SiteID = 1 limit 1000;
`
