import Vue from "vue";
import Buefy from "buefy";
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faCheck,
  faCheckCircle,
  faInfoCircle,
  faExclamationTriangle,
  faExclamationCircle,
  faArrowUp,
  faAngleRight,
  faAngleLeft,
  faAngleDown,
  faEye,
  faEyeSlash,
  faCaretDown,
  faCaretUp,
  faUpload,
  faUser,
  faLock,
  faPlus,
  faMinus,
  faUndo,
  faServer,
  faSkullCrossbones
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import filesize from "filesize";
import filters from "./data/filters";

library.add(
  faCheck,
  faCheckCircle,
  faInfoCircle,
  faExclamationTriangle,
  faExclamationCircle,
  faArrowUp,
  faAngleRight,
  faAngleLeft,
  faAngleDown,
  faEye,
  faEyeSlash,
  faCaretDown,
  faCaretUp,
  faUpload,
  faUser,
  faLock,
  faPlus,
  faMinus,
  faUndo,
  faServer,
  faSkullCrossbones
);
Vue.component("vue-fontawesome", FontAwesomeIcon);

Vue.use(Buefy, {
  defaultIconComponent: "vue-fontawesome",
  defaultIconPack: "fas"
});

Vue.config.productionTip = false;

for (const [key, value] of filters) {
  Vue.filter(key, value);
}

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");

Object.defineProperty(Vue.prototype, "$filesize", { value: filesize });
