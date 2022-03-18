<template>
  <v-card color="basil" class="text-center justify-center">
    <v-card-title primary-title class="justify-center py-6 text-center" >
      <h1 class="basil--text text-center justify-center">
        SF LOG
      </h1>
    </v-card-title>
    <v-tabs v-model="tab" background-color="basil" grow>
      <v-tab v-for="item in list" :key="item" v-bind:to="item.to">
        {{ item.title}}
      </v-tab>
    </v-tabs>
  </v-card>
</template>

<script>
export default {
  name: "Header",
  data() {
    return {
      tab: null,
      list: [
        {to: "/", title: "Dashboard"},
        {to: "/logs", title: "Logs"},
        {to: "/analytics", title: "Analytics"},
      ],
    }
  },
}
</script>

<style scoped>
  .basil {
    background-color: #FFFBE6 !important;
  }
  .basil--text {
    color: #356859 !important;
  }
</style>
