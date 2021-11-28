export type UserSystem = {
    id: number;
    userName: string;
    userAvatar: string;
    roleIds: string[];
    userSign: string;
    createDate: string;
    activatedAt: string;
    name: string;
    email: string;
    role: UserChilRole[];
    birthday: string;
    age: string;
}

export type UserChilRole = {
    roleId: string;
    roleName: string;
    describe: string;
    status: boolean;
    user: null;
    baseMenu: null;
    createdAt: string;
    updatedAt: string;
}

export type RegisterUserForm = {
    id?: number | null;
    userName: string;
    passWord?: string;
    nikeName: string;
    roleId: string | null;
    roleIds?: string[];
}

export type UserAsRoleParams = {
    id: number;
    roleIds: string[];
}