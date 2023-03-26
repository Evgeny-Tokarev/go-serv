import {defineStore} from "pinia"
import type {User} from "@/models"
import api from "@/api"

const transformUser = (res: { [name: string]: any }) => {
    return {
        CreatedAt: new Date(res.CreatedAt) || null,
        DeletedAt: new Date(res.DeletedAt) || null,
        CVList: res.Entries || null,
        ID: res.ID,
        UpdatedAt: new Date(res.UpdatedAt) || null,
        username: res.username,
        photo: res.photo
    }
}
export const useUserStore = defineStore({
    id: "userStore",
    state: () => ({
        currentUser: null as User | null
    }),
    actions: {
        async setCurrentUser() {
            try {
                const res = await api.getCurrentUser()
                if (res && res.status === 200 && res.data) {
                    this.currentUser = transformUser(res.data)
                    return true
                } else {
                    this.currentUser = null
                    return false
                }
            } catch (err: unknown) {
                console.log(err)
                this.currentUser = null
            }

        },
        async authenticateUser(email: string, password: string) {
            try {
                const res = await api.login(email, password)
                if (res && res.status === 200 && res.data) {
                    this.currentUser = transformUser(res.data)
                    return true
                } else return false
            } catch (err: unknown) {
                console.log(err)
                return false
            }
        },
        async changeUser(user: { name: string, CVList: [], photo?: File }) {
            try {
                const res = await api.updateUser(this.currentUser?.ID, user.name, user.CVList, user.photo)
            } catch (err: unknown) {
                console.log(err)
                return false

            }
        }
    }
})
