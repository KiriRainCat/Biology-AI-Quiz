import type { RouteRecordRaw } from "vue-router";
import HomeView from "../../views/HomeView.vue";

export const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/quiz",
    name: "quiz",
    component: () => import("../../views/QuizView.vue"),
  },
  {
    path: "/result",
    name: "result",
    component: () => import("../../views/ResultView.vue"),
  },
  {
    path: "/leaderboard",
    name: "leaderboard",
    component: () => import("../../views/LeaderboardView.vue"),
  },
];
