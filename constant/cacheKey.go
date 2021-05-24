package constant
type CacheKey string

const (
	BLOG_PREFIX = "blog:cache:"
	BLOG_ARTICLE_BY_ID = BLOG_PREFIX+"article:%d"
	BLOG_ARTICLE_BY_TYPE = BLOG_PREFIX+"article:typeId:%d"
	BLOG_ARTICLE_BY_TAG = BLOG_PREFIX+"article:tagId:%d"
	BLOG_ARTICLE_BY_YEAR = BLOG_PREFIX+"article:year:%d"

)