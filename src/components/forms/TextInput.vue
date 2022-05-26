<template>
 <!-- This is a reusable component that can be used for text input forms of any kind -->
  <div class="mb-3">
    <label :for="name" class="form-label">{{ label }}</label>
    <input
      :type="type"
      :name="name"
      :placeholder="placeholder"
      :required="required" 
      :min="min"
      :max="max"
      :value="value"
      :autocomplete="name + '-new'"
      class="form-control"
    />
    <!-- by using : in front of the value like :for ="" you are saying it will be assigned the type,name,placeholder etc-->
     <!-- you give it when you are using the component for a specific form. pretty cool and makes it super easy!-->
  </div>
</template>
<script>
export default {
  name: "TextInput",
  props: { // this is defining data type of input that will be on the component  
    name: String,
    type: String,
    label: String,
    placeholder: String,
    required: String,
    min: String,
    max: String,
    value: String,
  },
};
</script>