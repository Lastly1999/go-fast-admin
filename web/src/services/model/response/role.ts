import {RoleItem} from "@/store/modules/system"

export type LoginForm = {
    userName: string;
    passWord: string;
    codeAuth: string;
    code: string;
}

export type UserInfo = {
    id: number;
    nikeName: string;
    role: RoleItem[];
    roleId: string;
    userAvatar: string;
    userName: string;
}

export interface RoleForm {
    roleId: number | null;
    roleName: string;
    describe: string;
    createDate: string;
    state?: number;
}

export interface RolePermissionParam {
    roleId: number | null;
    permissionId: number[];
}

export type MenuInfo = {
    menuName?: string;
    menuIcon?: string | null;
    menuPath?: string;
    menuParentId?: number;
    menuParentName?: string;
}

export type ApiLogsParams = {
    timer: string[];
    page: number;
    pageSize: number;
}