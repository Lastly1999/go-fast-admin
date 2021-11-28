import interceptor from "./Interceptor"
import {AxiosResponse, Method} from "axios"

const api_url = "/api"

// request typings
export type HttpRequestOptions<T> = {
    method: Method;
    path: string;
    data?: T
}

/**
 * 服务层底层请求
 * @param options
 */
const httpRequest = <T>(options: HttpRequestOptions<T>): Promise<any> => {
    return new Promise((resolve, reject) => {
        interceptor({
            method: options.method,
            url: `${api_url}${options.path}`,
            data: options.data,
        }).then((res: AxiosResponse<any>) => {
            resolve(res)
        }).catch((err: any) => {
            reject(err)
        })
    })
}

export default httpRequest