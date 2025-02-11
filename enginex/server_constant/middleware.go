package server_constant

const (
	SecureMiddleware               = "SECURE"
	CSRFMiddleware                 = "CSRF"
	CSRFIncludeGETMethodMiddleware = "CSRF_GET"
	CORSMiddleware                 = "CORS"
	DBContextAppenderMiddleware    = "DB_CONTEXT_APPENDER"
	NoCacheMiddleware              = "NOCACHE"
	UUIDSessionGeneratorMiddleware = "UUID_SESSION_GENERATOR"
	CacheStoreAppenderMiddleware   = "CACHE_STORE_APPENDER"
	DefaultCSRFGetMethodTokenLookup = "query:_ctoken_"
	DynamoContextAppenderMiddleware = "DYNAMO_CONTEXT_APPENDER_MIDDLEWARE"
	SqsContextAppenderMiddleware = "SQS_CONTEXT_APPENDER_MIDDLEWARE"
)
