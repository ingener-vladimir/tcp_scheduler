package utils

import (
	"encoding/json"
	"fmt"
)

func JsonUnmarshal(bytes []byte, iface interface{}) error {
	if err := json.Unmarshal(bytes, &iface); err != nil {
		fmt.Println("Ошибка десериализации json: ", err)
		return err
	}
	return nil
}
