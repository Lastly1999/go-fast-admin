import md5 from "js-md5"

// 加密密码
export const createMd5Pass = (pass: string): string => {
    return md5(pass)
}