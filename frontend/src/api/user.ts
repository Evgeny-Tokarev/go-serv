import type {User} from '@/models'
import axios from './axios'

export default {
    async getCurrentUser(): Promise<{ status: number, data: User }> {
        try {
            const resp = await axios.get('user', {
                withCredentials: true
            })
            return {
                data: resp.data,
                status: resp.status
            }
        } catch (err: any) {
            return err.response.status
        }
    },
    async login(name: string, pass: string) {
        try {
            const resp = await axios.post('auth/login', {
                withCredentials: true,
                username: name,
                password: pass
            })
            return {
                data: resp.data,
                status: resp.status
            }
        } catch (err: any) {
            return err.response.status
        }
    },
    async updateUser(id: number, name: string, photo?: File) {
        if (id === undefined) return new Error("no id!")
        if (photo) {
            await this.updateUserImage(id, photo)
        }
        try {
            const resp = await axios.post(`user/${id}`, {
                withCredentials: true,
                username: name
            })
            return {
                data: resp.data,
                status: resp.status
            }
        } catch (err: any) {
            return err.response.status
        }
    },
    async updateUserImage(id: number, img: File) {
        try {
            console.log(id, img)
            // const res = await api.attachContentPhoto(img, id)
            // if (res && res.status === 200 && res.data) {
            // const contentItemToUpdateIndex = this.contentData.findIndex(item => item.id === id)
            // if (contentItemToUpdateIndex !== -1) {
            //     this.contentData[contentItemToUpdateIndex] = res.data
            // }
            return
            // } else return false
        } catch (err) {
            console.log(err)
            return false
        }
    }
}
