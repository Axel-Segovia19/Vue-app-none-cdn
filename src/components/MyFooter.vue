<template>
  <footer class="d-flex align-items-center justify-content-center mt-5">
    Copyright &copy; {{ new Date().getFullYear() }}
  </footer>
</template>


<style scoped>
footer {
  background: black;
  color: white;
  min-height: 3em;
  font-size: 0.8em;
}
</style>
<!-- scoped means it is only for this component if you leave that off it can spill over-->