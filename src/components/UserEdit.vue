<template>
  <div class="containter">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">User</h1>
      </div>
    </div>
  </div>
</template>

<script> // this logic protects this route from non logged in users from  accessing this route/page
import Security from './security.js' //logic to secure route is in here importing it to keep code dry
export default {
  beforeMount() {
    Security.requireToken();
  }
  
}
</script>
