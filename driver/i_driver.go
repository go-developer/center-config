// Package driver...
//
// Description : 定义驱动的接口约束
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-23 11:59 下午
package driver

// IDriver 驱动的接口约束
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:00 上午 2021/1/24
type IDriver interface {
	Set(namespace string, key string, value interface{}, expire int64)
	Delete(namespace string, key string)
	Expire(namespace string, key string, expire int64)
}
