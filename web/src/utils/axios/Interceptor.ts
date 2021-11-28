import axios, { AxiosRequestConfig, AxiosResponse } from "axios"
import { alertMsg } from '../antd/antd'
import router from "@/router"
import type { HttpResponse } from "@/services/model/request/public"

// axios instance
const axiosInstance = axios.create({
    timeout: 5000
})

/**
 * 请求拦截
 * @date 2021年8月9日19:44:07
 */
axiosInstance.interceptors.request.use((config: AxiosRequestConfig): AxiosRequestConfig => {
    config.headers['authorization'] = localStorage.getItem("auth-token")
    return config
})

/**
 * 响应拦截
 * @date 2021年8月9日19:44:12
 */
axiosInstance.interceptors.response.use((response: AxiosResponse): AxiosResponse => {
    serveResponseErrHandler(response)
    requestHandler(response)
    return response.data
}, (err: any) => {
    if (err.response) {
        // 错误处理
        errorsHandler(err.response.data)
    } else {
        alertMsg("error", "服务器异常：请检查服务器!")
    }

    // 异常抛出
    return Promise.reject(err)
})

function serveResponseErrHandler(res: AxiosResponse<HttpResponse>) {
    const statusCode: number = res.data.code
    const errorMsg: string = res.data.msg
    if (statusCode !== 200) alertMsg("error", errorMsg)
}

function requestHandler(response: AxiosResponse<HttpResponse>) {
    let httpStr = `request-url:${JSON.stringify(response.config.url)} method:${response.config.method}`
    console.log(`%c[fast-http-info]`, "color: #fff; font-weight: bold;background-color:green;padding:5px", httpStr)
}

/**
 * server error handler
 * @author lastly
 * @param data
 */
function errorsHandler(data: any) {
    let errorMsg = "服务器异常"
    switch (data.code) {
        case 20001:
            errorMsg = data.data
            router.push('/login').then(r => r)
        case 20002:
            errorMsg = data.data
            router.push('/login').then(r => r)
        case 500:
            errorMsg = data.data
        case 401:
            errorMsg = data.msg
        default:
            errorMsg = data.msg
    }
    alertMsg("error", errorMsg)
}

export default axiosInstance