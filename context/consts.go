package context

const (
	// AreaNone ...
	AreaNone = 1

	// AreaIPFilter ...
	AreaIPFilter = 2

	// AreaDomain ...
	AreaDomain = 4

	// AreaSubDomain ...
	AreaSubDomain = 8

	// AreaRedirect ...
	AreaRedirect = 16

	// AreaCache ...
	AreaCache = 32

	// AreaSettings ...
	AreaSettings = 64

	// AreaUser ...
	AreaUser = 128

	// AreaStatistics ...
	AreaStatistics = 256

	// AreaErrorPage ...
	AreaErrorPage = 512

	// AreaSsl ...
	AreaSsl = 1024

	// AreaAll ...
	AreaAll = 0xFFFFFFFFFFFFFFFF
)

const (
	// ContextUp ...
	ContextUp = ".."

	// ContextSelf ...
	ContextSelf = "."

	// ContextIPFilter ...
	ContextIPFilter = "ipFilter"

	// ContextCacheSettings ...
	ContextCacheSettings = "cacheSettings"

	// ContextRedirect ...
	ContextRedirect = "redirects"

	// ContextSettings ...
	ContextSettings = "settings"

	// ContextStatistics ...
	ContextStatistics = "statistics"

	// ContextErrorPages ...
	ContextErrorPages = "errorPages"

	// ContextSsl ...
	ContextSsl = "sslCert"
)
