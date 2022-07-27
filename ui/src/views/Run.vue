<template>
  <div role="group">
    <b-form-group>
      <b-container fluid>
        <b-row class="my-1">
          <b-col sm="2">
            <div>URL:</div>
          </b-col>
          <b-col sm="10">
            <b-form-input id="funcId" v-model="url" placeholder="Enter image id with params if any"></b-form-input>
          </b-col>
        </b-row>
        <b-row class="my-1">
          <b-col sm="2">
            <div for="input-small">Duration:</div>
          </b-col>
          <b-col sm="10">
            <b-form-input id="duration" v-model="duration" placeholder="Enter duration in seconds (eg: 5s)"></b-form-input>
          </b-col>
        </b-row>
        <div>
          <b-button block @click="run" variant="primary">Run</b-button>
        </div>
      </b-container>
    </b-form-group>
    <div id="app" class="container py-2">
        <div class="py-2">
            <div class="progress my-3">
                <div class="progress-bar" ref="elapsedTime" role="progressbar" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100">
                    {{ minutes>0 ? minutes + ' seconds' : '' }}
                </div>
            </div>
        </div>
    </div>
    <div>
      <codemirror v-model="output"
      :options="{ theme: 'base16-dark', lineNumbers: true, mode: 'javascript' }" class="mt-3"/>
    </div>
  </div>
</template>

<script>

export default {
  name: 'playground',
  data() {
    return {
      url: '',
      params: '',
      duration: '',
      output: 'Check your output here...',
      isLoading: false,
      minutes: 0,
      countdown: null,
    };
  },
  methods: {
    run() {
      this.isLoading = true;
      this.output = 'Loading...';
      this.startTimer();
      const body = { duration: this.duration };
      this.axios.post(`${this.$endpoint}/trigger/${this.url}`, body)
        .then((res) => {
          this.output = res.data;
        });
    },
    startTimer() {
      const et = this.$refs.elapsedTime;
      this.minutes = this.duration.substr(0, this.duration.length - 1);
      const seconds = this.minutes; // initial time
      et.style.width = '0%';
      this.countdown = setInterval(() => {
        this.minutes -= 1;
        et.style.width = `${((seconds - this.minutes) * 100) / seconds}%`;
        if (this.minutes === 0) {
          this.stopTimer();
        }
      }, 1000);
    },
    stopTimer() {
      clearInterval(this.countdown);
      this.minutes = 0;
    },
  },
};
</script>

<style>
#div1{
  display: flex;
}
</style>
