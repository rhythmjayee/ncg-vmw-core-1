<template>
    <div class="home">
    <b-table :fields="fields" :items="items" sort-by="Date" :sort-desc="true">
      <template slot="ID" slot-scope="data">
        <router-link :to="{ name: 'details', params: { id: data.value } }">
          {{ data.value | formatId }}
        </router-link>
      </template>
      <template slot="Date" slot-scope="data">
        {{ data.value | moment('from') }}
      </template>
      <template slot="Ready" slot-scope="data">
        {{ data.value | moment('from') }}
      </template>
      <template #cell(Contributor)="row">
        {{row.item.User}}
      </template>
      <template #cell(Actions)="row">
        <b-button size="sm"  v-b-modal="'modal' + row.index" class="ms-2 btn-success">Details
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
  name: 'home',
  data() {
    return {
      refreshTimer: null,
      items: [],
      fields: [
        { key: 'ID' },
        { key: 'Date' },
        { key: 'Ready', label: 'Image Built' },
        { key: 'Language' },
        { key: 'Actions', label: 'Actions' },
        { key: 'Contributor' },
      ],
    };
  },
  methods: {
    refresh() {
      this.axios.get(`${this.$endpoint}/list`).then((resp) => { this.items = resp.data; });
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
