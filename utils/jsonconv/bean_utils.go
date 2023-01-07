package jsonconv

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

const (
	Camel = 0 //驼峰
	Case  = 1 //下划线
)

/**
 * @Description: 利用gob进行深拷贝
 */
func DeepCopyByGob(src, dst interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(dst)
}

/**
 * @Description: 利用json进行深拷贝    obj,&objTo
 */
func DeepCopyByJson(src, dst any) error {
	if tmp, err := json.Marshal(&src); err != nil {
		return err
	} else {
		err = json.Unmarshal(tmp, dst)
		return err
	}
}
