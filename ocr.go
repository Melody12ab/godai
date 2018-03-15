package godai

import (
    "io/ioutil"
    "encoding/base64"
    "fmt"
    "net/url"
    "net/http"
)

func rec(baseUrl, path string) ([]byte, error) {
    filebytes, _ := ioutil.ReadFile(path)
    sourcestring := base64.StdEncoding.EncodeToString(filebytes)
    token := getToken()
    urlstr := fmt.Sprintf("%s?access_token=%s", baseUrl, token)
    //todo options参数抽出来
    params := url.Values{
        "image": {sourcestring},
    }
    res, err := http.PostForm(urlstr, params)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    //fmt.Println(string(body))
    return body, err
}

func GeneralBasic(path string) ([]byte, error) {
    //通用文字识别
    general_basic := "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
    return rec(general_basic, path)
}

func AccurateBasic(path string) ([]byte, error) {
    //通用文字识别（高精度版）
    accurate_basic := "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
    return rec(accurate_basic, path)
}

func General(path string) ([]byte, error) {
    //通用文字识别（含位置信息版）
    general := "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
    return rec(general, path)
}

func Accurate(path string) ([]byte, error) {
    //通用文字识别（含位置高精度版）
    accurate := "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate"
    return rec(accurate, path)
}

func GeneralEnhanced(path string) ([]byte, error) {
    //通用文字识别（含生僻字版）
    general_enhanced := "https://aip.baidubce.com/rest/2.0/ocr/v1/general_enhanced"
    return rec(general_enhanced, path)
}

func Webimage(path string) ([]byte, error) {
    //网络图片文字识别
    webimage := "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"
    return rec(webimage, path)
}
