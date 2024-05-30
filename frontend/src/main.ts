import { createApp } from "vue";

import App from "./App.vue";
import router from "./pkg/router";

// Varlet UI
import Varlet, { Snackbar, Locale } from "@varlet/ui";
import "@varlet/ui/es/style";
import "@varlet/touch-emulator";
Snackbar.allowMultiple(true);
Locale.add("en-US", Locale.enUS);
Locale.use("en-US");

// 全局样式
import "./assets/styles.scss";

createApp(App).use(router).use(Varlet).mount("#app");
