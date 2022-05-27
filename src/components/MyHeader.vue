<template>
  <!-- this is navbar code from bootstrap -->
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Navbar</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0" > 
          <li class="nav-item">
            <!-- This is a router link to the MyBody page as the home page -->
            <router-link class="nav-link active" aria-current="page" to="/"
              >Home</router-link
            >
          </li>

          <li class="nav-item">
            <router-link class="nav-link active" to="/books">Books</router-link>
          </li> 

          <li v-if="store.token !== ''" class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="navBarDropDown" role="button" data-bs-toggle="dropdown"
            aria-expanded="false">
            Admin
            </a>
            <ul class="dropdown-menu" aria-labelledby="navBarDropDown">
              <li>
                <router-link class="dropdown-item" to="/admin/users">Manage Users </router-link>
              </li>
              <li>
                <router-link class="dropdown-item" to="/admin/users/0">Add Users </router-link>
              </li>
              <li>
                <router-link class="dropdown-item" to="/admin/books">Manage Books </router-link>
              </li>
              <li>
                <!-- this link here for :to"" uses different syntax but works the same as above! -->
                <router-link class="dropdown-item" :to="{name:'BookEdit', params: {bookId: 0}}">Add Book</router-link>
              </li>
            </ul>
          </li>

          <li class="nav-item">
            <!-- This is a router link to the MyLogin page as the home page -->
            <router-link v-if="store.token == ''" class="nav-link" to="/login"
              >Login</router-link
            >
            <a
              href="javascript:void(0);"
              v-else
              class="nav-link"
              @click="logout"
              >logout</a
            >
            <!-- the router-link tag can be placed pretty much anywhere on the application to use them! -->
          </li>
        </ul>

        <span class="navbar-text">
          {{store.user.first_name ?? ''}} <!-- when logged in you will see your name in navbar and when youre not you wont! -->
        </span>


      </div>
    </div>
  </nav>
</template>


<script>
import { store } from "./store.js";
import router from "./../router/index.js";
import Security from './security.js' //this is using authentication logic from a different page to allow us to clean up code look at fetch to see how Security it being used
export default {
  data() {
    return {
      store,
    };
  },

  methods: {
    logout() {
      //this method redirects you to the login page logging you out but you have to also let the backend know you are logged out
      const payload = {
        token: store.token,
      };
      fetch(process.env.VUE_APP_API_URL + "/users/logout", Security.requestOptions(payload)) // process.env is the link to the localhost better subsitute and better way to change link when going into production
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            console.log(response.message);
          } else {
            store.token = "";
            store.user = {};

            document.cookie = '_site_data=; Path=/; ' +
            'SameSite=Strict; Secure; ' +
            'Expires=Thu, 01 Jan 1970 00:00:01 GMT;'

            router.push("/login");
          }
        });
    },
  },
};
</script>