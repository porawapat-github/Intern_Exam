import "./assets/tailwind.css";
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import axios from "axios";

axios.defaults.baseURL = "http://localhost:8080"; // ← backend รัน port นี้

const app = createApp(App);

app.use(router);
app.mount("#app");
