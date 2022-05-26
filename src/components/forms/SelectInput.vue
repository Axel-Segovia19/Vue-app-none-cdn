<template>
<!-- This is a select form component to allow you to select different items from a dropdown menu-->
  <div class="mb-3">
    <label :for="name" class="form-label">{{ label }}</label>
    <select :id="name" class="form-select" :name="name" :required="required">
      <option
        v-for="option in items"
        :value:="option.value"
        :key="option.value"
      >
      <!-- need :key="option.value" to allow vue to give it a specific value -->
        {{ option.text }}
      </option>
    </select>
  </div>
</template>

<script>
export default {
  name: "SelectInput",
  props: ["items", "name", "required", "label"],
};
</script>