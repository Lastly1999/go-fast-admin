import { message } from 'ant-design-vue'

export type NoticeType = 'info' | 'success' | 'error' | 'warning' | 'loading';

export const alertMsg = (type: NoticeType, msg: string) => {
  return message[type](msg)
}