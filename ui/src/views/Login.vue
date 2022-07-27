<template>
    <div class="vue-tempalte">
        <form>
            <h3>Sign In</h3>
            <div class="form-group">
                <label for="email">Username
                <input :value="userData.username"
                @input="userData.username = $event.target.value" id="email"
                type="email" class="form-control form-control-lg" />
                </label>
            </div>
            <div class="form-group">
                <label for="password">Password
                <input v-model.lazy="userData.password" type="password" id="password"
                 class="form-control form-control-lg mb-2" />
                </label>
            </div>
            <button @click.prevent="submit" class="btn btn-dark btn-lg btn-block">SignIn</button>
            <p class="forgot-password text-right mt-2 mb-4">
                <router-link to="/register">Have'nt Registered?</router-link>
            </p>
        </form>
    </div>
</template>
<script>
import { mapActions } from 'vuex';

export default {
  name: 'Login',
  data() {
    return {
      userData: {
        username: '',
        password: '',
      },
      showError: false,
      dataSwitch: true,
      isSubmitted: false,
    };
  },
  methods: {
    ...mapActions(['LogIn']),
    async submit() {
      try {
        await this.LogIn(this.userData);
        this.$router.push('/');
        this.showError = false;
      } catch (error) {
        this.showError = true;
      }
    },
  },
  components: {},
};
</script>
