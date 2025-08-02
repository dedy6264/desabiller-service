package helpers

import "desabiller/configs"

func SignIakEncrypt(additional string) (sign string) {

	// sign: md5({username}+{api_key}+{additional})
	key := configs.IakUsername + configs.IakApiKey + additional
	sign = createHash(key)
	return
}
