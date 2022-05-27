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
import { store } from "./components/store.js";

const getCookie = (name) => {
  // this will read the cookies find a cookie by that name and return a value
  return document.cookie.split("; ").reduce((r, v) => {
    const parts = v.split("=");
    return parts[0] === name ? decodeURIComponent(parts[1]) : r;
  }, "");
};
export default {
  name: "App",
  components: {
    MyHeader,
    MyFooter,
  },
  data() {
    return {
      store,
    };
  },
  beforeMount() {
    //check for a cookie
    let data = getCookie("_site_data");

    if (data !== "") {
      let cookieData = JSON.parse(data);

      //update store
      store.token = cookieData.token.token;
      store.user = {
        id: cookieData.user.id,
        first_name: cookieData.user.first_name,
        last_name: cookieData.user.last_name,
        email: cookieData.user.email,
      };
    }
  },
};
</script>

<style>
</style>