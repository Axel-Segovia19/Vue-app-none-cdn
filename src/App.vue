<template>
  <MyHeader />
  <div>
    <!-- defining where you want your routed components to show up between div tags by calling router-view -->
    <router-view />
  </div>
  <MyFooter />
</template>



<!-- every time you import a component and register it you have to use it! or else it wont compile -->
<script>
import MyHeader from "./components/MyHeader.vue";
import MyFooter from "./components/MyFooter.vue";

export default {
  name: "App",
  components: {
    MyHeader,
    MyFooter,
  },
};
</script>

<style>
</style>