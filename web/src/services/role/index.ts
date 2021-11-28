import httpRequest from "@/utils/axios/service";
import { HttpResponse } from "../model/request/public";
import { RoleForm, RolePermissionParam } from "../model/response/role"

/**
 * 获取角色列表
 * @param data
 */
export const getRoles = (): Promise<HttpResponse> => {
    return httpRequest({
        method: 'get',
        path: '/role/role'
    })
}

/**
 * 修改角色信息
 * @param data
 * @returns
 */
export const editRole = (data: RoleForm): Promise<HttpResponse> => {
    return httpRequest({
        method: 'patch',
        path: '/role/role',
        data
    })
}

/**
 * 删除角色信息
 * @param id
 */
export const delRole = (id: number | null) => {
    return httpRequest({
        method: 'delete',
        path: `/role/role/${id}`
    })
}

/**
 * 添加角色
 * @param data
 */
export const appendRole = (data: RoleForm): Promise<HttpResponse> => {
    return httpRequest({
        method: 'put',
        path: `/role/role`,
        data
    })
}

/**
 * 修改角色菜单
 * @param data
 */
export const editPermission = (data: RolePermissionParam): Promise<HttpResponse> => {
    return httpRequest({
        method: 'patch',
        path: `/role/menu`,
        data
    })
}

