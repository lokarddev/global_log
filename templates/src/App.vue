<template>
  <div>
    <Header></Header>
    <div class="w-full flex">
      <router-view></router-view>
    </div>
  </div>
</template>
<script>

import Header from "@/components/Header";
import { mapActions } from 'vuex';
export default {
  name: 'App',
  components: {Header},
  methods: mapActions(["createWSConnection", "closeWSConnection"]),
  data() {
    },
  mounted() {
    this.createWSConnection();
  },
  unmounted() {
    this.closeWSConnection();
  }
};
</script>
