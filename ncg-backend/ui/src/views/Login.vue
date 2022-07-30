<template>
    <div class="vue-tempalte">
        <form>
            <h3>Sign In</h3>
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
            <button @click.prevent="submit" class="btn btn-dark btn-lg btn-block">SignIn</button>
            <p class="forgot-password text-right mt-2 mb-4">
                <router-link to="/register">Have'nt Register?</router-link>
            </p>
        </form>
    </div>
</template>
<script>

export default {
  name: 'Login',
  data() {
    return {
      userData: {
        email: '',
        password: '',
      },
      showError: false,
      dataSwitch: true,
      isSubmitted: false,
    };
  },
};
</script>
