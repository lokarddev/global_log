import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    connection: null,
    all_logs: [],
  },
  mutations: {
    updateLogs(state, logs) {
      state.all_logs = logs
    },
    newLogFromWs(state, log) {
      state.all_logs.push(log)
    }
  },
  actions: {
    createWSConnection( {commit} ) {
      this.connection = new WebSocket("ws://localhost:8080/api/logger-ws-conn")
      this.connection.onmessage = function(message) {
        let j = JSON.parse(message.data)
        commit("newLogFromWs", j)
      }
      console.log(this.connection)
    },
    closeWSConnection () {
      this.connection.close()
    },
  },
  getters: {
    allLogs(state) {
      return state.all_logs
    }
  },
  modules: {
  }
})
