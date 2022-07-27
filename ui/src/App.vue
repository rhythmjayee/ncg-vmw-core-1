<template>
  <div id="app">
    <Nav />
    <div class="container" style="margin-top:50px;">
      <router-view class="mt-3"/>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Nav from '@/views/Nav.vue';

export default {
  components: {
    Nav,
  },
};
</script>

<style lang="scss">

@import "~bootstrap/scss/functions";
@import "~bootstrap/scss/variables";
@import "~bootstrap/scss/mixins";
@import "assets/_custom.scss";
@import "~bootstrap/scss/bootstrap";
@import '~bootstrap-vue/dist/bootstrap-vue.css';
@import '../node_modules/codemirror/lib/codemirror.css';
@import '../node_modules/codemirror/theme/base16-dark.css';

body {
  margin: 0;
  background-image: url('~@/assets/images/bg-img.webp');
}

#app {
  font-family: BlinkMacSystemFont, 'Segoe UI', Roboto,
    Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #3b3b3b;
}

</style>
