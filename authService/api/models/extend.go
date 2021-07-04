package models

// GetByOptionCache 功能选项模式获取
func (obj *_AuthInfoMgr) GetByOptionCache(opts ...Option) (result AuthInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	id, ok := options.query[AuthInfoColumns.ID]
	if ok {
		AuthInfoFunc{}.GetOneById(obj.ctx, id.(int64))
		return
	}
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where(options.query).Find(&result).Error

	return
}
