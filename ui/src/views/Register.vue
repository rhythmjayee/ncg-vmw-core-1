<template>
    <div class="vue-tempalte">
        <form>
            <h3>Register</h3>
             <div class="form-group">
                <label for="name">User Name
                <input :value="userData.username"
                @input="userData.username = $event.target.value" id="name"
                type="text" class="form-control form-control-lg" />
                </label>
            </div>
            <div class="form-group">
                <label for="email">Email address
                <input :value="userData.email"
                @input="userData.email = $event.target.value" id="email"
                type="email" class="form-control form-control-lg" />
                </label>
            </div>
            <div class="form-group">
                <label for="password">Password
                <input v-model.lazy="userData.password" type="password" id="password"
                 class="form-control form-control-lg mb-2" />
                </label>
            </div>
            <button @click.prevent="submit" class="btn btn-dark btn-lg btn-block">Submit</button>
            <p class="forgot-password text-right mt-2 mb-4">
                <router-link to="/login">Already a User?</router-link>
            </p>
        </form>
    </div>
</template>
<script>
import { mapActions } from 'vuex';

export default {
  name: 'Register',
  data() {
    return {
      userData: {
        username: '',
        email: '',
        password: '',
      },
      showError: false,
      dataSwitch: true,
      isSubmitted: false,
    };
  },
  methods: {
    ...mapActions(['Register']),
    async submit() {
      console.log('submitted');
      try {
        await this.Register(this.userData);
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
