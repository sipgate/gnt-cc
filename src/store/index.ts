import Vue from "vue";
import Vuex from "vuex";
import Api from "@/store/api";
import GntInstance from "@/model/GntInstance";
import GntNode from "@/model/GntNode";
import GntCluster from "@/model/GntCluster";

Vue.use(Vuex);

export interface StoreState {
  clusters: GntCluster[];
  nodes: Record<string, GntNode[]>;
  instances: Record<string, GntInstance[]>;
}

const initialState: StoreState = {
  clusters: [],
  nodes: {},
  instances: {},
};

export const Actions = {
  LoadClusters: "loadClusters",
  LoadNodes: "loadNodes",
  LoadInstances: "loadInstances",
  LoadInstance: "loadInstance",
  RemoveToken: "removeToken",
};

export const Mutations = {
  SetClusters: "setClusters",
  SetNodes: "setNodes",
  SetInstances: "setInstances",
  AddInstance: "addInstance",
  UnsetToken: "unsetToken",
};

export default new Vuex.Store({
  state: initialState,
  mutations: {
    [Mutations.SetClusters](state, clusters: GntCluster[]) {
      state.clusters = clusters;
    },
    [Mutations.SetNodes](state, { cluster, nodes }: { cluster: string; nodes: GntNode[] }) {
      state.nodes = {
        ...state.nodes,
        [cluster]: nodes,
      };
    },
    [Mutations.SetInstances](
      state,
      { cluster, instances }: { cluster: string; instances: GntInstance[] }
    ) {
      state.instances = {
        ...state.instances,
        [cluster]: instances,
      };
    },
    [Mutations.AddInstance](
      state,
      { cluster, instance }: { cluster: string; instance: GntInstance }
    ) {
      const instances = state.instances[cluster] || [];

      const oldIndex = instances.findIndex((el) => el.name === instance.name);

      if (oldIndex > -1) {
        instances[oldIndex] = instance;
      } else {
        instances.push(instance);
      }

      state.instances = {
        ...state.instances,
        [cluster]: instances,
      };
    },
    setToken(state, token: string) {
      localStorage.setItem(Api.tokenStorageKey, token);
    },
    [Mutations.UnsetToken](state, token: string) {
      localStorage.removeItem(Api.tokenStorageKey);
    },
  },
  actions: {
    async saveToken({ commit }, token: string) {
      commit("setToken", token);
    },
    async [Actions.RemoveToken]({ commit }, token: string) {
      commit(Mutations.UnsetToken, token);
    },
    async [Actions.LoadClusters]({ commit }) {
      const response = await Api.get("clusters");
      commit(Mutations.SetClusters, response.clusters);
      return response.clusters;
    },
    async [Actions.LoadNodes]({ commit }, { cluster }) {
      const response = await Api.get(`clusters/${cluster}/nodes`);
      commit(Mutations.SetNodes, {
        cluster,
        nodes: response.nodes,
      });
    },
    async [Actions.LoadInstances]({ commit }, { cluster }) {
      const response = await Api.get(`clusters/${cluster}/instances`);
      commit(Mutations.SetInstances, {
        cluster,
        instances: response.instances,
      });
    },
    async [Actions.LoadInstance]({ commit }, { cluster, instance }) {
      const response = await Api.get(`clusters/${cluster}/instances/${instance}`);
      commit(Mutations.AddInstance, {
        cluster,
        instance: response.instance,
      });
    },
  },
});
