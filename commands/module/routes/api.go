r.GET("/{moduleName}", Index)
r.GET("/{moduleName}", Show)
r.POST("/{moduleName}", Store)
r.PATCH("/{moduleName}/:id", Update)
r.DELETE("/{moduleName}/:id", Delete)
