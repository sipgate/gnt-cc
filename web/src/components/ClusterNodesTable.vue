<template>
  <section>
    <b-table
      :data="data"
      :columns="columns"
      :paginated="true"
      :pagination-simple="false"
      :pagination-position="paginationPosition"
      :default-sort-direction="defaultSortDirection"
      :sort-icon-size="sortIconSize"
      :per-page="20"
    >
    </b-table>
  </section>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";

interface Memory {
  memoryFree: number;
}

function sortFreeMemoryBySize(a: Memory, b: Memory, isAsc: boolean) {
  if (isAsc) return a.memoryFree - b.memoryFree;
  else return b.memoryFree - a.memoryFree;
}

@Component({
  name: "ClusterNodesTable"
})
export default class ClusterNodesTable extends Vue {
  paginationPosition = "bottom";
  defaultSortDirection = "asc";
  sortIconSize = "small";

  data = [
    {
      name: "abcd.srv.efgh.net",
      primaryInstance: 11,
      secondaryInstance: 12,
      cpus: 64,
      memoryFree: 14213542453,
      memoryTotal: 34313542453,
      diskFree: 1231354245354,
      diskTotal: 2231354245354,
      memoryFreeHuman: "14.2 GB",
      memoryTotalHuman: "32 GB",
      diskFreeHuman: "1.2 TB",
      diskTotalHuman: "2 TB",
      group: "newHardware"
    },
    {
      name: "hijk.srv.efgh.net",
      primaryInstance: 24,
      secondaryInstance: 14,
      cpus: 64,
      memoryFree: 9213542453,
      memoryTotal: 68627084906,
      diskFree: 631354245354,
      diskTotal: 2231354245354,
      memoryFreeHuman: "6.8 GB",
      memoryTotalHuman: "64 GB",
      diskFreeHuman: "522 GB",
      diskTotalHuman: "2 TB",
      group: "newHardware"
    },
    {
      name: "lmno.srv.dev.net",
      primaryInstance: 2,
      secondaryInstance: 23,
      cpus: 24,
      memoryFree: 58627084906,
      memoryTotal: 68627084906,
      diskFree: 1115677122677,
      diskTotal: 2231354245354,
      memoryFreeHuman: "28.4 GB",
      memoryTotalHuman: "32 GB",
      diskFreeHuman: "945 GB",
      diskTotalHuman: "2 TB",
      group: "oldHardware"
    }
  ];

  columns = [
    {
      field: "name",
      label: "Name",
      searchable: true,
      sortable: true
    },
    {
      field: "primaryInstance",
      label: "Primary Instance",
      numeric: true,
      sortable: true
    },
    {
      field: "secondaryInstance",
      label: "Secondary Instance",
      numeric: true,
      sortable: true
    },
    {
      field: "cpus",
      label: "CPUs",
      numeric: true,
      sortable: true
    },
    {
      field: "memoryFreeHuman",
      label: "Memory Free",
      sortable: true,
      customSort: sortFreeMemoryBySize
    },
    {
      field: "memoryTotalHuman",
      label: "Memory Total",
      sortable: true
    },
    {
      field: "diskFreeHuman",
      label: "Disk Free",
      sortable: true
    },
    {
      field: "diskTotalHuman",
      label: "Disk Total",
      sortable: true
    },
    {
      field: "group",
      label: "Group",
      sortable: true
    }
  ];
}
</script>

<style scoped></style>
