package httpclient

import "net/http"

// http提供的GET，POST方法没有办法自定义header、body等信息
// 这里提供方便生成request的方法

// 构建统一的client请求 调用该函数可以直接获取结构体 或者是error

// 提供一个结构体转map的函数 针对url中请求参数

// 此处需要将body不close的情况补充


func send(req *http.Request, )