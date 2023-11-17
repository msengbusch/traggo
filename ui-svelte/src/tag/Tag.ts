export class Tag {
    key: string
    color: string
    usages: number

    constructor(key: string, color: string, usages: number) {
        this.key = key
        this.color = color
        this.usages = usages
    }
}