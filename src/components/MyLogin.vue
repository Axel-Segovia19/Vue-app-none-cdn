<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <!-- @myevent is an event you will fire off from somewhere else in your code this component will listen for and then perform some action
        can also be written v-on:myevent="" will need a method call submit handler look in script -->
        <form-tag @myevent="submitHandler" name="myForm" event="myevent"> <!-- the event you emit from FormTag is being named here and also binds the event to myevent -->
          <!-- using the TextInput component for email and password that a user will use to login with email and password-->
           <!-- Make sure to register your component in the script tag to give it function with javascript-->
          <text-input
            v-model="email"
            label="Email"
            type="email"
            name="email"
            required="true"
          >
          </text-input>

          <text-input
            v-model="password" 
            label="Password"
            type="password"
            name="password"
            required="true"
          >
          </text-input>

          <hr />
          <input type="submit" class="btn btn-primary" value="Login" />
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>//inside the script tag is where you will enter your javascript
import FormTag from "./forms/FormTag.vue"; // importing the formtag component to be used here in the login screen 
import TextInput from "./forms/TextInput.vue"; //importing the TextInput component to give it some function with Javascript 
export default {
  name: "MyLogin",
  components: { // this registers your components to be used in you html vuejs code
    FormTag,
    TextInput, 
  },
  data() { //this gets the values from text-input for email and password make sure to bind them in vueJS with v-model=""
  // reason why they are empty strings is due to the fact that is what will hold the information to then run it through the backend
    return {
      email: "",
      password: "",
    };
  },
  methods: {// this gives function to the @myevent handle up top for logging in
    submitHandler() {
      console.log("submitHandler called - success!");

      const payload = { // this what the back end is expenting to find 
        email: this.email,
        password: this.password,
      };

      const requestOptions = { // 
        method: "POST",
        body: JSON.stringify(payload), // payload is converted to json 
      };

      fetch("http://localhost:8081/users/login", requestOptions) // that makes the call to the back end and sends the payload json
        .then((response) => response.json()) // after you make the call yuo get your response and convert it to response json
        .then((data) => { // catching the error if data was not received
          if (data.error) {
            console.log("Error:", data.message);
          } else {
            console.log(data);
          }
        });
    },
  },
};
</script>