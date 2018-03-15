package godai

import (
    "bufio"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

const CLIENT_ID = "***"
const CLIENT_SECRET = "***"

const TOKEN_URL_FORMAT = "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s"

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}

//获取配置信息
func getConfig(path string) (config map[string]string, err error) {
    if checkFileIsExist(path) {
        fmt.Println("已存在配置文件")
        //存在配置文件就读取
        f, err1 := os.Open(path)
        defer f.Close()
        if err1 != nil {
            return nil, err1
        }
        br := bufio.NewReader(f)
        config = make(map[string]string)
        for {
            str, _, eof := br.ReadLine()
            if eof == io.EOF {
                break
            } else {
                config_line := strings.Split(string(str), ":")
                config[config_line[0]] = config_line[1]
            }
        }
        return config, nil
    } else {
        fmt.Println("不存在配置文件")
    }
    return
}

func WriteMaptoFile(configs map[string]string, filePath string) error {
    f, err := os.Create(filePath)
    if err != nil {
        fmt.Printf("create map file error: %v\n", err)
        return err
    }
    defer f.Close()

    w := bufio.NewWriter(f)
    for k, v := range configs {
        lineStr := fmt.Sprintf("%s:%s", k, v)
        fmt.Fprintln(w, lineStr)
    }
    return w.Flush()
}

func getToken() string {

    //检查配置文件中是否存在token
    configFile := "config.txt"
    configs, err := getConfig(configFile)
    checkError(err)
    if expires, ok := configs["expires"]; ok {
        lastTime, _ := strconv.Atoi(configs["time"])
        expireTime, _ := strconv.Atoi(expires)
        if int64(lastTime)+int64(expireTime) > time.Now().Unix() {
            //还没过期
            fmt.Println("还没有过期")
            return configs["token"]
        }
    } else {
        fmt.Println("配置文件token已经过期")
    }

    //否则重新请求
    tokenUrl := fmt.Sprintf(TOKEN_URL_FORMAT, CLIENT_ID, CLIENT_SECRET)
    resp, err := http.Get(tokenUrl)
    checkError(err)
    data, err := ioutil.ReadAll(resp.Body)
    checkError(err)
    var tokens interface{}
    json.Unmarshal(data, &tokens)
    result := tokens.(map[string]interface{})
    if errStr, ok := result["error"].(string); ok {
        log.Fatal(errors.New(result["error_description"].(string) + errStr))
    } else {
        //写入配置文件
        token := result["access_token"].(string)
        configs := make(map[string]string)
        configs["expires"] = strconv.Itoa(int(result["expires_in"].(float64)))
        configs["time"] = strconv.Itoa(int(time.Now().Unix()))
        configs["token"] = token
        WriteMaptoFile(configs, configFile)
        return token
    }
    return ""
}

