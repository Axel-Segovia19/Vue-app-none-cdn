<template>
  <!-- this  -->
  <form
    @submit.prevent="submit"
    :ref="name"
    :event="event"
    autocomplete="off"
    :method="method"
    :action="action"
    class="needs-validation"
    novalidate
  >
    <slot></slot>
  </form>
</template>

<script>
export default {
  name: "FormTag",
  props: ["method", "action", "name", "event"],
  methods: {
    submit() {
      let myForm = this.$refs[this.$props.name]; //this is looking for the form the name is referenced to example $refs in form and $props in export default

      if (myForm.checkValidity()) {
        //
        console.log("my event name", this.$props["event"]);
        console.log("Name", this.$props.name);
        this.$emit(this.$props["event"]); //this will emit/send out the event happening here to be captured by the submitHandler in MyLogin
      }
      myForm.classList.add("was-validated"); //standard bootstrap clientside validation
    },
  },
};
</script>