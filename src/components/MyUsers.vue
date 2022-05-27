<template>
  <div class="containter mx-auto">
    <div class="row">
      <div class="col-auto">
        <h1 class="mt-3">All Users</h1>
      </div>
      <hr />


      <table class="table table-responsive table-striped">
        <thead>
          <tr>
            <th>User</th>
            <th>Email</th>
            <hr />

          </tr>
        </thead>
        <tbody>
          <tr v-for="u in this.users" v-bind:key="u.id">
            <td>
              <router-link :to="`/admin/users/${u.id}`"
                >{{ u.last_name }}, {{ u.first_name }}></router-link
              >
            </td>
            <td>{{ u.email }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
// this logic protects this route from non logged in users from  accessing this route/page
import Security from "./security.js";
import notie from "notie";
export default {
  data() {
    // we store the data we get to get a list of all users in the local state
    return {
      users: [],
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
          notie.alert({
            type: "error",
            text: response.message,
          });
        } else {
          this.users = response.data.users;
        }
      })
      .catch((error) => {
        notie.alert({
          type: "error",
          text: error,
        });
      });
  },
};
</script>
