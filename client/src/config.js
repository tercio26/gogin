export default Object.freeze({
    API_DOMAIN: 'http://localhost:8081/api',
    WS_DOMAIN: 'ws://localhost:8081/ws',

    API: {
        HOME: "/",
        PING: "/ping"
    },

    MESSAGE_TYPE_ALERT: 1,
    MESSAGE_TYPE_CHAT: 2,

    EndPoint(api) {
        return this.API_DOMAIN + api
    }
})