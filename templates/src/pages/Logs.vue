<template>
  <v-simple-table>
    <template v-slot:default>
      <thead>
      <tr>
        <th class="text-left">
          Id
        </th>
        <th class="text-left">
          Time
        </th>
        <th class="text-left">
          LogLevel
        </th>
        <th class="text-left">
          Payload
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in allLogs" :key="item">
        <td>{{ item.id }}</td>
        <td>{{ item.created_at }}</td>
        <td>{{ item.log_level_id }}</td>
        <td>{{ item.payload }}</td>
      </tr>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script>

import axios from "axios";
import store from "@/store";
import { mapGetters } from "vuex";

export default {
  name: "Logs",
  data() {
    return {
    }
  },
  mounted() {
    this.getAllLogs();
    },
  methods :{
    getAllLogs() {
      axios.get("http://localhost:8080/api/get-logs").then(r => {
        store.commit("updateLogs", r.data)
    })}
  },
  computed: mapGetters(["allLogs"]),
}
</script>

<style scoped>

</style>
