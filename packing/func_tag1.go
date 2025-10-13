//go:build !tag1 && !tag2
// +build !tag1,!tag2

// 上面的是新版
// 下面是旧版

package packing

func init() {
	TAG = append(TAG, "tag1")
}
