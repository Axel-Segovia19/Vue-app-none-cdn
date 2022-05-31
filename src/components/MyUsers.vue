<template>
  <div class="containter">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">All Users</h1>

        <hr />

        <table v-if="this.ready" class="table table-compact table-striped">
          <thead>
            <tr>
              <th>User</th>
              <th>Email</th>
              <th>Active</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="u in this.users" v-bind:key="u.id">
              <td>
                <router-link :to="`/admin/users/${u.id}`"
                  >{{ u.last_name }}, {{ u.first_name }}</router-link
                >
              </td>
              <td>{{ u.email }}</td>

              <td v-if="u.active === 1">
                <span class="badge bg-success">Active</span>
              </td>
              <td v-else>
                <span class="badge bg-danger">Inactive</span>
              </td>

              <td v-if="u.token.id > 0">
                <a href="javascript:void(0);">
                  <span class="badge bg-success" @click="logUserOut(u.id)"
                  >Logged in
                  </span>
                </a>
              </td>
              <td v-else>
                <span class="badge bg-danger">Not logged in</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
// this logic protects this route from non logged in users from  accessing this route/page
import Security from "./security.js";
import notie from "notie";
import { store } from "./store.js";

export default {
  name: 'MyUsers',
  data() {
    // we store the data we get to get a list of all users in the local state
    return {
      users: [],
      ready: false, // will need to do on other pages as well
      store,
    };
  },
  beforeMount() {
    Security.requireToken(); //verifying user is allowed to be here

    fetch(
      process.env.VUE_APP_API_URL + "/admin/users",
      Security.requestOptions("")
    ) // making a request to the back end to get a list of all users
      .then((response) => response.json()) //convert it to the format we need it in
      .then((response) => {
        if (response.error) {
          // this had notie replaced it with emit and changed data to response
          this.$emit("error", response.message);
        } else {
          this.users = response.data.users;
          this.ready = true;
        }
      })
      .catch((error) => {
        this.$emit("error", error);
      });
  },
  methods: {
    logUserOut(id) {
      if (id !== store.user.id) {
        notie.confirm({
          text: "Are you sure you want to log this user out?",
          submitText: "Log Out",
          submitCallback: () => {
            console.log("Would log out user id", id);
            fetch(process.env.VUE_APP_API_URL + "/admin/log-user-out/" + id, Security.requestOptions(""))
            .then((response) => response.json())
            .then((data) => {
              if (data.error) {
                this.$emit('error', data.message);
              } else {
                 this.$emit('success', data.message);
                 this.$emit('forceUpdate');
              }
            })
          },
        });
      } else {
        this.$emit("error", "You can't log yourself out!");
      }
    },
  },
};
</script>
