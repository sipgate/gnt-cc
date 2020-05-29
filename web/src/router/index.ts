import Vue from "vue";
import VueRouter from "vue-router";
import InstancesListView from "@/views/InstancesListView.vue";
import StatisticsView from "@/views/StatisticsView.vue";
import NodesView from "@/views/NodesView.vue";
import DashboardView from "@/views/DashboardView.vue";
import JobsView from "@/views/JobsView.vue";
import PageNames from "@/data/enum/PageNames";
import LoginView from "@/views/LoginView.vue";
import Params from "@/data/enum/Params";
import InstancesDetailView from "@/views/InstanceDetailView.vue";

Vue.use(VueRouter);

export const REDIRECT_TO_QUERY_KEY = "redirect-to";

const routes = [
  {
    path: "/login",
    name: PageNames.Login,
    component: LoginView
  },
  {
    path: `/:${Params.Cluster}?`,
    component: DashboardView,
    children: [
      {
        path: "",
        name: PageNames.Statistics,
        component: StatisticsView
      },
      {
        path: "instances",
        name: PageNames.InstancesList,
        component: InstancesListView
      },
      {
        path: `instances/:${Params.InstanceName}`,
        name: PageNames.InstancesDetail,
        component: InstancesDetailView
      },
      {
        path: "nodes",
        name: PageNames.Nodes,
        component: NodesView
      },
      {
        path: "jobs",
        name: PageNames.Jobs,
        component: JobsView
      }
    ]
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
