import axios from "axios";
import { Snackbar } from "@varlet/ui";
import type { Question, Record } from "./api_types";
import router from "../router";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: {
    Authorization: import.meta.env.VITE_API_AUTH,
  },
});

api.interceptors.response.use(
  (res) => {
    return res;
  },

  // 处理错误
  (error) => {
    try {
      if (error.code === "ERR_NETWORK" || error.response.status == 404 || error.response.status == 502) {
        Snackbar({ content: "服务器爆炸啦，也可能是你网炸了 :D", type: "error" });
        router.push("/");
        return Promise.reject(error);
      }

      if (error.response.data != undefined && error.response.data["msg"] != undefined) {
        Snackbar({ content: error.response.data["msg"], type: "error" });
        return Promise.reject(error);
      }
    } catch (_) {
      Snackbar({ content: `发生未知异常: ${error}`, type: "error" });
      router.push("/");
      return Promise.reject(error);
    }
  }
);

export default api;

async function getRecords(): Promise<Record[]> {
  let list: Record[] = (await api.get("/records")).data["data"];
  list.sort((a, b) => (a.accuracy > b.accuracy && a.time_taken < b.time_taken ? -1 : 1));
  return list;
}

async function addRecord(record: Record): Promise<boolean> {
  let res = await api.post(`/record`, record);
  if (res.status != 200) return false;
  return true;
}

async function generateQuiz(): Promise<Question[]> {
  return (await api.get("/generate-quiz")).data["data"];
}

export { getRecords, addRecord, generateQuiz };
