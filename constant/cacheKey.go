package constant
type CacheKey string

const (
	BLOG_PREFIX = "blog:cache:"
	BLOG_ARTICLE_BY_ID = BLOG_PREFIX+"article:%d"
)