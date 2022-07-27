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
          <p class="my-4"><b>Function Name:</b> {{row.item.Name}}</p>
          <p class="my-4"><b>Function Description:</b> {{row.item.Description}}</p>
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
      items: [],
      fields: [
        { key: 'ID' },
        { key: 'Date' },
        { key: 'Ready', label: 'Image Built' },
        { key: 'Language' },
        { key: 'Actions' },
      ],
    };
  },
  methods: {
    refresh() {
      this.axios.get(`${this.$endpoint}/Functions?users=${this.$store.getters.StateUser}`)
        .then((resp) => { this.items = resp.data; });
    },
    del(ID) {
      this.axios.delete(`${this.$endpoint}/Functions/${ID}?users=${this.$store.getters.StateUser}`)
        .then((resp) => { this.items = resp.data; });
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
