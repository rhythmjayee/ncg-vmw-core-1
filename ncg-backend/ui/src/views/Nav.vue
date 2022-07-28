<template>
    <div>
        <b-navbar toggleable="lg" type="light" variant="light">
            <b-navbar-brand href="#">DIaaS</b-navbar-brand>

            <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

            <b-collapse v-if="isLoggedIn" id="nav-collapse" is-nav>
            <b-navbar-nav>
                <b-nav-item><router-link class="nav-link" to="/">Home</router-link></b-nav-item>
                <b-nav-item><router-link class="nav-link" to="/build">Build</router-link></b-nav-item>
                <b-nav-item><router-link class="nav-link" to="/route">Route</router-link></b-nav-item>
                <b-nav-item>
                    <router-link class="nav-link" to="/functions">My Functions</router-link>
                </b-nav-item>
            </b-navbar-nav>
            </b-collapse>
            <b-navbar-nav v-else class="ml-auto">
                <b-nav-item><router-link class="nav-link" to="/login">Login</router-link></b-nav-item>
                <b-nav-item><router-link class="nav-link" to="/register">Register</router-link></b-nav-item>
            </b-navbar-nav>
            <b-button class="ml-auto" v-if="isLoggedIn" pill sm variant="dark" @click.prevent="logout">LogOut</b-button>
        </b-navbar>
    </div>
</template>

<script>
export default {
  name: 'Nav',
  computed: {
    isLoggedIn: function () {
      return this.$store.getters.isAuthenticated;
    },
  },
  methods: {
    async logout() {
      await this.$store.dispatch('LogOut');
      this.$router.push('/login');
    },
  },
};
</script>

<style>
body {
  margin: 0;
}

#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto,
    Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #3b3b3b;
}

#nav {
  margin-top: 2em;
}

#nav a:hover {
  text-decoration: none;
  outline: none;
}

#nav .router-link-exact-active {
  color: #3b3b3b;
  padding-bottom: 5px;
  border-bottom: 2px solid #3b3b3b;
  outline: none;
}

#nav a {
  font-weight: bold;
  margin-right: 1em;
}
</style>
