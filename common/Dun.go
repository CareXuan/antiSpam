package common

import (
	"antispam/base"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func genSignature(params url.Values, ns base.Dun) string {
	var paramStr string
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		paramStr += key + params[key][0]
	}
	paramStr += ns.SecretKey
	md5Reader := md5.New()
	md5Reader.Write([]byte(paramStr))
	return hex.EncodeToString(md5Reader.Sum(nil))
}

func BaseCheck(params url.Values, apiUrl string) (string, error) {
	ns := base.Conf.Dun
	params["secretId"] = []string{ns.SecretId}
	params["timestamp"] = []string{strconv.FormatInt(time.Now().UnixNano()/1000000, 10)}
	params["nonce"] = []string{strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000000), 10)}
	params["signature"] = []string{genSignature(params, ns)}

	resp, err := http.Post(apiUrl, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents), nil
}
