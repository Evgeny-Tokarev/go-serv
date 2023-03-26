export interface User {
    CreatedAt: Date | null
    DeletedAt: Date | null
    CVList: unknown | null
    ID: number
    UpdatedAt: Date | null
    username: string
    photo?: string
}
