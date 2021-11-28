import httpRequest from "@/utils/axios/service"
import {HttpResponse} from "../model/request/public"
import {ApiLogsParams} from "@/services/model/response/role";

/**
 * 获取系统图标下拉列表
 */
export const getSystemIcons = (): Promise<HttpResponse> => {
    return httpRequest({
        method: "get",
        path: "/sys/icons"
    })
}

/**
 * 系统api日志
 */
export const getSystemLogs = (data: ApiLogsParams): Promise<HttpResponse> => {
    return httpRequest({
        method: "post",
        path: "/logs/logs",
        data
    })
}