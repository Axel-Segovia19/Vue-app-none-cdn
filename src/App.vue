<template>
<MyHeader />
<div>
  <router-view/>
</div>
<MyFooter />
  
  
</template>




<script>
import MyHeader from "./components/MyHeader.vue"
import MyFooter from "./components/MyFooter.vue"


export default {
  name: 'App',
  components: {
    MyHeader,
    MyFooter,
  }
}
</script>

<style>

</style>