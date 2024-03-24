export default class ChatMessage {
    constructor(type, username, message) {
        this.type = type;
        this.username = username;
        this.message = message;
    }

    toJSON()     {
        return {
            type: this.type,
            username: this.username,
            message: this.message
        }
    }

    stringify() {
        return JSON.stringify({
            type: this.type,
            username: this.username,
            message: this.message
        })
    }
}