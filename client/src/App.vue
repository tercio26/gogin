<template>
  <div class="container-fluid">
    <main>
      <div class="row flex-grow-1">
        <div class="col-md-2 d-flex flex-column border-end pb-3">
          <h1>Goredan</h1>

          <template v-if="joined">
            <div class="d-flex justify-content-between">
              <h5>Hello, {{ username }}</h5>
              <a @click.prevent="logout" class="text-primary" href="javascript:;">Logout</a>
            </div>
            <hr/>
            <div class="chat-container mt-auto">
              <div class="chat-content p-2">
                <template v-if="chatContent.length > 0" v-for="msg in chatContent">
                  <div v-if="msg.type === config.MESSAGE_TYPE_ALERT" class="text-center small text-black-50">{{msg.message}}</div>

                  <div v-else-if="msg.type === config.MESSAGE_TYPE_CHAT">
                    <b :class="{'text-danger': msg.username !== username }">{{ msg.username }}:</b>
                    {{ msg.message }}
                  </div>
                </template>
              </div>

              <form @submit.prevent="sendMessage">
                <div class="input-group">
                  <input type="text" class="form-control" placeholder="Type a message..." v-model="message" minlength="0" maxlength="50" required>
                  <button class="input-group-text btn btn-success" type="submit">Send</button>
                </div>

              </form>
            </div>

          </template>

          <div v-else class="my-auto">
            <form class="text-center" @submit.prevent="join">
              <p>You're not joined yet.</p>

              <input type="text" class="form-control form-control-lg mb-3" v-model.trim="username" placeholder="Username" minlength="3" maxlength="32" required>

              <button type="submit" class="btn btn-success w-100">Join</button>
            </form>
          </div>
        </div>
        <div class="col-md-10">
          <nav>

          </nav>
        </div>
      </div>
    </main>
  </div>

</template>

<script>
import config from "./config.js";
import ChatMessage from "./models/ChatMessage.js";

export default {
  name: "App",
  components: {},
  data() {
    return {
      joined: false,
      username: '',

      chatContent: [],
      message: '',
      socket: null
    }
  },
  computed: {
    config() {
      return config
    }
  },

  created() {
    if (localStorage.getItem('username')) {
      this.username = localStorage.getItem('username')
      this.join(false)
    }
  },

  methods: {
    logout() {
      this.socket.close()
      this.joined = false;
      this.username = ''
      localStorage.setItem('username', '')
    },

    initSocket() {
      if (this.socket != null) {
        return
      }
      let me = this;
      this.socket = new WebSocket(config.WS_DOMAIN);

      this.socket.onmessage = function (e) {
        const response = JSON.parse(e.data)
        const msg = new ChatMessage(response.type, response.username, response.message)
        me.chatContent.unshift(msg)
      }

      this.socket.onerror = function (event) {
        console.log("WebSocket Error!");
      }

      return new Promise((resolve) => {
        if (this.socket && this.socket.readyState !== this.socket.OPEN) {
          this.socket.onopen = function (event) {
            console.log("WebSocket connected!");
            resolve()
          }
        } else {
          resolve()
        }
      })
    },

    async join(needAlert = true) {
      await this.initSocket()

      if (needAlert) {
        const msg = new ChatMessage(config.MESSAGE_TYPE_ALERT, this.username, `${this.username} has joined`)
        this.socket.send(msg.stringify())
      }

      this.joined = true;
      localStorage.setItem('username', this.username)
    },

    sendMessage() {
      const msg = new ChatMessage(config.MESSAGE_TYPE_CHAT, this.username, this.message)
      this.socket.send(msg.stringify())
      this.message = ""
    },

    async ping() {
      const response = await fetch(config.EndPoint(config.API.PING));
      const data = await response.json();

      this.message = data.message
    }
  }
}
</script>