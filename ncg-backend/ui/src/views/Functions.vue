<template>
  <div class="home">
    <b-table :fields="fields" :items="items" sort-by="Date" :sort-desc="true">
      <template slot="ID" slot-scope="data">
        <router-link :to="{ name: 'details', params: { id: data.value } }">
          {{ data.value | formatId }}
        </router-link>
      </template>
      <template slot="Ready" slot-scope="data">
        <span :class="{ 'text-success': data.value, 'text-danger': !data.value }">&#x25a0;</span>
      </template>
      <template slot="Date" slot-scope="data">
        {{ data.value | moment('from') }}
      </template>
      <template #cell(Actions)="row">
        <b-button size="sm" @click="del(row.item.ID)" class="me-2 btn-danger">Delete
        </b-button>
        <b-button size="sm" v-b-modal="'modal' + row.index" class="ms-2 btn-success">Details
        </b-button>
        <b-modal :id="'modal' + row.index" centered title="Details" hide-footer>
          <p class="my-4"><b>Function Name:</b> {{row.item.FunctionName}}</p>
          <p class="my-4"><b>Function Description:</b> {{row.item.FunctionDescription}}</p>
        </b-modal>
      </template>
    </b-table>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: 'functions',
  data() {
    return {
      refreshTimer: null,
      items: [{
        Ready: true,
        ID: 'a05d89103d13f9464484c0f1f09538231772f0c9ae0bcaf6105c3e9ee22a8bd6',
        Date: '2022-07-18T07:01:45.026043897Z',
        Language: 'java',
        FunctionName: 'Add two Numbers',
        FunctionDescription: 'Take 2 numbers and performs sum',
      },
      {
        Ready: true,
        ID: 'a05d89103d13f9464484c0f1f09538231772f0c9ae0bcaf6105c3e9ff22a8bd6',
        Date: '2022-07-19T07:01:45.026043897Z',
        Language: 'cpp',
        FunctionName: 'My function 2',
        FunctionDescription: 'This is a CPP function',
      },
      ],
      fields: [
        { key: 'Ready', label: '' },
        { key: 'ID' },
        { key: 'Date' },
        { key: 'Language' },
        { key: 'Actions' },
      ],
    };
  },
  methods: {
    refresh() {
      // this.axios.get(`${this.$endpoint}/list`).then((resp) => { this.items = resp.data; });
    },
    del(ID) {
      console.log('deleted', ID);
    //   this.axios.delete(`${this.$endpoint}/function/${ID}`).then((resp) => { this.items = resp.data; });
    },
  },
  created() {
    this.refresh();
    this.refreshTimer = window.setInterval(this.refresh, 10000);
  },
  beforeDestroy() {
    window.clearInterval(this.refreshTimer);
  },
};
</script>
