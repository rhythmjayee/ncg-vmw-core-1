<template>
  <div>
    <codemirror v-model="code"
      :options="{ theme: 'base16-dark', lineNumbers: true, mode: 'javascript' }" class="mb-3"/>
    <b-form>
      <b-form-group id="input-group-1" label="Function Name" label-for="input-1">
        <b-form-input
          id="input-1"
          v-model="form.functionName"
          type="text"
          placeholder="Enter Function Name"
          required
        ></b-form-input>
      </b-form-group>
       <b-form-group id="input-group-2" label="Function Description:"
        label-for="input-2">
        <b-form-textarea
          id="input-2"
          v-model="form.Description"
          placeholder="Enter Description"
          required
        ></b-form-textarea>
      </b-form-group>
      <b-dropdown class="mt-2" text="Deploy" variant="primary">
        <b-dropdown-item @click="deploy('go')">as Go</b-dropdown-item>
        <b-dropdown-item @click="deploy('swift')">as Swift</b-dropdown-item>
        <b-dropdown-item @click="deploy('js')">as NodeJS</b-dropdown-item>
        <b-dropdown-item @click="deploy('c')">as C</b-dropdown-item>
        <b-dropdown-item @click="deploy('cpp')">as Cpp</b-dropdown-item>
    </b-dropdown>
    </b-form>
  </div>
</template>

<script>
export default {
  name: 'build',
  data() {
    return {
      code: '// You can use JavaScript, Go, Swift, C and C++.',
      form: {
        functionName: '',
        Description: '',
      },
    };
  },
  methods: {
    deploy(lang) {
      const submitFunction = {
        code: this.code,
        lang,
        functionName: this.form.functionName,
        description: this.form.Description,
        user: 'Dummyuser',
      };
      console.log(submitFunction);
      // POST API call
      this.axios.post(`${this.$endpoint}/deploy?lang=${lang}`, submitFunction)
        .then(() => {
          this.$router.push('/');
        });
    },
  },
};
</script>
