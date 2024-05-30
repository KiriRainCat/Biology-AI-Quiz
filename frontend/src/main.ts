import { createApp } from "vue";

import App from "./App.vue";
import router from "./pkg/router";

// Varlet UI
import { Snackbar, Locale, Button, Loading, Paper, Table, Dialog, Input } from "@varlet/ui";
import "@varlet/ui/es/style";
import "@varlet/touch-emulator";
Snackbar.allowMultiple(true);
Locale.add("en-US", Locale.enUS);
Locale.use("en-US");

// 全局样式
import "./assets/styles.scss";

createApp(App).use(router).use(Button).use(Loading).use(Paper).use(Table).use(Dialog).use(Input).mount("#app");
